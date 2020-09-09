package service

import (
	"github.com/xabinapal/gopve/pkg/types/task"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

//go:generate mockery --case snake --name VirtualMachine --filename vm.go

type VirtualMachine interface {
	List() ([]vm.VirtualMachine, error)
	ListByKind(kind vm.Kind) ([]vm.VirtualMachine, error)
	Get(vmid uint) (vm.VirtualMachine, error)

	CreateQEMU(opts vm.QEMUCreateOptions) (task.Task, error)
	CreateLXC(opts vm.LXCCreateOptions) (task.Task, error)
}
