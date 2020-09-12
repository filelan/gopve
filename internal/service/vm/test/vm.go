package test

import (
	"github.com/xabinapal/gopve/internal/service/vm"
	"github.com/xabinapal/gopve/pkg/request/mocks"
	types "github.com/xabinapal/gopve/pkg/types/vm"
)

func NewVirtualMachine() (*vm.VirtualMachine, *mocks.Executor) {
	svc, exc := NewService()
	return vm.NewVirtualMachine(svc, "test_node", types.Kind("test_kind"), 100), exc
}

func NewQEMU() (*vm.QEMUVirtualMachine, *mocks.Executor) {
	svc, exc := NewService()
	return vm.NewQEMU(svc, "test_node", 100), exc
}

func NewLXC() (*vm.LXCVirtualMachine, *mocks.Executor) {
	svc, exc := NewService()
	return vm.NewLXC(svc, "test_node", 100), exc
}
