package lxc

import (
	"fmt"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
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
	vm.Properties
	GlobalProperties

	CPU    CPUProperties
	Memory MemoryProperties
}

func NewProperties(props types.Properties) (obj Properties, err error) {
	return obj, errors.ChainUntilFail(
		func() (err error) {
			obj.GlobalProperties, err = NewGlobalProperties(props)
			return err
		},
		func() (err error) {
			obj.CPU, err = NewCPUProperties(props)
			return err
		},
		func() (err error) {
			obj.Memory, err = NewMemoryProperties(props)
			return err
		},
	)

	return obj, nil
}

type GlobalProperties struct {
	OSType OSType

	RootFSStorage string
	RootFSSize    uint
}

const (
	mkGlobalPropertyOSType = "ostype"
)

func NewGlobalProperties(
	props types.Properties,
) (obj GlobalProperties, err error) {
	if err := props.SetRequiredFixedValue(mkGlobalPropertyOSType, &obj.OSType, nil); err != nil {
		return obj, err
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
