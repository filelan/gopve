package vm

import (
	"fmt"
	"strconv"

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
	CPU    QEMUCPUProperties
	Memory QEMUMemoryProperties
}

func NewQEMUProperties(props types.Properties) (*QEMUProperties, error) {
	obj := new(QEMUProperties)

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
	Type    string
	Sockets uint
	Cores   uint
	VCPUs   uint

	Limit uint
	Units uint

	NUMA bool

	FreezeAtStartup bool
}

const (
	mkQEMUCPUPropertySockets = "sockets"
	mkQEMUCPUPropertyCores   = "cores"
	mkQEMUCPUPropertyVCPUs   = "vcpus"
	mkQEMUCPUPropertLimit    = "cpulimit"
	mkQEMUCPUPropertyUnits   = "cpuunits"

	mkQEMUCPUPropertyNUMA            = "numa"
	mkQEMUCPUPropertyFreezeAtStartup = "freeze"

	DefaultQEMUCPUPropertyLimit uint = 0
	DefaultQEMUCPUPropertyUnits uint = 1024

	DefaultQEMUCPUPropertyNUMA            bool = false
	DefaultQEMUCPUPropertyFreezeAtStartup bool = false
)

func NewQEMUCPUProperties(props types.Properties) (*QEMUCPUProperties, error) {
	obj := new(QEMUCPUProperties)

	if v, ok := props[mkQEMUCPUPropertySockets].(float64); ok {
		if v != float64(int(v)) || v < 0 || v > 4 {
			err := errors.ErrInvalidProperty
			err.AddKey("name", mkQEMUCPUPropertySockets)
			err.AddKey("value", v)
			return nil, err
		}

		obj.Sockets = uint(v)
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", mkQEMUCPUPropertySockets)
		return nil, err
	}

	if v, ok := props[mkQEMUCPUPropertyCores].(float64); ok {
		if v != float64(int(v)) || v < 0 || v > 128 {
			err := errors.ErrInvalidProperty
			err.AddKey("name", mkQEMUCPUPropertyCores)
			err.AddKey("value", v)
			return nil, err
		}

		obj.Cores = uint(v)
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", mkQEMUCPUPropertyCores)
		return nil, err
	}

	if v, ok := props[mkQEMUCPUPropertyVCPUs].(float64); ok {
		if v != float64(int(v)) || v < 0 || uint(v) > obj.Sockets*obj.Cores {
			err := errors.ErrInvalidProperty
			err.AddKey("name", mkQEMUCPUPropertyVCPUs)
			err.AddKey("value", v)
			return nil, err
		}

		obj.VCPUs = uint(v)
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", mkQEMUCPUPropertyVCPUs)
		return nil, err
	}

	if v, ok := props[mkQEMUCPUPropertLimit].(string); ok {
		if v == "" {
			obj.Limit = DefaultQEMUCPUPropertyLimit
		} else {
			limit, err := strconv.Atoi(v)
			if err != nil || limit < 0 || limit > 128 {
				err := errors.ErrInvalidProperty
				err.AddKey("name", mkQEMUCPUPropertLimit)
				err.AddKey("value", v)
				return nil, err
			} else {
				obj.Limit = uint(limit)
			}
		}
	} else {
		obj.Limit = DefaultQEMUCPUPropertyLimit
	}

	if v, ok := props[mkQEMUCPUPropertyUnits].(float64); ok {
		if v != float64(int(v)) || v < 8 || v > 500000 {
			err := errors.ErrInvalidProperty
			err.AddKey("name", mkQEMUCPUPropertyUnits)
			err.AddKey("value", v)
			return nil, err
		} else {
			obj.Units = uint(v)
		}
	} else {
		obj.Units = DefaultQEMUCPUPropertyUnits
	}

	if v, ok := props[mkQEMUCPUPropertyNUMA].(float64); ok {
		obj.NUMA = internal_types.NewPVEBoolFromFloat64(v).Bool()
	} else {
		obj.NUMA = DefaultQEMUCPUPropertyNUMA
	}

	if v, ok := props[mkQEMUCPUPropertyFreezeAtStartup].(float64); ok {
		obj.FreezeAtStartup = internal_types.NewPVEBoolFromFloat64(v).Bool()
	} else {
		obj.FreezeAtStartup = DefaultQEMUCPUPropertyFreezeAtStartup
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

	if v, ok := props[mkQEMUMemoryPropertyMemory].(float64); ok {
		if v != float64(int(v)) || v < 0 || v > 4178944 {
			err := errors.ErrInvalidProperty
			err.AddKey("name", mkQEMUMemoryPropertyMemory)
			err.AddKey("value", v)
			return nil, err
		}

		obj.Memory = uint(v)
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", mkQEMUMemoryPropertyMemory)
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
