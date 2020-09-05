package service

import (
	"github.com/xabinapal/gopve/pkg/types/cluster"
)

//go:generate mockery --case snake --name Cluster

type Cluster interface {
	HA() HighAvailability
}

type HighAvailability interface {
	ListGroups() ([]cluster.HighAvailabilityGroup, error)
	GetGroup(name string) (cluster.HighAvailabilityGroup, error)
	CreateGroup(name string, props cluster.HighAvailabilityGroupProperties, nodes cluster.HighAvailabilityGroupNodes) (cluster.HighAvailabilityGroup, error)
}
