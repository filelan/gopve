package vm

import (
	"github.com/xabinapal/gopve/pkg/types/node"
	"github.com/xabinapal/gopve/pkg/types/task"
)

type VirtualMachine interface {
	Node() (node.Node, error)
	Kind() Kind

	VMID() uint
	Name() string
	Status() Status
	IsTemplate() bool

	Clone(options CloneOptions) (task.Task, error)

	Start() (task.Task, error)
	Stop() (task.Task, error)
	Reset() (task.Task, error)
	Shutdown() (task.Task, error)
	Reboot() (task.Task, error)
	Suspend() (task.Task, error)
	Resume() (task.Task, error)
}
