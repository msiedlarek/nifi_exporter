package collectors

import (
	"github.com/msiedlarek/nifi_exporter/nifi/client"
	"github.com/prometheus/client_golang/prometheus"
)

type generalMetrics struct {
	info   *prometheus.Desc
	osInfo *prometheus.Desc
}

type jvmMetrics struct {
	info                    *prometheus.Desc
	nonHeapBytes            *prometheus.Desc
	usedNonHeapBytes        *prometheus.Desc
	freeNonHeapBytes        *prometheus.Desc
	maxNonHeapBytes         *prometheus.Desc
	heapBytes               *prometheus.Desc
	usedHeapBytes           *prometheus.Desc
	freeHeapBytes           *prometheus.Desc
	maxHeapBytes            *prometheus.Desc
	availableProcessorCount *prometheus.Desc
	processorLoadAvg        *prometheus.Desc
	threadCount             *prometheus.Desc
	daemonThreadCount       *prometheus.Desc
	gcCollectionTotal       *prometheus.Desc
	gcCollectionTimeMsTotal *prometheus.Desc
}

type storageMetrics struct {
	freeBytes *prometheus.Desc
	sizeBytes *prometheus.Desc
	usedBytes *prometheus.Desc
}

type DiagnosticsCollector struct {
	api *client.Client

	generalMetrics
	jvmMetrics
	storageMetrics
}

func NewDiagnosticsCollector(api *client.Client, labels map[string]string) *DiagnosticsCollector {
	basicLabels := []string{"node_id"}
	jvmMetricsPrefix := MetricNamePrefix + "jvm_"
	storageMetricsPrefix := MetricNamePrefix + "stor_"
	return &DiagnosticsCollector{
		api: api,

		generalMetrics: generalMetrics{
			info: prometheus.NewDesc(
				MetricNamePrefix+"info",
				"NiFi version info.",
				append(basicLabels, "version"),
				labels,
			),
			osInfo: prometheus.NewDesc(
				MetricNamePrefix+"os_info",
				"Operating system version info.",
				append(basicLabels, "name", "version", "arch"),
				labels,
			),
		},
		jvmMetrics: jvmMetrics{
			info: prometheus.NewDesc(
				jvmMetricsPrefix+"info",
				"JVM version info.",
				append(basicLabels, "vendor", "version"),
				labels,
			),
			nonHeapBytes: prometheus.NewDesc(
				jvmMetricsPrefix+"non_heap_bytes",
				"Total number of bytes allocated to the JVM not used for heap.",
				basicLabels,
				labels,
			),
			usedNonHeapBytes: prometheus.NewDesc(
				jvmMetricsPrefix+"used_non_heap_bytes",
				"Total number of bytes used by the JVM not in the heap space.",
				basicLabels,
				labels,
			),
			freeNonHeapBytes: prometheus.NewDesc(
				jvmMetricsPrefix+"free_non_heap_bytes",
				"Total number of free non-heap bytes available to the JVM.",
				basicLabels,
				labels,
			),
			maxNonHeapBytes: prometheus.NewDesc(
				jvmMetricsPrefix+"max_non_heap_bytes",
				"The maximum number of bytes that the JVM can use for non-heap purposes.",
				basicLabels,
				labels,
			),
			heapBytes: prometheus.NewDesc(
				jvmMetricsPrefix+"heap_bytes",
				"The total number of bytes that are available for the JVM heap to use.",
				basicLabels,
				labels,
			),
			usedHeapBytes: prometheus.NewDesc(
				jvmMetricsPrefix+"used_heap_bytes",
				"The number of bytes of JVM heap that are currently being used.",
				basicLabels,
				labels,
			),
			freeHeapBytes: prometheus.NewDesc(
				jvmMetricsPrefix+"free_heap_bytes",
				"The number of bytes that are allocated to the JVM heap but not currently being used.",
				basicLabels,
				labels,
			),
			maxHeapBytes: prometheus.NewDesc(
				jvmMetricsPrefix+"max_heap_bytes",
				"The maximum number of bytes that can be used by the JVM.",
				basicLabels,
				labels,
			),
			availableProcessorCount: prometheus.NewDesc(
				jvmMetricsPrefix+"available_processor_count",
				"Number of available processors if supported by the underlying system.",
				basicLabels,
				labels,
			),
			processorLoadAvg: prometheus.NewDesc(
				jvmMetricsPrefix+"processor_load_avg",
				"The processor load average if supported by the underlying system.",
				basicLabels,
				labels,
			),
			threadCount: prometheus.NewDesc(
				jvmMetricsPrefix+"thread_count",
				"Total number of threads.",
				basicLabels,
				labels,
			),
			daemonThreadCount: prometheus.NewDesc(
				jvmMetricsPrefix+"daemon_thread_count",
				"Number of daemon threads.",
				basicLabels,
				labels,
			),
			gcCollectionTotal: prometheus.NewDesc(
				jvmMetricsPrefix+"gc_collection_total",
				"The number of times garbage collection has run.",
				append(basicLabels, "gc"),
				labels,
			),
			gcCollectionTimeMsTotal: prometheus.NewDesc(
				jvmMetricsPrefix+"gc_collection_time_ms_total",
				"The total number of milliseconds spent garbage collecting.",
				append(basicLabels, "gc"),
				labels,
			),
		},
		storageMetrics: storageMetrics{
			freeBytes: prometheus.NewDesc(
				storageMetricsPrefix+"free_bytes",
				"The number of bytes of free space.",
				append(basicLabels, "type", "location"),
				labels,
			),
			sizeBytes: prometheus.NewDesc(
				storageMetricsPrefix+"size_bytes",
				"The number of bytes of total space.",
				append(basicLabels, "type", "location"),
				labels,
			),
			usedBytes: prometheus.NewDesc(
				storageMetricsPrefix+"used_bytes",
				"The number of bytes of used space.",
				append(basicLabels, "type", "location"),
				labels,
			),
		},
	}
}

