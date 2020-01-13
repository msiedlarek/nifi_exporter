package collectors

import (
	"github.com/msiedlarek/nifi_exporter/nifi/client"
	"github.com/prometheus/client_golang/prometheus"
)

type ProcessGroupsCollector struct {
	api *client.Client

	bulletin5mCount *prometheus.Desc
	componentCount  *prometheus.Desc

	inFlowFiles5mCount          *prometheus.Desc
	inBytes5mCount              *prometheus.Desc
	queuedFlowFilesCount        *prometheus.Desc
	queuedBytes                 *prometheus.Desc
	readBytes5mCount            *prometheus.Desc
	writtenBytes5mCount         *prometheus.Desc
	outFlowFiles5mCount         *prometheus.Desc
	outBytes5mCount             *prometheus.Desc
	transferredFlowFiles5mCount *prometheus.Desc
	transferredBytes5mCount     *prometheus.Desc
	receivedBytes5mCount        *prometheus.Desc
	receivedFlowFiles5mCount    *prometheus.Desc
	sentBytes5mCount            *prometheus.Desc
	sentFlowFiles5mCount        *prometheus.Desc
	activeThreadCount           *prometheus.Desc
}

func NewProcessGroupsCollector(api *client.Client, labels map[string]string) *ProcessGroupsCollector {
	prefix := MetricNamePrefix + "pg_"
	//statLabels := []string{"node_id", "group"}
	statLabelsId := []string{"node_id", "group", "entity_id"}
	return &ProcessGroupsCollector{
		api: api,

		bulletin5mCount: prometheus.NewDesc(
			prefix+"bulletin_5m_count",
			"Number of bulletins posted during last 5 minutes.",
			[]string{"group", "level", "entity_id"},
			labels,
		),
		componentCount: prometheus.NewDesc(
			prefix+"component_count",
			"The number of components in this process group.",
			[]string{"group", "status", "entity_id"},
			labels,
		),

		inFlowFiles5mCount: prometheus.NewDesc(
			prefix+"in_flow_files_5m_count",
			"The number of FlowFiles that have come into this ProcessGroup in the last 5 minutes",
			statLabelsId,
			labels,
		),
		inBytes5mCount: prometheus.NewDesc(
			prefix+"in_bytes_5m_count",
			"The number of bytes that have come into this ProcessGroup in the last 5 minutes",
			statLabelsId,
			labels,
		),
		queuedFlowFilesCount: prometheus.NewDesc(
			prefix+"queued_flow_files_count",
			"The number of FlowFiles that are queued up in this ProcessGroup right now",
			statLabelsId,
			labels,
		),
		queuedBytes: prometheus.NewDesc(
			prefix+"queued_bytes",
			"The number of bytes that are queued up in this ProcessGroup right now",
			statLabelsId,
			labels,
		),
		readBytes5mCount: prometheus.NewDesc(
			prefix+"read_bytes_5m_count",
			"The number of bytes read by components in this ProcessGroup in the last 5 minutes",
			statLabelsId,
			labels,
		),
		writtenBytes5mCount: prometheus.NewDesc(
			prefix+"written_bytes_5m_count",
			"The number of bytes written by components in this ProcessGroup in the last 5 minutes",
			statLabelsId,
			labels,
		),
		outFlowFiles5mCount: prometheus.NewDesc(
			prefix+"out_flow_files_5m_count",
			"The number of FlowFiles transferred out of this ProcessGroup in the last 5 minutes",
			statLabelsId,
			labels,
		),
		outBytes5mCount: prometheus.NewDesc(
			prefix+"out_bytes_5m_count",
			"The number of bytes transferred out of this ProcessGroup in the last 5 minutes",
			statLabelsId,
			labels,
		),
		transferredFlowFiles5mCount: prometheus.NewDesc(
			prefix+"transferred_flow_files_5m_count",
			"The number of FlowFiles transferred in this ProcessGroup in the last 5 minutes",
			statLabelsId,
			labels,
		),
		transferredBytes5mCount: prometheus.NewDesc(
			prefix+"transferred_bytes_5m_count",
			"The number of bytes transferred in this ProcessGroup in the last 5 minutes",
			statLabelsId,
			labels,
		),
		receivedBytes5mCount: prometheus.NewDesc(
			prefix+"received_bytes_5m_count",
			"The number of bytes received from external sources by components within this ProcessGroup in the last 5 minutes",
			statLabelsId,
			labels,
		),
		receivedFlowFiles5mCount: prometheus.NewDesc(
			prefix+"received_flow_files_5m_count",
			"The number of FlowFiles received from external sources by components within this ProcessGroup in the last 5 minutes",
			statLabelsId,
			labels,
		),
		sentBytes5mCount: prometheus.NewDesc(
			prefix+"sent_bytes_5m_count",
			"The number of bytes sent to an external sink by components within this ProcessGroup in the last 5 minutes",
			statLabelsId,
			labels,
		),
		sentFlowFiles5mCount: prometheus.NewDesc(
			prefix+"sent_flow_files_5m_count",
			"The number of FlowFiles sent to an external sink by components within this ProcessGroup in the last 5 minutes",
			statLabelsId,
			labels,
		),
		activeThreadCount: prometheus.NewDesc(
			prefix+"active_thread_count",
			"The active thread count for this process group.",
			statLabelsId,
			labels,
		),
	}
}

