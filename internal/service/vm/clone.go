package vm

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/task"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

func (obj *VirtualMachine) Clone(options vm.CloneOptions) (task.Task, error) {
	obj.svc.client.StartAtomicBlock()
	defer obj.svc.client.EndAtomicBlock()

	vmid := options.VMID
	if vmid == 0 {
		freeVMID, err := obj.svc.GetNextVMID()
		if err != nil {
			return nil, err
		}

		vmid = freeVMID
	}

	values := request.Values{
		"newid": {strconv.Itoa(int(vmid))},
	}

	if options.Name != "" {
		switch obj.kind {
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
	if obj.isTemplate {
		values.AddBool("full", options.TemplateFullClone)
	} else {
		fullClone = true
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

	var task string
	err := obj.svc.client.Request(
		http.MethodPost,
		fmt.Sprintf("nodes/%s/qemu/%d/clone", obj.node, obj.vmid),
		values,
		&task,
	)
	if err != nil {
		return nil, err
	}

	return obj.svc.api.Task().Get(task)
}
