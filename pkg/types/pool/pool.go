package pool

type Pool interface {
	Name() string
	Description() (string, error)

	GetProperties() (PoolProperties, error)
	SetProperties(prop PoolProperties) error

	Delete(force bool) error

	ListMembers() ([]PoolMember, error)
}
