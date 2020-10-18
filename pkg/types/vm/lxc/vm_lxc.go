package lxc

import (
	"fmt"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

type VirtualMachine interface {
	vm.VirtualMachine

	CPU() (CPUProperties, error)
	Memory() (MemoryProperties, error)

	GetLXCProperties() (Properties, error)
	SetLXCProperties(props Properties) error
}

type CreateOptions struct {
	VMID uint
	Node string

	OSTemplateStorage string
	OSTemplate        string

	Properties Properties
}

func (obj CreateOptions) MapToValues() (request.Values, error) {
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

type Properties struct {
	GlobalProperties
	CPU    CPUProperties
	Memory MemoryProperties
}

func NewProperties(props types.Properties) (*Properties, error) {
	obj := new(Properties)

	if v, err := NewGlobalProperties(props); err != nil {
		return nil, err
	} else {
		obj.GlobalProperties = *v
	}

	if v, err := NewCPUProperties(props); err != nil {
		return nil, err
	} else {
		obj.CPU = *v
	}

	if v, err := NewMemoryProperties(props); err != nil {
		return nil, err
	} else {
		obj.Memory = *v
	}

	return obj, nil
}

type GlobalProperties struct {
	OSType OSType

	Protected bool

	StartAtBoot bool

	RootFSStorage string
	RootFSSize    uint
}

const (
	mkGlobalPropertyOSType      = "ostype"
	mkGlobalPropertyProtected   = "protection"
	mkGlobalPropertyStartAtBoot = "onboot"

	DefaultGlobalPropertyProtected   bool = false
	DefaultGlobalPropertyStartAtBoot bool = false
)

func NewGlobalProperties(
	props types.Properties,
) (*GlobalProperties, error) {
	obj := new(GlobalProperties)

	if err := props.SetRequiredFixedValue(mkGlobalPropertyOSType, &obj.OSType, nil); err != nil {
		return nil, err
	}

	if err := props.SetBool(mkGlobalPropertyProtected, &obj.Protected, DefaultGlobalPropertyProtected, nil); err != nil {
		return nil, err
	}

	if err := props.SetBool(mkGlobalPropertyStartAtBoot, &obj.StartAtBoot, DefaultGlobalPropertyStartAtBoot, nil); err != nil {
		return nil, err
	}

	return obj, nil
}

func (obj Properties) MapToValues() (request.Values, error) {
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
