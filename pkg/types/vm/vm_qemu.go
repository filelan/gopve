package vm

import (
	"fmt"

	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
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
	CPU    QEMUCPUProperties
	Memory QEMUMemoryProperties
}

func NewQEMUProperties(props types.Properties) (*QEMUProperties, error) {
	obj := new(QEMUProperties)

	if v, err := NewQEMUGlobalProperties(props); err != nil {
		return nil, err
	} else {
		obj.QEMUGlobalProperties = *v
	}

	if v, err := NewQEMUCPUProperties(props); err != nil {
		return nil, err
	} else {
		obj.CPU = *v
	}

	if v, err := NewQEMUMemoryProperties(props); err != nil {
		return nil, err
	} else {
		obj.Memory = *v
	}

	return obj, nil
}

type QEMUGlobalProperties struct {
	OSType QEMUOSType

	Protected bool

	StartAtBoot bool
}

const (
	mkQEMUGlobalPropertyOSType      = "ostype"
	mkQEMUGlobalPropertyProtected   = "protection"
	mkQEMUGlobalPropertyStartAtBoot = "onboot"

	DefaultQEMUGlobalPropertyProtected   bool = false
	DefaultQEMUGlobalPropertyStartAtBoot bool = false
)

func NewQEMUGlobalProperties(props types.Properties) (*QEMUGlobalProperties, error) {
	obj := new(QEMUGlobalProperties)

	if err := props.SetRequiredFixedValue(mkQEMUGlobalPropertyOSType, &obj.OSType, nil); err != nil {
		return nil, err
	}

	if err := props.SetBool(mkQEMUGlobalPropertyProtected, &obj.Protected, DefaultQEMUGlobalPropertyProtected, nil); err != nil {
		return nil, err
	}

	if err := props.SetBool(mkQEMUGlobalPropertyStartAtBoot, &obj.StartAtBoot, DefaultQEMUGlobalPropertyStartAtBoot, nil); err != nil {
		return nil, err
	}

	return obj, nil
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

type QEMUCPUProperties struct {
	Kind  QEMUCPUKind
	Flags []QEMUCPUFlags

	Architecture QEMUCPUArchitecture

	Sockets uint
	Cores   uint
	VCPUs   uint

	Limit uint
	Units uint

	NUMA bool

	FreezeAtStartup bool
}

const (
	mkQEMUCPUPropertyCPU   = "cpu"
	mkQEMUCPUPropertyFlags = "flags"

	mkQEMUCPUPropertyArchitecture = "arch"

	mkQEMUCPUPropertySockets = "sockets"
	mkQEMUCPUPropertyCores   = "cores"
	mkQEMUCPUPropertyVCPUs   = "vcpus"
	mkQEMUCPUPropertyLimit   = "cpulimit"
	mkQEMUCPUPropertyUnits   = "cpuunits"

	mkQEMUCPUPropertyNUMA            = "numa"
	mkQEMUCPUPropertyFreezeAtStartup = "freeze"

	DefaultQEMUCPUPropertyKind         QEMUCPUKind         = QEMUCPUKindKVM64
	DefaultQEMUCPUPropertyArchitecture QEMUCPUArchitecture = QEMUCPUArchitectureHost

	DefaultQEMUCPUPropertyLimit uint = 0
	DefaultQEMUCPUPropertyUnits uint = 1024

	DefaultQEMUCPUPropertyNUMA            bool = false
	DefaultQEMUCPUPropertyFreezeAtStartup bool = false
)

func NewQEMUCPUProperties(props types.Properties) (*QEMUCPUProperties, error) {
	obj := new(QEMUCPUProperties)

	if err := props.SetFixedValue(mkQEMUCPUPropertyCPU, &obj.Kind, DefaultQEMUCPUPropertyKind, nil); err != nil {
		return nil, err
	}

	obj.Flags = []QEMUCPUFlags{}
	if v, ok := props[mkQEMUCPUPropertyCPU].(string); ok {
		cpuOptions := internal_types.PVEDictionary{
			ListSeparator:     ",",
			KeyValueSeparator: "=",
			AllowNoValue:      true,
		}

		if err := (&cpuOptions).Unmarshal(v); err != nil {
			err := errors.ErrInvalidProperty
			err.AddKey("name", mkQEMUCPUPropertyCPU)
			err.AddKey("value", v)
			return nil, err
		}

		for _, vv := range cpuOptions.List() {
			if !vv.HasValue() {
				if err := (&obj.Kind).Unmarshal(vv.Key()); err != nil {
					err := errors.ErrInvalidProperty
					err.AddKey("name", mkQEMUCPUPropertyCPU)
					err.AddKey("value", v)
					return nil, err
				}
			} else {
				switch vv.Key() {
				case mkQEMUCPUPropertyFlags:
					flags := internal_types.PVEList{
						Separator: ";",
					}

					if err := (&flags).Unmarshal(vv.Value()); err != nil {
						err := errors.ErrInvalidProperty
						err.AddKey("name", mkQEMUCPUPropertyCPU)
						err.AddKey("value", v)
						return nil, err
					}

					for _, v := range flags.List() {
						var flag QEMUCPUFlags
						if err := (&flag).Unmarshal(v); err != nil {
							err := errors.ErrInvalidProperty
							err.AddKey("name", mkQEMUCPUPropertyCPU)
							err.AddKey("value", v)
							return nil, err
						}

						obj.Flags = append(obj.Flags, flag)
					}
				}
			}
		}
	} else {
		obj.Kind = DefaultQEMUCPUPropertyKind
	}

	if err := props.SetFixedValue(mkQEMUCPUPropertyArchitecture, &obj.Architecture, DefaultQEMUCPUPropertyArchitecture, nil); err != nil {
		return nil, err
	}

	if err := props.SetRequiredUint(mkQEMUCPUPropertySockets, &obj.Sockets, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= 4
		},
	}); err != nil {
		return nil, err
	}

	if err := props.SetRequiredUint(mkQEMUCPUPropertyCores, &obj.Cores, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= 128
		},
	}); err != nil {
		return nil, err
	}

	if err := props.SetRequiredUint(mkQEMUCPUPropertyVCPUs, &obj.VCPUs, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= obj.Sockets*obj.Cores
		},
	}); err != nil {
		return nil, err
	}

	if err := props.SetUintFromString(mkQEMUCPUPropertyLimit, &obj.Limit, DefaultQEMUCPUPropertyLimit, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= 128
		},
	}); err != nil {
		return nil, err
	}

	if err := props.SetUint(mkQEMUCPUPropertyUnits, &obj.Units, DefaultQEMUCPUPropertyUnits, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val >= 8 && val <= 500000
		},
	}); err != nil {
		return nil, err
	}

	if err := props.SetBool(mkQEMUCPUPropertyNUMA, &obj.NUMA, DefaultQEMUCPUPropertyNUMA, nil); err != nil {
		return nil, err
	}

	if err := props.SetBool(mkQEMUCPUPropertyFreezeAtStartup, &obj.FreezeAtStartup, DefaultQEMUCPUPropertyFreezeAtStartup, nil); err != nil {
		return nil, err
	}

	return obj, nil
}

