package types

type PoolService interface {
	List() ([]Pool, error)
	Get(id string) (Pool, error)
}

type Pool interface {
	Name() string
	Description() (string, error)

	GetProperties() (PoolProperties, error)
	SetProperties(prop PoolProperties) error

	Delete(force bool) error
}

type PoolProperties struct {
	Description string
}
