package vm

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/node"
	"github.com/xabinapal/gopve/pkg/types/task"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

func (svc *Service) CreateQEMU(opts vm.QEMUCreateOptions) (task.Task, error) {
	svc.client.StartAtomicBlock()
	defer svc.client.EndAtomicBlock()

	values := request.Values{}

	if err := svc.createQEMUAddCPUProperties(values, opts.CPU); err != nil {
		return nil, err
	}

	if err := svc.createQEMUAddMemoryProperties(values, opts.Memory); err != nil {
		return nil, err
	}

	vmid, err := svc.getVMID(opts.VMID)
	if err != nil {
		return nil, err
	}

	values.AddUint("vmid", vmid)

	values.ConditionalAddString("name", opts.Name, opts.Name != "")
	values.ConditionalAddString("description", opts.Description, opts.Description != "")

	node, err := svc.getNode(opts.Node)
	if err != nil {
		return nil, err
	}

	var task string
	if err := svc.client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/qemu", node), values, &task); err != nil {
		return nil, err
	}

	return svc.api.Task().Get(task)
}

func (svc *Service) CreateLXC(opts vm.LXCCreateOptions) (task.Task, error) {
	svc.client.StartAtomicBlock()
	defer svc.client.EndAtomicBlock()

	values := request.Values{}

	if err := svc.createLXCAddCPUProperties(values, opts.CPU); err != nil {
		return nil, err
	}

	if err := svc.createLXCAddMemoryProperties(values, opts.Memory); err != nil {
		return nil, err
	}

	vmid, err := svc.getVMID(opts.VMID)
	if err != nil {
		return nil, err
	}

	values.AddUint("vmid", vmid)

	values.ConditionalAddString("hostname", opts.Name, opts.Name != "")
	values.ConditionalAddString("description", opts.Description, opts.Description != "")

	values.AddString("ostemplate", fmt.Sprintf("%s:vztmpl/%s", opts.OSTemplateStorage, opts.OSTemplate))
	values.AddString("rootfs", fmt.Sprintf("%s:%d", opts.RootFSStorage, opts.RootFSSize))

	node, err := svc.getNode(opts.Node)
	if err != nil {
		return nil, err
	}

	var task string
	if err := svc.client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/lxc", node), values, &task); err != nil {
		return nil, err
	}

	return svc.api.Task().Get(task)
}

func (svc *Service) createQEMUAddCPUProperties(values request.Values, props vm.QEMUCPUProperties) error {
	sockets := props.Sockets
	if sockets == 0 {
		sockets = 1
	} else if sockets > 4 {
		return fmt.Errorf("Invalid CPU sockets, the maximum allowed is 4")
	}
	values.AddUint("sockets", sockets)

	cores := props.Cores
	if cores == 0 {
		cores = 1
	} else if cores > 128 {
		return fmt.Errorf("Invalid CPU sockets, the maximum allowed is 128")
	}
	values.AddUint("cores", cores)

	if props.VCPUs != 0 && (props.VCPUs > sockets*cores) {
		return fmt.Errorf("Invalid CPU hotplugged cores, can't be greater than sockets * cores")
	} else if props.VCPUs != 0 {
		values.AddUint("vcpus", props.VCPUs)
	}

	if props.Limit > 128 {
		return fmt.Errorf("Invalid CPU limit, must be between 0 and 128")
	} else if props.Limit != 0 {
		values.AddUint("cpulimit", props.Limit)
	}

	if props.Units != 0 && (props.Units < 2 || props.Units > 262144) {
		return fmt.Errorf("Invalid CPU units, must be between 2 and 262144")
	} else if props.Units != 0 {
		values.AddUint("cpuunits", props.Units)
	}

	values.AddBool("numa", props.NUMA)

	values.AddBool("freeze", props.FreezeAtStartup)

	return nil
}

func (svc *Service) createQEMUAddMemoryProperties(values request.Values, props vm.QEMUMemoryProperties) error {
	memory := props.Memory
	if memory == 0 {
		memory = 512
	} else if memory < 16 || memory > 4178944 {
		return fmt.Errorf("Invalid memory, must be between 16 and 4178944")
	}
	values.AddUint("memory", memory)

	if props.Ballooning {
		minimumMemory := props.MinimumMemory
		if minimumMemory == 0 {
			minimumMemory = memory
		} else if minimumMemory > memory {
			return fmt.Errorf("Invalid Memory ballooning minimum, can't be greater than total memory")
		}

		values.AddUint("balloon", minimumMemory)

		if minimumMemory == memory {
			values.AddUint("shares", 0)
		} else {
			values.AddUint("shares", props.Shares)
		}

	} else {
		values.AddUint("balloon", 0)
		values.AddUint("shares", 0)
	}

	return nil
}

func (svc *Service) createLXCAddCPUProperties(values request.Values, props vm.LXCCPUProperties) error {
	cores := props.Cores
	if cores == 0 {
		cores = 1
	} else if cores > 128 {
		return fmt.Errorf("Invalid CPU sockets, the maximum allowed is 128")
	}
	values.AddUint("cores", cores)

	if props.Limit > 128 {
		return fmt.Errorf("Invalid CPU limit, must be between 0 and 128")
	} else if props.Limit != 0 {
		values.AddUint("cpulimit", props.Limit)
	}

	if props.Units != 0 && (props.Units < 2 || props.Units > 500000) {
		return fmt.Errorf("Invalid CPU units, must be between 2 and 500000")
	} else if props.Units != 0 {
		values.AddUint("cpuunits", props.Units)
	}

	return nil
}

func (svc *Service) createLXCAddMemoryProperties(values request.Values, props vm.LXCMemoryProperties) error {
	memory := props.Memory
	if memory == 0 {
		memory = 512
	} else if memory < 16 {
		return fmt.Errorf("Invalid memory, must at least 16")
	}
	values.AddUint("memory", memory)

	values.AddUint("swap", props.Swap)

	return nil
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

		return "", fmt.Errorf("cannot create virtual machine, there are no online nodes")
	}

	if n, err := svc.api.Node().Get(n); err != nil {
		return "", err
	} else if n.Status() != node.StatusOnline {
		return "", fmt.Errorf("cannot create virtual machine, target node status is not online")
	}

	return n, nil
}
