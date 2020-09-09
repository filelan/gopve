package vm

import (
	"fmt"
	"net/http"
	"sort"
	"strconv"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

type listResponseJSON struct {
	VMID       uint          `json:"vmid"`
	Name       string        `json:"name"`
	Node       string        `json:"node"`
	Kind       vm.Kind       `json:"type"`
	IsTemplate types.PVEBool `json:"template"`
}

func (res listResponseJSON) Map(svc *Service) (vm.VirtualMachine, error) {
	if err := res.Kind.IsValid(); err != nil {
		return nil, fmt.Errorf("unsupported virtual machine kind")
	}

	out := VirtualMachine{
		svc:  svc,
		full: false,

		node:       res.Node,
		kind:       res.Kind,
		vmid:       res.VMID,
		name:       res.Name,
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

type getQEMUResponseJSON struct {
	Name       string        `json:"name"`
	IsTemplate types.PVEBool `json:"template"`

	CPUType            string        `json:"cpu"`
	CPUSockets         uint          `json:"sockets"`
	CPUCores           uint          `json:"cores"`
	CPUVCPUs           uint          `json:"vcpus"`
	CPULimit           string        `json:"cpulimit"`
	CPUUnits           uint          `json:"cpuunits"`
	NUMA               types.PVEBool `json:"numa"`
	FreezeCPUAtStartup types.PVEBool `json:"freeze"`

	Memory           uint  `json:"memory"`
	MemoryBallooning *uint `json:"balloon"`
	MemoryShares     *uint `json:"shares"`
}

func (res getQEMUResponseJSON) Map(svc *Service, vmid uint, node string) (vm.QEMUVirtualMachine, error) {
	limit, err := strconv.Atoi(res.CPULimit)
	if err != nil {
		return nil, err
	}

	vm := &QEMUVirtualMachine{
		VirtualMachine: VirtualMachine{
			svc:  svc,
			full: true,

			node:       node,
			kind:       vm.KindQEMU,
			vmid:       vmid,
			name:       res.Name,
			isTemplate: res.IsTemplate.Bool(),
		},

		cpu: vm.QEMUCPUProperties{
			Type:    res.CPUType,
			Sockets: res.CPUSockets,
			Cores:   res.CPUCores,
			VCPUs:   res.CPUVCPUs,

			Limit: uint(limit),
			Units: res.CPUUnits,

			NUMA: res.NUMA.Bool(),

			FreezeAtStartup: res.FreezeCPUAtStartup.Bool(),
		},

		memory: vm.QEMUMemoryProperties{
			Memory: res.Memory,
		},
	}

	if vm.cpu.VCPUs == 0 {
		vm.cpu.VCPUs = vm.cpu.Cores * vm.cpu.Sockets
	}

	if res.MemoryBallooning == nil {
		vm.memory.Ballooning = true
		vm.memory.MinimumMemory = vm.memory.Memory
		vm.memory.Shares = 0
	} else if *res.MemoryBallooning == 0 {
		vm.memory.Ballooning = false
		vm.memory.MinimumMemory = vm.memory.Memory
		vm.memory.Shares = 0
	} else {
		vm.memory.Ballooning = true
		vm.memory.MinimumMemory = uint(*res.MemoryBallooning)
		if res.MemoryShares == nil {
			vm.memory.Shares = 1000
		} else {
			vm.memory.Shares = *res.MemoryShares
		}
	}

	return vm, nil
}

type getLXCResponseJSON struct {
	Name       string        `json:"hostname"`
	IsTemplate types.PVEBool `json:"template"`

	OSType string `json:"ostype"`

	CPUCores uint `json:"cores"`
	CPULimit uint `json:"cpulimit"`
	CPUUnits uint `json:"cpuunits"`

	Memory uint `json:"memory"`
	Swap   uint `json:"swap"`

	RootFS string `json:"rootfs"`
}

func (res getLXCResponseJSON) Map(svc *Service, vmid uint, node string) (vm.LXCVirtualMachine, error) {
	return &LXCVirtualMachine{
		VirtualMachine: VirtualMachine{
			svc:  svc,
			full: true,

			node:       node,
			kind:       vm.KindLXC,
			vmid:       vmid,
			name:       res.Name,
			isTemplate: res.IsTemplate.Bool(),
		},

		cpu: vm.LXCCPUProperties{
			Cores: res.CPUCores,
			Limit: res.CPULimit,
			Units: res.CPUUnits,
		},

		memory: vm.LXCMemoryProperties{
			Memory: res.Memory,
			Swap:   res.Swap,
		},
	}, nil
}

func (svc *Service) Get(vmid uint) (vm.VirtualMachine, error) {
	vms, err := svc.List()
	if err != nil {
		return nil, nil
	}

	for _, virtualMachine := range vms {
		if virtualMachine.VMID() == vmid {
			switch virtualMachine.Kind() {
			case vm.KindQEMU:
				var res getQEMUResponseJSON
				if err := svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/qemu/%d/config", virtualMachine.Node(), virtualMachine.VMID()), nil, &res); err != nil {
					return nil, err
				}

				return res.Map(svc, virtualMachine.VMID(), virtualMachine.Node())

			case vm.KindLXC:
				var res getLXCResponseJSON
				if err := svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/lxc/%d/config", virtualMachine.Node(), virtualMachine.VMID()), nil, &res); err != nil {
					return nil, err
				}

				return res.Map(svc, virtualMachine.VMID(), virtualMachine.Node())

			default:
				panic("This should never happen")
			}
		}
	}

	return nil, fmt.Errorf("virtual machine not found")
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
