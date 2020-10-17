package vm

import (
	"fmt"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/lxc"
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
	LXCGlobalProperties
	CPU    LXCCPUProperties
	Memory LXCMemoryProperties
}

func NewLXCProperties(props types.Properties) (*LXCProperties, error) {
	obj := new(LXCProperties)

	if v, err := NewLXCGlobalProperties(props); err != nil {
		return nil, err
	} else {
		obj.LXCGlobalProperties = *v
	}

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

type LXCGlobalProperties struct {
	OSType lxc.OSType

	Protected bool

	StartAtBoot bool

	RootFSStorage string
	RootFSSize    uint
}

const (
	mkLXCGlobalPropertyOSType      = "ostype"
	mkLXCGlobalPropertyProtected   = "protection"
	mkLXCGlobalPropertyStartAtBoot = "onboot"

	DefaultLXCGlobalPropertyProtected   bool = false
	DefaultLXCGlobalPropertyStartAtBoot bool = false
)

func NewLXCGlobalProperties(
	props types.Properties,
) (*LXCGlobalProperties, error) {
	obj := new(LXCGlobalProperties)

	if err := props.SetRequiredFixedValue(mkLXCGlobalPropertyOSType, &obj.OSType, nil); err != nil {
		return nil, err
	}

	if err := props.SetBool(mkLXCGlobalPropertyProtected, &obj.Protected, DefaultLXCGlobalPropertyProtected, nil); err != nil {
		return nil, err
	}

	if err := props.SetBool(mkLXCGlobalPropertyStartAtBoot, &obj.StartAtBoot, DefaultLXCGlobalPropertyStartAtBoot, nil); err != nil {
		return nil, err
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
