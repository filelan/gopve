package vm

import (
	"fmt"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
)

type QEMUVirtualMachine interface {
	VirtualMachine

	CPU() (QEMUCPUProperties, error)
	Memory() (QEMUMemoryProperties, error)

	GetQEMUProperties() (QEMUProperties, error)
	SetQEMUProperties(props QEMUProperties) error
}

type QEMUCreateOptions struct {
	VMID uint
	Node string

	Properties QEMUProperties
}

func (obj QEMUCreateOptions) MapToValues() (request.Values, error) {
	values, err := obj.Properties.MapToValues()
	if err != nil {
		return nil, err
	}

	return values, nil
}

type QEMUProperties struct {
	QEMUGlobalProperties
	CPU     QEMUCPUProperties
	Memory  QEMUMemoryProperties
	Storage QEMUStorageProperties
	Network []QEMUNetworkInterfaceProperties
}

const (
	maxQEMUNetworkInterfacePropertiesArrayCapacity = 32
)

func NewQEMUProperties(props types.Properties) (QEMUProperties, error) {
	obj := QEMUProperties{}

	err := errors.ChainUntilFail(
		func() (err error) {
			obj.QEMUGlobalProperties, err = NewQEMUGlobalProperties(props)
			return err
		},
		func() (err error) {
			obj.CPU, err = NewQEMUCPUProperties(props)
			return err
		},
		func() (err error) {
			obj.Memory, err = NewQEMUMemoryProperties(props)
			return err
		},
		func() (err error) {
			obj.Storage, err = NewQEMUStorageProperties(props)
			return err
		},
		func() (err error) {
			for i := 0; i < maxQEMUNetworkInterfacePropertiesArrayCapacity; i++ {
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

				if network, err := NewQEMUNetworkInterfaceProperties(x); err == nil {
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

type QEMUGlobalProperties struct {
	OSType qemu.OSType

	Protected bool

	StartOnBoot bool
}

const (
	mkQEMUGlobalPropertyOSType      = "ostype"
	mkQEMUGlobalPropertyProtected   = "protection"
	mkQEMUGlobalPropertyStartOnBoot = "onboot"

	DefaultQEMUGlobalPropertyProtected   bool = false
	DefaultQEMUGlobalPropertyStartOnBoot bool = false
)

func NewQEMUGlobalProperties(
	props types.Properties,
) (QEMUGlobalProperties, error) {
	obj := QEMUGlobalProperties{}

	err := errors.ChainUntilFail(
		func() error {
			return props.SetRequiredFixedValue(
				mkQEMUGlobalPropertyOSType,
				&obj.OSType,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkQEMUGlobalPropertyProtected,
				&obj.Protected,
				DefaultQEMUGlobalPropertyProtected,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkQEMUGlobalPropertyStartOnBoot,
				&obj.StartOnBoot,
				DefaultQEMUGlobalPropertyStartOnBoot,
				nil,
			)
		},
	)

	return obj, err
}

func (obj QEMUProperties) MapToValues() (request.Values, error) {
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
