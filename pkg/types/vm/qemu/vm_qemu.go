package qemu

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

	GetQEMUProperties() (Properties, error)
	SetQEMUProperties(props Properties) error
}

type CreateOptions struct {
	VMID uint
	Node string

	Properties Properties
}

func (obj CreateOptions) MapToValues() (request.Values, error) {
	values, err := obj.Properties.MapToValues()
	if err != nil {
		return nil, err
	}

	return values, nil
}

type Properties struct {
	GlobalProperties
	CPU     CPUProperties
	Memory  MemoryProperties
	Storage StorageProperties
	Network []NetworkInterfaceProperties
}

const (
	maxNetworkInterfacePropertiesArrayCapacity = 32
)

func NewProperties(props types.Properties) (Properties, error) {
	obj := Properties{}

	err := errors.ChainUntilFail(
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
		func() (err error) {
			obj.Storage, err = NewStorageProperties(props)
			return err
		},
		func() (err error) {
			for i := 0; i < maxNetworkInterfacePropertiesArrayCapacity; i++ {
				propName := fmt.Sprintf("net%d", i)
				prop, ok := props[propName]
				if !ok {
					continue
				}

				x, ok := prop.(string)
				if !ok {
					err := errors.ErrInvalidProperty
					err.AddKey("name", propName)
					err.AddKey("value", prop)
					return err
				}

				if network, err := NewNetworkInterfaceProperties(x); err == nil {
					obj.Network = append(obj.Network, network)
				} else {
					return err
				}
			}

			return nil
		},
	)

	return obj, err
}

type GlobalProperties struct {
	OSType OSType

	Protected bool

	StartOnBoot bool
}

const (
	mkGlobalPropertyOSType      = "ostype"
	mkGlobalPropertyProtected   = "protection"
	mkGlobalPropertyStartOnBoot = "onboot"

	DefaultGlobalPropertyProtected   bool = false
	DefaultGlobalPropertyStartOnBoot bool = false
)

func NewGlobalProperties(
	props types.Properties,
) (GlobalProperties, error) {
	obj := GlobalProperties{}

	err := errors.ChainUntilFail(
		func() error {
			return props.SetRequiredFixedValue(
				mkGlobalPropertyOSType,
				&obj.OSType,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkGlobalPropertyProtected,
				&obj.Protected,
				DefaultGlobalPropertyProtected,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkGlobalPropertyStartOnBoot,
				&obj.StartOnBoot,
				DefaultGlobalPropertyStartOnBoot,
				nil,
			)
		},
	)

	return obj, err
}

func (obj Properties) MapToValues() (request.Values, error) {
	values := request.Values{}

	cpuValues, err := obj.CPU.MapToValues()
	if err != nil {
		return nil, err
	} else {
		for k, v := range cpuValues {
			values[k] = v
		}
	}

	memoryValues, err := obj.Memory.MapToValues()
	if err != nil {
		return nil, err
	} else {
		for k, v := range memoryValues {
			values[k] = v
		}
	}

	return values, nil
}
