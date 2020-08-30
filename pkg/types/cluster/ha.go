package cluster

import (
	"github.com/xabinapal/gopve/pkg/types/node"
)

type HighAvailabilityGroup interface {
	Name() string
	Description() (string, error)
	RestrictedResourceExecution() (bool, error)
	MigrateResourcesToHigherPriority() (bool, error)
	Nodes() (HighAvailabilityGroupNodes, error)

	GetProperties() (HighAvailabilityGroupProperties, error)
	SetProperties(props HighAvailabilityGroupProperties) error

	AddNodes(nodes map[string]uint) error
	DeleteNodes(nodes []string) error

	Delete() error
}

type HighAvailabilityGroupProperties struct {
	Description                      string
	RestrictedResourceExecution      bool
	MigrateResourcesToHigherPriority bool
}

type HighAvailabilityGroupNodes map[HighAvailabilityGroupNode]uint

type HighAvailabilityGroupNode interface {
	Name() string
	Get() (node.Node, error)
}
