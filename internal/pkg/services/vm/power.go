package vm

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/types"
)

func postStatus(vm *VirtualMachine, command string) (types.Task, error) {
	if vm.isTemplate {
		return nil, fmt.Errorf("unsupported action on template virtual machine")
	}

	var task string
	err := vm.svc.Client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/%s/%d/status/%s", vm.node, string(vm.category), vm.vmid, command), nil, &task)
	if err != nil {
		return nil, err
	}

	return vm.svc.Client.WaitForTask(task), nil
}

func (vm *VirtualMachine) Start() (types.Task, error) {
	return postStatus(vm, "start")
}

func (vm *VirtualMachine) Stop() (types.Task, error) {
	return postStatus(vm, "stop")
}

func (vm *VirtualMachine) Reset() (types.Task, error) {
	switch vm.category {
	case types.LXC:
		vm.svc.Client.Lock()
		defer vm.svc.Client.Unlock()
		postStatus(vm, "stop")
		return postStatus(vm, "start")
	default:
		return postStatus(vm, "reset")
	}
}

func (vm *VirtualMachine) Shutdown() (types.Task, error) {
	return postStatus(vm, "shutdown")
}

func (vm *VirtualMachine) Reboot() (types.Task, error) {
	return postStatus(vm, "reboot")
}

func (vm *VirtualMachine) Suspend() (types.Task, error) {
	return postStatus(vm, "suspend")
}

func (vm *VirtualMachine) Resume() (types.Task, error) {
	return postStatus(vm, "resume")
}