func (c *DiagnosticsCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- c.generalMetrics.info
	ch <- c.generalMetrics.osInfo

	ch <- c.jvmMetrics.info
	ch <- c.jvmMetrics.nonHeapBytes
	ch <- c.jvmMetrics.usedNonHeapBytes
	ch <- c.jvmMetrics.freeNonHeapBytes
	ch <- c.jvmMetrics.maxNonHeapBytes
	ch <- c.jvmMetrics.heapBytes
	ch <- c.jvmMetrics.usedHeapBytes
	ch <- c.jvmMetrics.freeHeapBytes
	ch <- c.jvmMetrics.maxHeapBytes
	ch <- c.jvmMetrics.availableProcessorCount
	ch <- c.jvmMetrics.processorLoadAvg
	ch <- c.jvmMetrics.threadCount
	ch <- c.jvmMetrics.daemonThreadCount

	ch <- c.storageMetrics.freeBytes
	ch <- c.storageMetrics.sizeBytes
	ch <- c.storageMetrics.usedBytes
}

func (c *DiagnosticsCollector) Collect(ch chan<- prometheus.Metric) {
	diagnostics, err := c.api.GetSystemDiagnostics(true, "")
	if err != nil {
		ch <- prometheus.NewInvalidMetric(c.generalMetrics.info, err)
		return
	}

	nodes := make(map[string]*client.SystemDiagnosticsSnapshotDTO)
	if len(diagnostics.NodeSnapshots) > 0 {
		for i := range diagnostics.NodeSnapshots {
			snapshot := &diagnostics.NodeSnapshots[i]
			nodes[snapshot.NodeID] = &snapshot.Snapshot
		}
	} else if diagnostics.AggregateSnapshot != nil {
		nodes[AggregateNodeID] = diagnostics.AggregateSnapshot
	}

	for nodeID, snapshot := range nodes {
		ch <- prometheus.MustNewConstMetric(
			c.generalMetrics.info,
			prometheus.GaugeValue,
			float64(1),
			nodeID,
			snapshot.VersionInfo.NiFiVersion,
		)
		ch <- prometheus.MustNewConstMetric(
			c.generalMetrics.osInfo,
			prometheus.GaugeValue,
			float64(1),
			nodeID,
			snapshot.VersionInfo.OsName,
			snapshot.VersionInfo.OsVersion,
			snapshot.VersionInfo.OsArchitecture,
		)

		ch <- prometheus.MustNewConstMetric(
			c.jvmMetrics.info,
			prometheus.GaugeValue,
			float64(1),
			nodeID,
			snapshot.VersionInfo.JavaVendor,
			snapshot.VersionInfo.JavaVersion,
		)
		ch <- prometheus.MustNewConstMetric(
			c.jvmMetrics.nonHeapBytes,
			prometheus.GaugeValue,
			float64(snapshot.TotalNonHeapBytes),
			nodeID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.jvmMetrics.usedNonHeapBytes,
			prometheus.GaugeValue,
			float64(snapshot.UsedNonHeapBytes),
			nodeID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.jvmMetrics.freeNonHeapBytes,
			prometheus.GaugeValue,
			float64(snapshot.FreeNonHeapBytes),
			nodeID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.jvmMetrics.maxNonHeapBytes,
			prometheus.GaugeValue,
			float64(snapshot.MaxNonHeapBytes),
			nodeID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.jvmMetrics.heapBytes,
			prometheus.GaugeValue,
			float64(snapshot.TotalHeapBytes),
			nodeID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.jvmMetrics.usedHeapBytes,
			prometheus.GaugeValue,
			float64(snapshot.UsedHeapBytes),
			nodeID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.jvmMetrics.freeHeapBytes,
			prometheus.GaugeValue,
			float64(snapshot.FreeHeapBytes),
			nodeID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.jvmMetrics.maxHeapBytes,
			prometheus.GaugeValue,
			float64(snapshot.MaxHeapBytes),
			nodeID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.jvmMetrics.availableProcessorCount,
			prometheus.GaugeValue,
			float64(snapshot.AvailableProcessors),
			nodeID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.jvmMetrics.processorLoadAvg,
			prometheus.GaugeValue,
			float64(snapshot.ProcessorLoadAverage),
			nodeID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.jvmMetrics.threadCount,
			prometheus.GaugeValue,
			float64(snapshot.TotalThreads),
			nodeID,
		)
		ch <- prometheus.MustNewConstMetric(
			c.jvmMetrics.daemonThreadCount,
			prometheus.GaugeValue,
			float64(snapshot.DaemonThreads),
			nodeID,
		)
		for i := range snapshot.GarbageCollection {
			stats := &snapshot.GarbageCollection[i]
			ch <- prometheus.MustNewConstMetric(
				c.jvmMetrics.gcCollectionTotal,
				prometheus.CounterValue,
				float64(stats.CollectionCount),
				nodeID,
				stats.Name,
			)
			ch <- prometheus.MustNewConstMetric(
				c.jvmMetrics.gcCollectionTimeMsTotal,
				prometheus.CounterValue,
				float64(stats.CollectionMillis),
				nodeID,
				stats.Name,
			)
		}

		c.collectStorageUsage(
			ch,
			nodeID,
			"flow_file",
			"default",
			&snapshot.FlowFileRepositoryStorageUsage,
		)
		for i := range snapshot.ContentRepositoryStorageUsage {
			usage := &snapshot.ContentRepositoryStorageUsage[i]
			c.collectStorageUsage(
				ch,
				nodeID,
				"content",
				usage.Identifier,
				usage,
			)
		}
		for i := range snapshot.ProvenanceRepositoryStorageUsage {
			usage := &snapshot.ProvenanceRepositoryStorageUsage[i]
			c.collectStorageUsage(
				ch,
				nodeID,
				"provenance",
				usage.Identifier,
				usage,
			)
		}
	}
}

func (c *DiagnosticsCollector) collectStorageUsage(ch chan<- prometheus.Metric, nodeID, storageType, location string, usage *client.StorageUsageDTO) {
	ch <- prometheus.MustNewConstMetric(
		c.storageMetrics.freeBytes,
		prometheus.GaugeValue,
		float64(usage.FreeSpaceBytes),
		nodeID,
		storageType,
		location,
	)
	ch <- prometheus.MustNewConstMetric(
		c.storageMetrics.sizeBytes,
		prometheus.GaugeValue,
		float64(usage.TotalSpaceBytes),
		nodeID,
		storageType,
		location,
	)
	ch <- prometheus.MustNewConstMetric(
		c.storageMetrics.usedBytes,
		prometheus.GaugeValue,
		float64(usage.UsedSpaceBytes),
		nodeID,
		storageType,
		location,
	)
}
