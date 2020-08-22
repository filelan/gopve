package vm

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/xabinapal/gopve/internal/pkg/utils"

	"github.com/xabinapal/gopve/pkg/types"
)

func (vm *VirtualMachine) Clone(options types.VMCloneOptions) (types.Task, error) {
	vm.svc.Client.Lock()
	defer vm.svc.Client.Unlock()

	vmid := options.VMID
	if vmid == 0 {
		freeVMID, err := vm.svc.GetNextVMID()
		if err != nil {
			return nil, err
		}
		vmid = freeVMID
	}

	values := utils.RequestValues{
		"newid": {strconv.Itoa(int(vmid))},
	}

	if options.Name != "" {
		switch vm.category {
		case types.QEMU:
			values.AddString("name", options.Name)
		case types.LXC:
			values.AddString("hostname", options.Name)
		default:
			return nil, fmt.Errorf("unknown virtual machine category")
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
	if vm.isTemplate {
		values.AddBool("full", options.TemplateFullClone)
	} else {
		fullClone = true
	}

	if options.ImageFormat != "" {
		if vm.category == types.QEMU {
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
	err := vm.svc.Client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/qemu/%d/clone", vm.node, vm.vmid), values, &task)
	if err != nil {
		return nil, err
	}

	return vm.svc.Client.WaitForTask(task), nil
}
