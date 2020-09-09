package vm

import (
	"fmt"

	"github.com/xabinapal/gopve/pkg/request"
)

type LXCVirtualMachine interface {
	VirtualMachine

	CPU() (LXCCPUProperties, error)
	Memory() (LXCMemoryProperties, error)

	SetProperties(props LXCProperties) error
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

	values.AddString("ostemplate", fmt.Sprintf("%s:vztmpl/%s", obj.OSTemplateStorage, obj.OSTemplate))

	return values, nil
}

type LXCProperties struct {
	Name        string
	Description string

	CPU    LXCCPUProperties
	Memory LXCMemoryProperties

	RootFSStorage string
	RootFSSize    uint
}

func (obj LXCProperties) MapToValues() (request.Values, error) {
	var values request.Values

	values.ConditionalAddString("hostname", obj.Name, obj.Name != "")
	values.ConditionalAddString("description", obj.Description, obj.Description != "")

	values.AddString("rootfs", fmt.Sprintf("%s:%d", obj.RootFSStorage, obj.RootFSSize))

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

func (obj LXCCPUProperties) MapToValues() (request.Values, error) {
	var values request.Values

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
		return nil, fmt.Errorf("Invalid CPU units, must be between 2 and 500000")
	} else if obj.Units != 0 {
		values.AddUint("cpuunits", obj.Units)
	}

	return values, nil
}

type LXCMemoryProperties struct {
	Memory uint
	Swap   uint
}

func (obj LXCMemoryProperties) MapToValues() (request.Values, error) {
	var values request.Values

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
