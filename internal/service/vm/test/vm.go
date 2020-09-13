package test

import (
	"github.com/xabinapal/gopve/internal/service/vm"
	"github.com/xabinapal/gopve/pkg/client/test"
	"github.com/xabinapal/gopve/pkg/request/mocks"
	types "github.com/xabinapal/gopve/pkg/types/vm"
)

func NewVirtualMachine() (*vm.VirtualMachine, *test.API, *mocks.Executor) {
	svc, api, exc := NewService()
	return vm.NewVirtualMachine(svc, "test_node", types.Kind("test_kind"), 100), api, exc
}

func NewQEMU() (*vm.QEMUVirtualMachine, *test.API, *mocks.Executor) {
	svc, api, exc := NewService()
	return vm.NewQEMU(svc, "test_node", 100), api, exc
}

func NewLXC() (*vm.LXCVirtualMachine, *test.API, *mocks.Executor) {
	svc, api, exc := NewService()
	return vm.NewLXC(svc, "test_node", 100), api, exc
}
