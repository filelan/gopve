package vm

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/xabinapal/gopve/internal/pkg/utils"
	"github.com/xabinapal/gopve/pkg/types"
)

type getResponseJSON struct {
	VMID       uint                         `json:"vmid"`
	Name       string                       `json:"name"`
	Node       string                       `json:"node"`
	Category   types.VirtualMachineCategory `json:"type"`
	Status     types.VirtualMachineStatus   `json:"status"`
	IsTemplate utils.PVEBool                `json:"template"`
}

func (res getResponseJSON) ConvertToEntity(svc *Service) (types.VirtualMachine, error) {
	if err := res.Category.IsValid(); err != nil {
		return nil, fmt.Errorf("unsupported virtual machine type")
	}

	if err := res.Status.IsValid(); err != nil {
		return nil, fmt.Errorf("unsupported virtual machine status")
	}

	out := VirtualMachine{
		svc: svc,

		node:       res.Node,
		category:   res.Category,
		vmid:       res.VMID,
		name:       res.Name,
		status:     res.Status,
		isTemplate: res.IsTemplate.Bool(),
	}

	switch res.Category {
	case types.QEMU:
		return &QEMUVirtualMachine{VirtualMachine: out}, nil
	case types.LXC:
		return &LXCVirtualMachine{VirtualMachine: out}, nil
	default:
		return nil, fmt.Errorf("unsupported virtual machine type")
	}
}

func (svc *Service) GetAll() ([]types.VirtualMachine, error) {
	var res []getResponseJSON
	if err := svc.Client.Request(http.MethodGet, "cluster/resources", url.Values{
		"type": {"vm"},
	}, &res); err != nil {
		return nil, err
	}

	vms := make([]types.VirtualMachine, len(res))
	for i, vm := range res {
		out, err := vm.ConvertToEntity(svc)
		if err != nil {
			return nil, err
		}

		vms[i] = out
	}

	return vms, nil
}

func (svc *Service) GetAllByType(category types.VirtualMachineCategory) ([]types.VirtualMachine, error) {
	var res []getResponseJSON
	if err := svc.Client.Request(http.MethodGet, "cluster/resources", url.Values{
		"type": {"vm"},
	}, &res); err != nil {
		return nil, err
	}

	vms := make([]types.VirtualMachine, len(res))
	for i, vm := range res {
		if vm.Category == category {
			out, err := vm.ConvertToEntity(svc)
			if err != nil {
				return nil, err
			}

			vms[i] = out
		}
	}

	return vms, nil
}

func (svc *Service) GetNextVMID() (uint, error) {
	var res string
	err := svc.Client.Request(http.MethodGet, "cluster/nextid", nil, &res)
	if err != nil {
		return 0, err
	}

	vmid, err := strconv.Atoi(res)
	if err != nil {
		return 0, err
	}

	return uint(vmid), nil
}
