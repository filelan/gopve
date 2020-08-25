package types

type HighAvailabilityService interface {
	ListGroups() ([]HighAvailabilityGroup, error)
	GetGroup(name string) (HighAvailabilityGroup, error)
	CreateGroup(name string, props HighAvailabilityGroupProperties, nodes HighAvailabilityGroupNodes) (HighAvailabilityGroup, error)
}

type HighAvailabilityGroup interface {
	Name() string
	Description() (string, error)
	RestrictedResourceExecution() (bool, error)
	MigrateResourcesToHigherPriority() (bool, error)
	Nodes() (HighAvailabilityGroupNodes, error)

	GetProperties() (HighAvailabilityGroupProperties, error)
	SetProperties(props HighAvailabilityGroupProperties) error
	Delete() error

	AddNodes(nodes map[string]uint) error
	DeleteNodes(nodes []string) error
}

type HighAvailabilityGroupProperties struct {
	Description                      string
	RestrictedResourceExecution      bool
	MigrateResourcesToHigherPriority bool
}

type HighAvailabilityGroupNodes map[HighAvailabilityGroupNode]uint

type HighAvailabilityGroupNode interface {
	Name() string
	Get() (Node, error)
}
