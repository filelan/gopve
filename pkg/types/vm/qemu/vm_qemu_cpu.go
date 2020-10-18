package qemu

import (
	"fmt"

	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type CPUProperties struct {
	Kind  CPUType
	Flags []CPUFlags

	Architecture CPUArchitecture

	Sockets uint
	Cores   uint
	VCPUs   uint

	Limit uint
	Units uint

	NUMA bool

	FreezeAtStartup bool
}

const (
	mkCPUPropertyCPU   = "cpu"
	mkCPUPropertyFlags = "flags"

	mkCPUPropertyArchitecture = "arch"

	mkCPUPropertySockets = "sockets"
	mkCPUPropertyCores   = "cores"
	mkCPUPropertyVCPUs   = "vcpus"
	mkCPUPropertyLimit   = "cpulimit"
	mkCPUPropertyUnits   = "cpuunits"

	mkCPUPropertyNUMA            = "numa"
	mkCPUPropertyFreezeAtStartup = "freeze"

	DefaultCPUPropertyKind         CPUType         = CPUTypeKVM64
	DefaultCPUPropertyArchitecture CPUArchitecture = CPUArchitectureHost

	DefaultCPUPropertyLimit uint = 0
	DefaultCPUPropertyUnits uint = 1024

	DefaultCPUPropertyNUMA            bool = false
	DefaultCPUPropertyFreezeAtStartup bool = false
)

func NewCPUProperties(props types.Properties) (CPUProperties, error) {
	obj := CPUProperties{}

	if err := props.SetFixedValue(mkCPUPropertyCPU, &obj.Kind, DefaultCPUPropertyKind, nil); err != nil {
		return obj, err
	}

	obj.Flags = []CPUFlags{}
	if v, ok := props[mkCPUPropertyCPU].(string); ok {
		cpuOptions := internal_types.PVEDictionary{
			ListSeparator:     ",",
			KeyValueSeparator: "=",
			AllowNoValue:      true,
		}

		if err := (&cpuOptions).Unmarshal(v); err != nil {
			err := errors.ErrInvalidProperty
			err.AddKey("name", mkCPUPropertyCPU)
			err.AddKey("value", v)
			return obj, err
		}

		for _, vv := range cpuOptions.List() {
			if !vv.HasValue() {
				if err := (&obj.Kind).Unmarshal(vv.Key()); err != nil {
					err := errors.ErrInvalidProperty
					err.AddKey("name", mkCPUPropertyCPU)
					err.AddKey("value", v)
					return obj, err
				}
			} else {
				switch vv.Key() {
				case mkCPUPropertyFlags:
					flags := internal_types.PVEList{
						Separator: ";",
					}

					if err := (&flags).Unmarshal(vv.Value()); err != nil {
						err := errors.ErrInvalidProperty
						err.AddKey("name", mkCPUPropertyCPU)
						err.AddKey("value", v)
						return obj, err
					}

					for _, v := range flags.List() {
						var flag CPUFlags
						if err := (&flag).Unmarshal(v); err != nil {
							err := errors.ErrInvalidProperty
							err.AddKey("name", mkCPUPropertyCPU)
							err.AddKey("value", v)
							return obj, err
						}

						obj.Flags = append(obj.Flags, flag)
					}
				}
			}
		}
	} else {
		obj.Kind = DefaultCPUPropertyKind
	}

	if err := props.SetFixedValue(mkCPUPropertyArchitecture, &obj.Architecture, DefaultCPUPropertyArchitecture, nil); err != nil {
		return obj, err
	}

	if err := props.SetRequiredUint(mkCPUPropertySockets, &obj.Sockets, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= 4
		},
	}); err != nil {
		return obj, err
	}

	if err := props.SetRequiredUint(mkCPUPropertyCores, &obj.Cores, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= 128
		},
	}); err != nil {
		return obj, err
	}

	if err := props.SetUint(mkCPUPropertyVCPUs, &obj.VCPUs, obj.Sockets*obj.Cores, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= obj.Sockets*obj.Cores
		},
	}); err != nil {
		return obj, err
	}

	if err := props.SetUintFromString(mkCPUPropertyLimit, &obj.Limit, DefaultCPUPropertyLimit, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= 128
		},
	}); err != nil {
		return obj, err
	}

	if err := props.SetUint(mkCPUPropertyUnits, &obj.Units, DefaultCPUPropertyUnits, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val >= 8 && val <= 500000
		},
	}); err != nil {
		return obj, err
	}

	if err := props.SetBool(mkCPUPropertyNUMA, &obj.NUMA, DefaultCPUPropertyNUMA, nil); err != nil {
		return obj, err
	}

	if err := props.SetBool(mkCPUPropertyFreezeAtStartup, &obj.FreezeAtStartup, DefaultCPUPropertyFreezeAtStartup, nil); err != nil {
		return obj, err
	}

	return obj, nil
}

func (obj CPUProperties) MapToValues() (request.Values, error) {
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
