package vm

import (
	"github.com/xabinapal/gopve/internal/client"
	"github.com/xabinapal/gopve/pkg/types/node"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

type Service struct {
	client client.Client
	api    client.API
}

func NewService(cli client.Client, api client.API) *Service {
	return &Service{
		client: cli,
		api:    api,
	}
}

type VirtualMachine struct {
	svc *Service

	node       string
	kind       vm.Kind
	vmid       uint
	name       string
	status     vm.Status
	isTemplate bool
}

type QEMUVirtualMachine struct {
	VirtualMachine
}

type LXCVirtualMachine struct {
	VirtualMachine
}

func (vm *VirtualMachine) Node() (node.Node, error) {
	return vm.svc.api.Node().Get(vm.node)
}

func (vm *VirtualMachine) Kind() vm.Kind {
	return vm.kind
}

func (vm *VirtualMachine) VMID() uint {
	return vm.vmid
}

func (vm *VirtualMachine) Name() string {
	return vm.name
}
func (vm *VirtualMachine) Status() vm.Status {
	return vm.status
}
func (vm *VirtualMachine) IsTemplate() bool {
	return vm.isTemplate
}
