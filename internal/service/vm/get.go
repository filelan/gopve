package vm

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"

	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

type listResponseJSON struct {
	VMID     uint                   `json:"vmid"`
	Kind     vm.Kind                `json:"type"`
	Node     string                 `json:"node"`
	Name     string                 `json:"name"`
	Template internal_types.PVEBool `json:"template"`
}

func (res listResponseJSON) Map(svc *Service) (vm.VirtualMachine, error) {
	return NewDynamicVirtualMachine(
		svc,
		res.VMID,
		res.Kind,
		res.Node,
		res.Name,
		res.Template.Bool(),
		nil,
		nil,
	)
}

func (svc *Service) List() ([]vm.VirtualMachine, error) {
	var res []listResponseJSON
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

	sort.Slice(vms, func(i, j int) bool {
		return vms[i].VMID() < vms[j].VMID()
	})

	return vms, nil
}

func (svc *Service) ListByKind(kind vm.Kind) ([]vm.VirtualMachine, error) {
	var res []listResponseJSON
	if err := svc.client.Request(http.MethodGet, "cluster/resources", request.Values{
		"type": {"vm"},
	}, &res); err != nil {
		return nil, err
	}

	var vms []vm.VirtualMachine

	for _, vm := range res {
		if vm.Kind == kind {
			out, err := vm.Map(svc)
			if err != nil {
				return nil, err
			}

			vms = append(vms, out)
		}
	}

	sort.Slice(vms, func(i, j int) bool {
		return vms[i].VMID() < vms[j].VMID()
	})

	return vms, nil
}

type getResponseJSON struct {
	Name     string                 `json:"name"`
	Template internal_types.PVEBool `json:"template"`

	ExtraProperties types.Properties `json:"-"`
}

func (res *getResponseJSON) UnmarshalJSON(b []byte) error {
	type UnmarshalJSON getResponseJSON

	var x UnmarshalJSON

	if err := json.Unmarshal(b, &x); err != nil {
		return err
	}

	if err := json.Unmarshal(b, &x.ExtraProperties); err != nil {
		return err
	}

	*res = getResponseJSON(x)

	return nil
}

func (res getResponseJSON) Map(
	svc *Service,
	vmid uint,
	kind vm.Kind,
	node string,
) (vm.VirtualMachine, error) {
	props, err := vm.NewProperties(res.ExtraProperties)
	if err != nil {
		return nil, err
	}

	return NewDynamicVirtualMachine(
		svc,
		vmid,
		kind,
		node,
		res.Name,
		res.Template.Bool(),
		&props,
		res.ExtraProperties,
	)
}

func (svc *Service) Get(vmid uint) (vm.VirtualMachine, error) {
	vms, err := svc.List()
	if err != nil {
		return nil, err
	}

	for _, virtualMachine := range vms {
		if virtualMachine.VMID() == vmid {
			var res getResponseJSON

			switch virtualMachine.Kind() {
			case vm.KindQEMU:
				if err := svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/qemu/%d/config", virtualMachine.Node(), virtualMachine.VMID()), nil, &res); err != nil {
					return nil, err
				}

			case vm.KindLXC:
				if err := svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/lxc/%d/config", virtualMachine.Node(), virtualMachine.VMID()), nil, &res); err != nil {
					return nil, err
				}

			default:
				panic("This should never happen")
			}

			return res.Map(
				svc,
				virtualMachine.VMID(),
				virtualMachine.Kind(),
				virtualMachine.Node(),
			)
		}
	}

	return nil, vm.ErrNotFound
}

func (svc *Service) GetNextVMID() (uint, error) {
	var res string
	if err := svc.client.Request(http.MethodGet, "cluster/nextid", nil, &res); err != nil {
		return 0, err
	}

	vmid, err := strconv.Atoi(res)
	if err != nil {
		return 0, err
	}

	return uint(vmid), nil
}
