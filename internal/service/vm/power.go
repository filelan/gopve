package vm

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/types/task"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

func postStatus(obj *VirtualMachine, command string) (task.Task, error) {
	if obj.template {
		return nil, fmt.Errorf("unsupported action on template virtual machine")
	}

	var task string
	if err := obj.svc.client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/%s/%d/status/%s", obj.node, string(obj.kind), obj.vmid, command), nil, &task); err != nil {
		return nil, err
	}

	return obj.svc.api.Task().Get(task)
}

func (obj *VirtualMachine) Start() (task.Task, error) {
	return postStatus(obj, "start")
}

func (obj *VirtualMachine) Stop() (task.Task, error) {
	return postStatus(obj, "stop")
}

func (obj *VirtualMachine) Reset() (task.Task, error) {
	switch obj.kind {
	case vm.KindLXC:
		obj.svc.client.StartAtomicBlock()
		defer obj.svc.client.EndAtomicBlock()
		postStatus(obj, "stop")
		return postStatus(obj, "start")
	default:
		return postStatus(obj, "reset")
	}
}

func (obj *VirtualMachine) Shutdown() (task.Task, error) {
	return postStatus(obj, "shutdown")
}

func (obj *VirtualMachine) Reboot() (task.Task, error) {
	return postStatus(obj, "reboot")
}

func (obj *VirtualMachine) Suspend() (task.Task, error) {
	return postStatus(obj, "suspend")
}

func (obj *VirtualMachine) Resume() (task.Task, error) {
	return postStatus(obj, "resume")
}
