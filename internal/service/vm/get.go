package vm

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

type getResponseJSON struct {
	VMID       uint          `json:"vmid"`
	Name       string        `json:"name"`
	Node       string        `json:"node"`
	Kind       vm.Kind       `json:"type"`
	Status     vm.Status     `json:"status"`
	IsTemplate types.PVEBool `json:"template"`
}

func (res getResponseJSON) Map(svc *Service) (vm.VirtualMachine, error) {
	if err := res.Kind.IsValid(); err != nil {
		return nil, fmt.Errorf("unsupported virtual machine kind")
	}

	if err := res.Status.IsValid(); err != nil {
		return nil, fmt.Errorf("unsupported virtual machine status")
	}

	out := VirtualMachine{
		svc: svc,

		node:       res.Node,
		kind:       res.Kind,
		vmid:       res.VMID,
		name:       res.Name,
		status:     res.Status,
		isTemplate: res.IsTemplate.Bool(),
	}

	switch res.Kind {
	case vm.KindQEMU:
		return &QEMUVirtualMachine{VirtualMachine: out}, nil
	case vm.KindLXC:
		return &LXCVirtualMachine{VirtualMachine: out}, nil
	default:
		return nil, fmt.Errorf("unsupported virtual machine type")
	}
}

func (svc *Service) List() ([]vm.VirtualMachine, error) {
	var res []getResponseJSON
	if err := svc.client.Request(http.MethodGet, "cluster/resources", request.Values{
		"type": {"vm"},
	}, &res); err != nil {
		return nil, err
	}

	vms := make([]vm.VirtualMachine, len(res))
	for i, vm := range res {
		out, err := vm.Map(svc)
		if err != nil {
			return nil, err
		}

		vms[i] = out
	}

	return vms, nil
}

func (svc *Service) GetAllByType(kind vm.Kind) ([]vm.VirtualMachine, error) {
	var res []getResponseJSON
	if err := svc.client.Request(http.MethodGet, "cluster/resources", request.Values{
		"type": {"vm"},
	}, &res); err != nil {
		return nil, err
	}

	vms := make([]vm.VirtualMachine, len(res))
	for i, vm := range res {
		if vm.Kind == kind {
			out, err := vm.Map(svc)
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
	err := svc.client.Request(http.MethodGet, "cluster/nextid", nil, &res)
	if err != nil {
		return 0, err
	}

	vmid, err := strconv.Atoi(res)
	if err != nil {
		return 0, err
	}

	return uint(vmid), nil
}
