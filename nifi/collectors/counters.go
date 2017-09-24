package collectors

import (
	"github.com/juju/errors"
	"github.com/msiedlarek/nifi_exporter/nifi/client"
	"github.com/prometheus/client_golang/prometheus"
)

type CountersCollector struct {
	alias              string
	api                *client.Client
	counterTotalMetric *prometheus.Desc
}

func NewCountersCollector(api *client.Client, labels map[string]string) *CountersCollector {
	return &CountersCollector{
		api: api,
		counterTotalMetric: prometheus.NewDesc(
			MetricNamePrefix+"counter_total",
			"The value of the counter.",
			[]string{"node_id", "id", "context", "name"},
			labels,
		),
	}
}

func (c *CountersCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.counterTotalMetric
}

func (c *CountersCollector) Collect(ch chan<- prometheus.Metric) {
	counterStats, err := c.api.GetCounters(true, "")
	if err != nil {
		ch <- prometheus.NewInvalidMetric(
			c.counterTotalMetric,
			errors.Annotatef(err, "Cannot retrieve metrics for node '%s'", c.alias),
		)
		return
	}

	nodes := make(map[string][]client.CounterDTO)
	if len(counterStats.NodeSnapshots) > 0 {
		for i := range counterStats.NodeSnapshots {
			snapshot := &counterStats.NodeSnapshots[i]
			nodes[snapshot.NodeID] = snapshot.Snapshot.Counters
		}
	} else if counterStats.AggregateSnapshot != nil {
		nodes[AggregateNodeID] = counterStats.AggregateSnapshot.Counters
	}

	for nodeID, counters := range nodes {
		for i := range counters {
			counter := &counters[i]
			ch <- prometheus.MustNewConstMetric(
				c.counterTotalMetric,
				prometheus.CounterValue,
				float64(counter.ValueCount),
				nodeID,
				counter.ID,
				counter.Context,
				counter.Name,
			)
		}
	}
}
