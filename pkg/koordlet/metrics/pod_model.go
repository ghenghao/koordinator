package metrics

import "github.com/prometheus/client_golang/prometheus"

var (
	PodUsedCPU = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: KoordletSubsystem,
		Name:      "pod_cpu_usage",
		Help:      "cpu usage of pod in realtime",
	}, []string{NodeKey, ClusterKey, PZoneKey, RegionKey, PartitionKey, WorkloadName, PodName, PodNamespace, PodUID})

	PodUsedMemory = prometheus.NewGaugeVec(prometheus.GaugeOpts{
		Subsystem: KoordletSubsystem,
		Name:      "pod_memory_usage",
		Help:      "memory usage of pod in realtime",
	}, []string{NodeKey, ClusterKey, PZoneKey, RegionKey, PartitionKey, WorkloadName, PodName, PodNamespace, PodUID})

	PodCpuUsageHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: KoordletSubsystem,
		Name:      "pod_cpu_usage_hist",
		Help:      "cpu usage of pod in Histogram",
		Buckets:   prometheus.ExponentialBuckets(0.025, 1.5, 28),
	}, []string{NodeKey, ClusterKey, PZoneKey, RegionKey, PartitionKey, WorkloadName, PodName, PodNamespace, PodUID})

	PodMemoryUsageHist = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: KoordletSubsystem,
		Name:      "pod_memory_usage_hist",
		Help:      "memory usage of pod in Histogram",
		Buckets:   prometheus.ExponentialBuckets(5<<20, 2, 20),
	}, []string{NodeKey, ClusterKey, PZoneKey, RegionKey, PartitionKey, WorkloadName, PodName, PodNamespace, PodUID})

	PodModelCollectors = []prometheus.Collector{
		PodUsedCPU,
		PodUsedMemory,
		PodCpuUsageHist,
	}
)

func RecordPodUsedCPU(namespace, podName, podUID, workload string, value float64) {
	labels := genPodTopologyLabels()
	if labels == nil {
		return
	}
	labels[PodNamespace] = namespace
	labels[PodName] = podName
	labels[PodUID] = podUID
	labels[WorkloadName] = workload

	PodUsedCPU.With(labels).Set(value)
	PodCpuUsageHist.With(labels).Observe(value)
}

func RecordPodUsedMemory(namespace, podName, podUID, workload string, value float64) {
	labels := genPodTopologyLabels()
	if labels == nil {
		return
	}
	labels[PodNamespace] = namespace
	labels[PodName] = podName
	labels[PodUID] = podUID
	labels[WorkloadName] = workload

	PodUsedMemory.With(labels).Set(value)
	PodMemoryUsageHist.With(labels).Observe(value)
}
