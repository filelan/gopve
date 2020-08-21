package vm

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

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

	values := url.Values{
		"vmid":  {strconv.Itoa(int(vm.vmid))},
		"newid": {strconv.Itoa(int(vmid))},
	}

	var task string
	err := vm.svc.Client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/qemu/%d/clone", vm.node, vm.vmid), values, &task)
	if err != nil {
		return nil, err
	}

	return vm.svc.Client.WaitForTask(task), nil
}
