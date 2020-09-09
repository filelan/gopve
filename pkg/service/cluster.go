package service

import (
	"github.com/xabinapal/gopve/pkg/types/cluster"
	"github.com/xabinapal/gopve/pkg/types/task"
)

//go:generate mockery --case snake --name Cluster

type Cluster interface {
	HA() HighAvailability

	Get() (cluster.Cluster, error)
	Create(name string, props cluster.NodeProperties) (task.Task, error)
	Join(hostname, password, fingerprint string, props cluster.NodeProperties) (task.Task, error)
}

type HighAvailability interface {
	ListGroups() ([]cluster.HighAvailabilityGroup, error)
	GetGroup(name string) (cluster.HighAvailabilityGroup, error)
	CreateGroup(name string, props cluster.HighAvailabilityGroupProperties, nodes cluster.HighAvailabilityGroupNodes) (cluster.HighAvailabilityGroup, error)
}
