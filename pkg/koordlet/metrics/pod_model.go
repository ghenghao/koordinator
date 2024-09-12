package metrics

import "github.com/prometheus/client_golang/prometheus"

/*
labels:
	podName
	namespace
	cluster
	node
	partition
	pzone
	workload?
 */

const (
	// NameSpace represents pod namespace.
	NameSpace = "namespace"
)

var (
	podCpuUsage = prometheus.NewHistogramVec(prometheus.HistogramOpts{
		Subsystem: KoordletSubsystem,
		Name:      "pod_cpu_usage_his",
		Help:      "",
		Buckets: prometheus.ExponentialBuckets(0.01, 4, 8),
	}, []string{PodName, PodNamespace, NodeKey, ClusterKey, PartitionKey, PZoneKey})

	PodModelCollectors = []prometheus.Collector{
		podCpuUsage,
	}
)