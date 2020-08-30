package vm

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/task"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

func (virtualMachine *VirtualMachine) Clone(options vm.CloneOptions) (task.Task, error) {
	virtualMachine.svc.client.StartAtomicBlock()
	defer virtualMachine.svc.client.EndAtomicBlock()

	vmid := options.VMID
	if vmid == 0 {
		freeVMID, err := virtualMachine.svc.GetNextVMID()
		if err != nil {
			return nil, err
		}
		vmid = freeVMID
	}

	values := request.Values{
		"newid": {strconv.Itoa(int(vmid))},
	}

	if options.Name != "" {
		switch virtualMachine.kind {
		case vm.KindQEMU:
			values.AddString("name", options.Name)
		case vm.KindLXC:
			values.AddString("hostname", options.Name)
		default:
			return nil, fmt.Errorf("unknown virtual machine kind")
		}
	}

	if options.Description != "" {
		values.AddString("description", options.Description)
	}

	if options.PoolName != "" {
		values.AddString("pool", options.PoolName)
	}

	if options.SnapshotName != "" {
		values.AddString("snapname", options.SnapshotName)
	}

	if options.BandwidthLimit != 0 {
		values.AddUint("bwlimit", options.BandwidthLimit)
	}

	fullClone := options.TemplateFullClone
	if virtualMachine.isTemplate {
		values.AddBool("full", options.TemplateFullClone)
	} else {
		fullClone = true
	}

	if options.ImageFormat != "" {
		if virtualMachine.kind == vm.KindQEMU {
			if !fullClone {
				return nil, fmt.Errorf("image format can only be specified when performing a full clone")
			} else if err := options.ImageFormat.IsValid(); err != nil {
				return nil, err
			} else {
				// TODO check if image format is compatible with target storage
				values.AddString("format", string(options.ImageFormat))
			}
		} else {
			return nil, fmt.Errorf("image format can only be specified in QEMU virtual machines")
		}
	}

	if options.TargetNode != "" {
		// TODO check if VM is on shared storage
		values.AddString("target", options.TargetNode)
	}

	if options.TargetStorage != "" {
		if fullClone {
			values.AddString("storage", options.TargetStorage)
		} else {
			return nil, fmt.Errorf("target storage can only be specified in full clone operations")
		}
	}

	fmt.Printf("%s\n", values)

	var task string
	err := virtualMachine.svc.client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/qemu/%d/clone", virtualMachine.node, virtualMachine.vmid), values, &task)
	if err != nil {
		return nil, err
	}

	return virtualMachine.svc.api.Task().Get(task)
}
