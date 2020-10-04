package test

import (
	"github.com/xabinapal/gopve/internal/service/vm"
	"github.com/xabinapal/gopve/pkg/client/test"
	"github.com/xabinapal/gopve/pkg/request/mocks"
	types "github.com/xabinapal/gopve/pkg/types/vm"
)

func NewVirtualMachine() (*vm.VirtualMachine, *test.API, *mocks.Executor) {
	svc, api, exc := NewService()

	return vm.NewVirtualMachine(
		svc,
		100,
		types.Kind("test_kind"),
		"test_node",
		false,
		nil,
	), api, exc
}

func NewQEMU() (*vm.QEMUVirtualMachine, *test.API, *mocks.Executor) {
	svc, api, exc := NewService()

	obj, _ := vm.NewDynamicVirtualMachine(
		svc,
		100,
		types.KindQEMU,
		"test_node",
		false,
		nil,
		nil,
	)

	return obj.(*vm.QEMUVirtualMachine), api, exc
}

func NewLXC() (*vm.LXCVirtualMachine, *test.API, *mocks.Executor) {
	svc, api, exc := NewService()

	obj, _ := vm.NewDynamicVirtualMachine(
		svc,
		100,
		types.KindLXC,
		"test_node",
		false,
		nil,
		nil,
	)

	return obj.(*vm.LXCVirtualMachine), api, exc
}
