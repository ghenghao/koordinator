/*
Copyright 2022 The Koordinator Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/klog/v2"
)

const (
	KoordletSubsystem = "koordlet"

	NodeKey      = "node"
	ClusterKey   = "cluster"
	PZoneKey     = "pzone"
	RegionKey    = "region"
	PartitionKey = "partition"
	WorkloadName = "workload_name"

	ClusterLabelKeyName   = "topology.mizar-k8s.io/cluster"
	PZoneLabelKeyName     = "topology.mizar-k8s.io/pzone"
	RegionLabelKeyName    = "topology.mizar-k8s.io/region"
	PartitionLabelKeyName = "topology.mizar-k8s.io/partition"
	WorkloadLabelKeyName  = "workload.plan.mizar-universe.io/name"

	PriorityKey = "priority"

	PredictorKey = "predictor"

	StatusKey     = "status"
	StatusSucceed = "succeeded"
	StatusFailed  = "failed"

	EvictionReasonKey = "reason"
	BESuppressTypeKey = "type"

	ContainerID   = "container_id"
	ContainerName = "container_name"

	PodUID       = "pod_uid"
	PodName      = "pod_name"
	PodNamespace = "pod_namespace"

	ResourceKey = "resource"

	UnitKey     = "unit"
	UnitCore    = "core"
	UnitByte    = "byte"
	UnitInteger = "integer"
)

var (
	NodeName string
	// ClusterName from label key 'topology.mizar-k8s.io/cluster'
	ClusterName string

	// PZoneName from label key 'topology.mizar-k8s.io/pzone'
	PZoneName string

	// RegionName from label key 'topology.mizar-k8s.io/region'
	RegionName string

	// PartitionName from label key 'topology.mizar-k8s.io/partition'
	PartitionName string

	Node *corev1.Node

	nodeLock sync.RWMutex
)

const (
	DefaultHTTPPath = "/metrics"
)

// Register registers the metrics with the node object
func Register(node *corev1.Node) {
	nodeLock.Lock()
	defer nodeLock.Unlock()

	if node != nil {
		NodeName = node.Name
		labels := node.GetLabels()
		if labels != nil {
			ClusterName = labels[ClusterLabelKeyName]
			PZoneName = labels[PZoneLabelKeyName]
			RegionName = labels[RegionLabelKeyName]
			PartitionName = labels[PartitionLabelKeyName]
		} else {
			ClusterName = ""
			PZoneName = ""
			RegionName = ""
			PartitionName = ""
			klog.Warning("register nil node for metrics")
		}
	} else {
		NodeName = ""
		klog.Warning("register nil node for metrics")
	}
	Node = node
}

func genNodeLabels() prometheus.Labels {
	nodeLock.RLock()
	defer nodeLock.RUnlock()
	if Node == nil {
		return nil
	}

	return prometheus.Labels{
		NodeKey: NodeName,
	}
}

func genPodTopologyLabels() prometheus.Labels {
	nodeLock.RLock()
	defer nodeLock.RUnlock()
	if Node == nil {
		return nil
	}

	return prometheus.Labels{
		NodeKey:      NodeName,
		ClusterKey:   ClusterName,
		PZoneKey:     PZoneName,
		RegionKey:    RegionName,
		PartitionKey: PartitionName,
	}
}
