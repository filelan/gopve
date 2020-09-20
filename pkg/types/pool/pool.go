package pool

type Pool interface {
	Name() string
	Description() (string, error)

	GetProperties() (PoolProperties, error)
	SetProperties(prop PoolProperties) error

	ListMembers() ([]PoolMember, error)

	AddVirtualMachine(vmid uint) error
	AddStorage(name string) error

	DeleteMember(member PoolMember) error
	DeleteVirtualMachine(vmid uint) error
	DeleteStorage(name string) error
}
