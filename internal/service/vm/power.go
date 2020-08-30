package vm

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/types/task"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

func postStatus(vm *VirtualMachine, command string) (task.Task, error) {
	if vm.isTemplate {
		return nil, fmt.Errorf("unsupported action on template virtual machine")
	}

	var task string
	if err := vm.svc.client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/%s/%d/status/%s", vm.node, string(vm.kind), vm.vmid, command), nil, &task); err != nil {
		return nil, err
	}

	return vm.svc.api.Task().Get(task)
}

func (vm *VirtualMachine) Start() (task.Task, error) {
	return postStatus(vm, "start")
}

func (vm *VirtualMachine) Stop() (task.Task, error) {
	return postStatus(vm, "stop")
}

func (virtualMachine *VirtualMachine) Reset() (task.Task, error) {
	switch virtualMachine.kind {
	case vm.KindLXC:
		virtualMachine.svc.client.StartAtomicBlock()
		defer virtualMachine.svc.client.EndAtomicBlock()
		postStatus(virtualMachine, "stop")
		return postStatus(virtualMachine, "start")
	default:
		return postStatus(virtualMachine, "reset")
	}
}

func (vm *VirtualMachine) Shutdown() (task.Task, error) {
	return postStatus(vm, "shutdown")
}

func (vm *VirtualMachine) Reboot() (task.Task, error) {
	return postStatus(vm, "reboot")
}

func (vm *VirtualMachine) Suspend() (task.Task, error) {
	return postStatus(vm, "suspend")
}

func (vm *VirtualMachine) Resume() (task.Task, error) {
	return postStatus(vm, "resume")
}
