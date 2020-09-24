package task

type Task interface {
	UPID() string

	Node() string
	Action() Action
	ID() string
	User() string

	GetStatus() (Status, error)
	Wait() error
}

type VirtualMachineTask interface {
	Task

	VMID() uint
	IsQEMU() bool
	IsLXC() bool
}
