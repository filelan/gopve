package service

import (
	"github.com/xabinapal/gopve/pkg/types/task"
	"github.com/xabinapal/gopve/pkg/types/vm"
	"github.com/xabinapal/gopve/pkg/types/vm/lxc"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
)

//go:generate mockery --case snake --name VirtualMachine --filename vm.go

type VirtualMachine interface {
	List() ([]vm.VirtualMachine, error)
	ListByKind(kind vm.Kind) ([]vm.VirtualMachine, error)
	Get(vmid uint) (vm.VirtualMachine, error)

	CreateQEMU(opts qemu.CreateOptions) (task.Task, error)
	CreateLXC(opts lxc.CreateOptions) (task.Task, error)

	DeleteQEMU(vmid uint, purge bool, force bool) (task.Task, error)
	DeleteLXC(vmid uint, purge bool, force bool) (task.Task, error)
}
