package vm

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/types/task"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

func (svc *Service) deleteVM(
	kind string,
	vmid uint,
	node string,
	purge bool,
	force bool,
) (task.Task, error) {
	var task string
	if err := svc.client.Request(http.MethodDelete, fmt.Sprintf("nodes/%s/%s/%d", node, kind, vmid), nil, &task); err != nil {
		return nil, err
	}

	return svc.api.Task().Get(task)
}

func (svc *Service) DeleteQEMU(
	vmid uint,
	purge bool,
	force bool,
) (task.Task, error) {
	virtualMachine, err := svc.Get(vmid)
	if err != nil {
		return nil, err
	}

	if virtualMachine.Kind() == vm.KindQEMU {
		return svc.deleteVM("qemu", vmid, virtualMachine.Node(), purge, force)
	}

	return nil, fmt.Errorf("invalid virtual machine kind")
}

func (svc *Service) DeleteLXC(
	vmid uint,
	purge bool,
	force bool,
) (task.Task, error) {
	virtualMachine, err := svc.Get(vmid)
	if err != nil {
		return nil, err
	}

	if virtualMachine.Kind() == vm.KindLXC {
		return svc.deleteVM("lxc", vmid, virtualMachine.Node(), purge, force)
	}

	return nil, fmt.Errorf("invalid virtual machine kind")
}
