package vm

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/node"
	"github.com/xabinapal/gopve/pkg/types/task"
	"github.com/xabinapal/gopve/pkg/types/vm/lxc"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
)

func (svc *Service) createVM(
	kind string,
	vmid uint,
	node string,
	values request.Values,
) (task.Task, error) {
	svc.client.StartAtomicBlock()
	defer svc.client.EndAtomicBlock()

	var err error

	vmid, err = svc.getVMID(vmid)
	if err != nil {
		return nil, err
	}

	values.AddUint("vmid", vmid)

	node, err = svc.getNode(node)
	if err != nil {
		return nil, err
	}

	var task string
	if err := svc.client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/%s", node, kind), values, &task); err != nil {
		return nil, err
	}

	return svc.api.Task().Get(task)
}

func (svc *Service) CreateQEMU(opts qemu.CreateOptions) (task.Task, error) {
	values, err := opts.MapToValues()
	if err != nil {
		return nil, err
	}

	return svc.createVM("qemu", opts.VMID, opts.Node, values)
}

func (svc *Service) CreateLXC(opts lxc.CreateOptions) (task.Task, error) {
	values, err := opts.MapToValues()
	if err != nil {
		return nil, err
	}

	return svc.createVM("lxc", opts.VMID, opts.Node, values)
}

func (svc *Service) getVMID(vmid uint) (uint, error) {
	if vmid == 0 {
		freeVMID, err := svc.GetNextVMID()
		if err != nil {
			return 0, err
		}

		return freeVMID, nil
	}

	return vmid, nil
}

func (svc *Service) getNode(n string) (string, error) {
	if n == "" {
		nodes, err := svc.api.Node().List()
		if err != nil {
			return "", err
		}

		for _, n := range nodes {
			if n.Status() == node.StatusOnline {
				return n.Name(), nil
			}
		}

		return "", fmt.Errorf(
			"cannot create virtual machine, there are no online nodes",
		)
	}

	if n, err := svc.api.Node().Get(n); err != nil {
		return "", err
	} else if n.Status() != node.StatusOnline {
		return "", fmt.Errorf("cannot create virtual machine, target node status is not online")
	}

	return n, nil
}
