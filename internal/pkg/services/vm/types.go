package vm

import (
	"github.com/xabinapal/gopve/internal/pkg/utils"

	"github.com/xabinapal/gopve/pkg/types"
)

type Service struct {
	utils.Service
}

func NewService(api utils.API, client utils.Client) *Service {
	return &Service{
		Service: utils.Service{api, client},
	}
}

type VirtualMachine struct {
	svc *Service

	node       string
	category   types.VirtualMachineCategory
	vmid       uint
	name       string
	status     types.VirtualMachineStatus
	isTemplate bool
}

type QEMUVirtualMachine struct {
	VirtualMachine
}

type LXCVirtualMachine struct {
	VirtualMachine
}

func (vm *VirtualMachine) Node() (types.Node, error) {
	return vm.svc.API.Node().Get(vm.node)
}

func (vm *VirtualMachine) Category() types.VirtualMachineCategory {
	return vm.category
}

func (vm *VirtualMachine) VMID() uint {
	return vm.vmid
}

func (vm *VirtualMachine) Name() string {
	return vm.name
}
func (vm *VirtualMachine) Status() types.VirtualMachineStatus {
	return vm.status
}
func (vm *VirtualMachine) IsTemplate() bool {
	return vm.isTemplate
}