func (obj QEMUCPUProperties) MapToValues() (request.Values, error) {
	values := request.Values{}

	sockets := obj.Sockets
	if sockets == 0 {
		sockets = 1
	} else if sockets > 4 {
		return nil, fmt.Errorf("Invalid CPU sockets, the maximum allowed is 4")
	}
	values.AddUint("sockets", sockets)

	cores := obj.Cores
	if cores == 0 {
		cores = 1
	} else if cores > 128 {
		return nil, fmt.Errorf("Invalid CPU cores, the maximum allowed is 128")
	}
	values.AddUint("cores", cores)

	if obj.VCPUs != 0 && (obj.VCPUs > sockets*cores) {
		return nil, fmt.Errorf(
			"Invalid CPU hotplugged cores, can't be greater than sockets * cores",
		)
	} else if obj.VCPUs != 0 {
		values.AddUint("vcpus", obj.VCPUs)
	}

	if obj.Limit > 128 {
		return nil, fmt.Errorf("Invalid CPU limit, must be between 0 and 128")
	} else if obj.Limit != 0 {
		values.AddUint("cpulimit", obj.Limit)
	}

	if obj.Units != 0 && (obj.Units < 2 || obj.Units > 262144) {
		return nil, fmt.Errorf(
			"Invalid CPU units, must be between 2 and 262144",
		)
	} else if obj.Units != 0 && obj.Units != 1024 {
		values.AddUint("cpuunits", obj.Units)
	}

	values.AddBool("numa", obj.NUMA)

	values.AddBool("freeze", obj.FreezeAtStartup)

	return values, nil
}