func (c *ProcessGroupsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.bulletin5mCount
	ch <- c.componentCount

	ch <- c.inFlowFiles5mCount
	ch <- c.inBytes5mCount
	ch <- c.queuedFlowFilesCount
	ch <- c.queuedBytes
	ch <- c.readBytes5mCount
	ch <- c.writtenBytes5mCount
	ch <- c.outFlowFiles5mCount
	ch <- c.outBytes5mCount
	ch <- c.transferredFlowFiles5mCount
	ch <- c.transferredBytes5mCount
	ch <- c.receivedBytes5mCount
	ch <- c.receivedFlowFiles5mCount
	ch <- c.sentBytes5mCount
	ch <- c.sentFlowFiles5mCount
	ch <- c.activeThreadCount
}

func (c *ProcessGroupsCollector) Collect(ch chan<- prometheus.Metric) {
	entities, err := c.api.GetDeepProcessGroups(rootProcessGroupID)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.componentCount, err)
		return
	}

	for i := range entities {
		c.collect(ch, &entities[i])
	}
}

func (c *ProcessGroupsCollector) collect(ch chan<- prometheus.Metric, entity *client.ProcessGroupEntity) {
	bulletinCount := map[string]int{
		"INFO":    0,
		"WARNING": 0,
		"ERROR":   0,
	}
	for i := range entity.Bulletins {
		bulletinCount[entity.Bulletins[i].Bulletin.Level]++
	}


	for level, count := range bulletinCount {
		ch <- prometheus.MustNewConstMetric(
			c.bulletin5mCount,
			prometheus.GaugeValue,
			float64(count),
			entity.Component.Name,
			level,
			entity.Component.ID,
		)
	}

	nodes := make(map[string]*client.ProcessGroupStatusSnapshotDTO)
	if len(entity.Status.NodeSnapshots) > 0 {
		for i := range entity.Status.NodeSnapshots {
			snapshot := &entity.Status.NodeSnapshots[i]
			nodes[snapshot.NodeID] = &snapshot.StatusSnapshot
		}
	} else if entity.Status.AggregateSnapshot != nil {
		nodes[AggregateNodeID] = entity.Status.AggregateSnapshot
	}

	ch <- prometheus.MustNewConstMetric(
		c.componentCount,
		prometheus.GaugeValue,
		float64(entity.RunningCount),
		entity.Component.Name,
		"running",
		entity.Component.ID,
	)
	ch <- prometheus.MustNewConstMetric(
		c.componentCount,
		prometheus.GaugeValue,
		float64(entity.StoppedCount),
		entity.Component.Name,
		"stopped",
		entity.Component.ID,
	)
	ch <- prometheus.MustNewConstMetric(
		c.componentCount,
		prometheus.GaugeValue,
		float64(entity.InvalidCount),
		entity.Component.Name,
		"invalid",
		entity.Component.ID,
	)
	ch <- prometheus.MustNewConstMetric(
		c.componentCount,
		prometheus.GaugeValue,
		float64(entity.DisabledCount),
		entity.Component.Name,
		"disabled",
		entity.Component.ID,
	)

	for nodeID, snapshot := range nodes {
		ch <- prometheus.MustNewConstMetric(
			c.inFlowFiles5mCount,
			prometheus.GaugeValue,
			float64(snapshot.FlowFilesIn),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.inBytes5mCount,
			prometheus.GaugeValue,
			float64(snapshot.BytesIn),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.queuedFlowFilesCount,
			prometheus.GaugeValue,
			float64(snapshot.FlowFilesQueued),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.queuedBytes,
			prometheus.GaugeValue,
			float64(snapshot.BytesQueued),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.readBytes5mCount,
			prometheus.GaugeValue,
			float64(snapshot.BytesRead),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.writtenBytes5mCount,
			prometheus.GaugeValue,
			float64(snapshot.BytesWritten),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.outFlowFiles5mCount,
			prometheus.GaugeValue,
			float64(snapshot.FlowFilesOut),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.outBytes5mCount,
			prometheus.GaugeValue,
			float64(snapshot.BytesOut),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.transferredFlowFiles5mCount,
			prometheus.GaugeValue,
			float64(snapshot.FlowFilesTransferred),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.transferredBytes5mCount,
			prometheus.GaugeValue,
			float64(snapshot.BytesTransferred),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.receivedBytes5mCount,
			prometheus.GaugeValue,
			float64(snapshot.BytesReceived),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.receivedFlowFiles5mCount,
			prometheus.GaugeValue,
			float64(snapshot.FlowFilesReceived),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.sentBytes5mCount,
			prometheus.GaugeValue,
			float64(snapshot.BytesSent),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.sentFlowFiles5mCount,
			prometheus.GaugeValue,
			float64(snapshot.FlowFilesSent),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.activeThreadCount,
			prometheus.GaugeValue,
			float64(snapshot.ActiveThreadCount),
			nodeID,
			snapshot.Name,
			entity.Component.ID,
		)
	}
}
