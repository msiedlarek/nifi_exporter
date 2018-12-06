package collectors

import (
	"github.com/msiedlarek/nifi_exporter/nifi/client"
	"github.com/prometheus/client_golang/prometheus"
)

// ConnectionsCollector holds the metrics for each connection
type ConnectionsCollector struct {
	api *client.Client

	queuedCount *prometheus.Desc
	// bulletin5mCount *prometheus.Desc
	// componentCount  *prometheus.Desc

	// inFlowFiles5mCount          *prometheus.Desc
	// inBytes5mCount              *prometheus.Desc
	// queuedFlowFilesCount        *prometheus.Desc
	// queuedBytes                 *prometheus.Desc
	// readBytes5mCount            *prometheus.Desc
	// writtenBytes5mCount         *prometheus.Desc
	// outFlowFiles5mCount         *prometheus.Desc
	// outBytes5mCount             *prometheus.Desc
	// transferredFlowFiles5mCount *prometheus.Desc
	// transferredBytes5mCount     *prometheus.Desc
	// receivedBytes5mCount        *prometheus.Desc
	// receivedFlowFiles5mCount    *prometheus.Desc
	// sentBytes5mCount            *prometheus.Desc
	// sentFlowFiles5mCount        *prometheus.Desc
	// activeThreadCount           *prometheus.Desc
}

// NewConnectionsCollector initialises a collector
func NewConnectionsCollector(api *client.Client, labels map[string]string) *ConnectionsCollector {
	prefix := MetricNamePrefix + "conn_"
	statLabels := []string{"node_id", "source_name", "destination_name"}
	return &ConnectionsCollector{
		api: api,
		queuedCount: prometheus.NewDesc(
			prefix+"queued_count",
			"Number of items queued in connection",
			statLabels,
			labels,
		),
	}
}

// Describe makes the metrics descriptions available to Prometheus
func (c *ConnectionsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.queuedCount
}

// Collect retrieves the data that for the metrics
func (c *ConnectionsCollector) Collect(ch chan<- prometheus.Metric) {
	entities, err := c.api.GetProcessGroups(rootProcessGroupID)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.queuedCount, err)
		return
	}

	for i := range entities {
		c.collect(ch, &entities[i])
	}
}

func (c *ConnectionsCollector) collect(ch chan<- prometheus.Metric, entity *client.ProcessGroupEntity) {
	// bulletinCount := map[string]int{
	// 	"INFO":    0,
	// 	"WARNING": 0,
	// 	"ERROR":   0,
	// }
	// for i := range entity.Bulletins {
	// 	bulletinCount[entity.Bulletins[i].Bulletin.Level]++
	// }
	// for level, count := range bulletinCount {
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.bulletin5mCount,
	// 		prometheus.GaugeValue,
	// 		float64(count),
	// 		entity.Component.Name,
	// 		level,
	// 	)
	// }

	// nodes := make(map[string]*client.ProcessGroupStatusSnapshotDTO)
	// if len(entity.Status.NodeSnapshots) > 0 {
	// 	for i := range entity.Status.NodeSnapshots {
	// 		snapshot := &entity.Status.NodeSnapshots[i]
	// 		nodes[snapshot.NodeID] = &snapshot.StatusSnapshot
	// 	}
	// } else if entity.Status.AggregateSnapshot != nil {
	// 	nodes[AggregateNodeID] = entity.Status.AggregateSnapshot
	// }

	// ch <- prometheus.MustNewConstMetric(
	// 	c.componentCount,
	// 	prometheus.GaugeValue,
	// 	float64(entity.RunningCount),
	// 	entity.Component.Name,
	// 	"running",
	// )
	// ch <- prometheus.MustNewConstMetric(
	// 	c.componentCount,
	// 	prometheus.GaugeValue,
	// 	float64(entity.StoppedCount),
	// 	entity.Component.Name,
	// 	"stopped",
	// )
	// ch <- prometheus.MustNewConstMetric(
	// 	c.componentCount,
	// 	prometheus.GaugeValue,
	// 	float64(entity.InvalidCount),
	// 	entity.Component.Name,
	// 	"invalid",
	// )
	// ch <- prometheus.MustNewConstMetric(
	// 	c.componentCount,
	// 	prometheus.GaugeValue,
	// 	float64(entity.DisabledCount),
	// 	entity.Component.Name,
	// 	"disabled",
	// )

	// for nodeID, snapshot := range nodes {
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.inFlowFiles5mCount,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.FlowFilesIn),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.inBytes5mCount,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.BytesIn),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.queuedFlowFilesCount,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.FlowFilesQueued),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.queuedBytes,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.BytesQueued),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.readBytes5mCount,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.BytesRead),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.writtenBytes5mCount,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.BytesWritten),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.outFlowFiles5mCount,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.FlowFilesOut),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.outBytes5mCount,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.BytesOut),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.transferredFlowFiles5mCount,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.FlowFilesTransferred),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.transferredBytes5mCount,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.BytesTransferred),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.receivedBytes5mCount,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.BytesReceived),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.receivedFlowFiles5mCount,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.FlowFilesReceived),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.sentBytes5mCount,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.BytesSent),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.sentFlowFiles5mCount,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.FlowFilesSent),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// 	ch <- prometheus.MustNewConstMetric(
	// 		c.activeThreadCount,
	// 		prometheus.GaugeValue,
	// 		float64(snapshot.ActiveThreadCount),
	// 		nodeID,
	// 		snapshot.Name,
	// 	)
	// }
}