type QEMUMemoryProperties struct {
	Memory uint

	Ballooning    bool
	MinimumMemory uint
	Shares        uint
}

const (
	mkQEMUMemoryPropertyMemory  = "memory"
	mkQEMUMemoryPropertyBalloon = "balloon"
	mkQEMUMemoryPropertyShares  = "shares"

	DefaultQEMUMemoryShares uint = 1000
)

func NewQEMUMemoryProperties(
	props types.Properties,
) (*QEMUMemoryProperties, error) {
	obj := new(QEMUMemoryProperties)

	if err := props.SetRequiredUint(mkQEMUMemoryPropertyMemory, &obj.Memory, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= 4178944
		},
	}); err != nil {
		return nil, err
	}

	if v, ok := props[mkQEMUMemoryPropertyBalloon].(float64); ok {
		if v != float64(int(v)) || v < 0 || uint(v) > obj.Memory {
			err := errors.ErrInvalidProperty
			err.AddKey("name", mkQEMUMemoryPropertyBalloon)
			err.AddKey("value", v)
			return nil, err
		} else if v == 0 {
			obj.Ballooning = false
			obj.MinimumMemory = obj.Memory
			obj.Shares = 0
		} else if uint(v) == obj.Memory {
			obj.Ballooning = true
			obj.MinimumMemory = obj.Memory
			obj.Shares = 0
		} else {
			obj.Ballooning = true
			obj.MinimumMemory = uint(v)

			if v, ok := props[mkQEMUMemoryPropertyShares].(float64); ok {
				if v != float64(int(v)) || v < 0 || v > 50000 {
					err := errors.ErrInvalidProperty
					err.AddKey("name", mkQEMUMemoryPropertyShares)
					err.AddKey("value", v)
					return nil, err
				}

				obj.Shares = uint(v)
			} else {
				obj.Shares = DefaultQEMUMemoryShares
			}
		}
	} else {
		obj.Ballooning = true
		obj.MinimumMemory = obj.Memory
		obj.Shares = 0
	}

	return obj, nil
}

func (obj *QEMUMemoryProperties) MapToValues() (request.Values, error) {
	values := request.Values{}

	memory := obj.Memory
	if memory == 0 {
		memory = 512
	} else if memory < 16 || memory > 4178944 {
		return nil, fmt.Errorf("Invalid memory, must be between 16 and 4178944")
	}
	values.AddUint("memory", memory)

	if obj.Ballooning {
		minimumMemory := obj.MinimumMemory
		if minimumMemory == 0 {
			minimumMemory = memory
		} else if minimumMemory > memory {
			return nil, fmt.Errorf("Invalid memory ballooning minimum, can't be greater than total memory")
		}

		values.AddUint("balloon", minimumMemory)

		if minimumMemory == memory {
			values.AddUint("shares", 0)
		} else {
			if obj.Shares == 0 {
				values.AddUint("shares", 1000)
			} else {
				values.AddUint("shares", obj.Shares)
			}
		}
	} else {
		values.AddUint("balloon", 0)
		values.AddUint("shares", 0)
	}

	return values, nil
}
