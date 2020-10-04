package vm

import (
	"fmt"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
)

type LXCVirtualMachine interface {
	VirtualMachine

	CPU() (LXCCPUProperties, error)
	Memory() (LXCMemoryProperties, error)

	GetLXCProperties() (LXCProperties, error)
	SetLXCProperties(props LXCProperties) error
}

type LXCCreateOptions struct {
	VMID uint
	Node string

	OSTemplateStorage string
	OSTemplate        string

	Properties LXCProperties
}

func (obj LXCCreateOptions) MapToValues() (request.Values, error) {
	values, err := obj.Properties.MapToValues()
	if err != nil {
		return nil, err
	}

	values.AddString(
		"ostemplate",
		fmt.Sprintf("%s:vztmpl/%s", obj.OSTemplateStorage, obj.OSTemplate),
	)

	return values, nil
}

type LXCProperties struct {
	CPU    LXCCPUProperties
	Memory LXCMemoryProperties

	RootFSStorage string
	RootFSSize    uint
}

func NewLXCProperties(props types.Properties) (*LXCProperties, error) {
	obj := new(LXCProperties)

	if v, err := NewLXCCPUProperties(props); err != nil {
		return nil, err
	} else {
		obj.CPU = *v
	}

	if v, err := NewLXCMemoryProperties(props); err != nil {
		return nil, err
	} else {
		obj.Memory = *v
	}

	return obj, nil
}

func (obj LXCProperties) MapToValues() (request.Values, error) {
	values := request.Values{}

	values.AddString(
		"rootfs",
		fmt.Sprintf("%s:%d", obj.RootFSStorage, obj.RootFSSize),
	)

	if cpuValues, err := obj.CPU.MapToValues(); err != nil {
		return nil, err
	} else {
		for k, v := range cpuValues {
			values[k] = v
		}
	}

	if memoryValues, err := obj.Memory.MapToValues(); err != nil {
		return nil, err
	} else {
		for k, v := range memoryValues {
			values[k] = v
		}
	}

	return values, nil
}

type LXCCPUProperties struct {
	Cores uint

	Limit uint
	Units uint
}

const (
	mkLXCCPUPropertyCores = "cores"
	mkLXCCPUPropertyLimit = "cpulimit"
	mkLXCCPUPropertyUnits = "cpuunits"

	DefaultLXCCPUPropertyLimit uint = 0
	DefaultLXCCPUPropertyUnits uint = 1024
)

func NewLXCCPUProperties(props types.Properties) (*LXCCPUProperties, error) {
	obj := new(LXCCPUProperties)

	if err := props.SetRequiredUint(mkLXCCPUPropertyCores, &obj.Cores, func(v uint) bool {
		return v <= 128
	}); err != nil {
		return nil, err
	}

	if err := props.SetUint(mkLXCCPUPropertyLimit, &obj.Limit, DefaultLXCCPUPropertyLimit, func(v uint) bool {
		return v <= 128
	}); err != nil {
		return nil, err
	}

	if err := props.SetUint(mkLXCCPUPropertyUnits, &obj.Units, DefaultLXCCPUPropertyUnits, func(v uint) bool {
		return v >= 8 && v <= 500000
	}); err != nil {
		return nil, err
	}

	return obj, nil
}

func (obj LXCCPUProperties) MapToValues() (request.Values, error) {
	values := request.Values{}

	cores := obj.Cores
	if cores == 0 {
		cores = 1
	} else if cores > 128 {
		return nil, fmt.Errorf("Invalid CPU sockets, the maximum allowed is 128")
	}
	values.AddUint("cores", cores)

	if obj.Limit > 128 {
		return nil, fmt.Errorf("Invalid CPU limit, must be between 0 and 128")
	} else if obj.Limit != 0 {
		values.AddUint("cpulimit", obj.Limit)
	}

	if obj.Units != 0 && (obj.Units < 2 || obj.Units > 500000) {
		return nil, fmt.Errorf(
			"Invalid CPU units, must be between 2 and 500000",
		)
	} else if obj.Units != 0 {
		values.AddUint("cpuunits", obj.Units)
	}

	return values, nil
}

type LXCMemoryProperties struct {
	Memory uint
	Swap   uint
}

const (
	mkLXCMemoryPropertyMemory = "memory"
	mkLXCMemoryPropertySwap   = "swap"
)

func NewLXCMemoryProperties(
	props types.Properties,
) (*LXCMemoryProperties, error) {
	obj := new(LXCMemoryProperties)

	if err := props.SetRequiredUint(mkQEMUMemoryPropertyMemory, &obj.Memory, nil); err != nil {
		return nil, err
	}

	if err := props.SetRequiredUint(mkLXCMemoryPropertySwap, &obj.Swap, nil); err != nil {
		return nil, err
	}

	return obj, nil
}

func (obj LXCMemoryProperties) MapToValues() (request.Values, error) {
	values := request.Values{}

	memory := obj.Memory
	if memory == 0 {
		memory = 512
	} else if memory < 16 {
		return nil, fmt.Errorf("Invalid memory, must at least 16")
	}
	values.AddUint("memory", memory)

	values.AddUint("swap", obj.Swap)

	return values, nil
}
