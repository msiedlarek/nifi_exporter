package collectors

import (
	"github.com/msiedlarek/nifi_exporter/nifi/client"
	"github.com/prometheus/client_golang/prometheus"
)

// ConnectionsCollector holds the metrics for each connection
type ConnectionsCollector struct {
	api *client.Client

	flowFilesQueued *prometheus.Desc
}

// NewConnectionsCollector initialises a collector
func NewConnectionsCollector(api *client.Client, labels map[string]string) *ConnectionsCollector {
	prefix := MetricNamePrefix + "conn_"
	statLabels := []string{"node_id", "connection_name", "connection_id", "group_id", "source_name", "destination_name"}
	return &ConnectionsCollector{
		api: api,
		flowFilesQueued: prometheus.NewDesc(
			prefix+"flow_files_queued",
			"The number of FlowFiles that are currently queued in the connection",
			statLabels,
			labels,
		),
	}
}

// Describe makes the metrics descriptions available to Prometheus
func (c *ConnectionsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.flowFilesQueued
}

// Collect retrieves the data that for the metrics
func (c *ConnectionsCollector) Collect(ch chan<- prometheus.Metric) {
	entities, err := c.api.GetConnections(rootProcessGroupID)
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.flowFilesQueued, err)
		return
	}

	for i := range entities {
		c.collect(ch, &entities[i])
	}
}

func (c *ConnectionsCollector) collect(ch chan<- prometheus.Metric, entity *client.ConnectionEntity) {
	nodes := make(map[string]*client.ConnectionStatusSnapshotDTO)
	if entity.Status.AggregateSnapshot != nil {
		nodes[AggregateNodeID] = entity.Status.AggregateSnapshot
	}

	for nodeID, snapshot := range nodes {
		ch <- prometheus.MustNewConstMetric(
			c.flowFilesQueued,
			prometheus.GaugeValue,
			float64(snapshot.FlowFilesQueued),
			nodeID,
			snapshot.Name,
			snapshot.ID,
			snapshot.GroupID,
			snapshot.SourceName,
			snapshot.DestinationName,
		)
	}
}
