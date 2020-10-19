package qemu

import (
	"fmt"

	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
	"github.com/xabinapal/gopve/pkg/types/schema"
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
	mkCPUDictPropertyCPU  = "cpu"
	mkCPUKeyPropertyFlags = "flags"

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

	return obj, errors.ChainUntilFail(
		func() error {
			obj.Flags = []CPUFlags{}
			cpuOptions, err := props.GetAsDict(
				mkCPUDictPropertyCPU,
				",",
				"=",
				true,
			)
			if err != nil {
				if errors.ErrMissingProperty.IsBase(err) {
					obj.Kind = DefaultCPUPropertyKind
					return nil
				}

				return err
			}

			for _, vv := range cpuOptions.List() {
				if !vv.HasValue() {
					if err := (&obj.Kind).Unmarshal(vv.Key()); err != nil {
						return errors.NewErrInvalidProperty(
							mkCPUDictPropertyCPU,
							vv.Key(),
						)
					}
				} else {
					switch vv.Key() {
					case mkCPUKeyPropertyFlags:
						flags := internal_types.PVEList{
							Separator: ";",
						}

						if err := (&flags).Unmarshal(vv.Value()); err != nil {
							return errors.NewErrInvalidProperty(mkCPUKeyPropertyFlags, vv.Value())
						}

						for _, v := range flags.List() {
							var flag CPUFlags
							if err := (&flag).Unmarshal(v); err != nil {
								return errors.NewErrInvalidProperty(mkCPUKeyPropertyFlags, vv.Value())
							}

							obj.Flags = append(obj.Flags, flag)
						}
					}
				}
			}

			return nil
		},
		func() error {
			return props.SetFixedValue(
				mkCPUPropertyArchitecture,
				&obj.Architecture,
				DefaultCPUPropertyArchitecture,
				nil,
			)
		},
		func() error {
			return props.SetRequiredUint(
				mkCPUPropertySockets,
				&obj.Sockets,
				&schema.UintFunctions{
					ValidateFunc: func(val uint) bool {
						return val <= 4
					},
				},
			)
		},
		func() error {
			return props.SetRequiredUint(
				mkCPUPropertyCores,
				&obj.Cores,
				&schema.UintFunctions{
					ValidateFunc: func(val uint) bool {
						return val <= 128
					},
				},
			)
		},
		func() error {
			return props.SetUint(
				mkCPUPropertyVCPUs,
				&obj.VCPUs,
				obj.Sockets*obj.Cores,
				&schema.UintFunctions{
					ValidateFunc: func(val uint) bool {
						return val <= obj.Sockets*obj.Cores
					},
				},
			)
		},
		func() error {
			return props.SetUintFromString(
				mkCPUPropertyLimit,
				&obj.Limit,
				DefaultCPUPropertyLimit,
				&schema.UintFunctions{
					ValidateFunc: func(val uint) bool {
						return val <= 128
					},
				},
			)
		},
		func() error {
			return props.SetUint(
				mkCPUPropertyUnits,
				&obj.Units,
				DefaultCPUPropertyUnits,
				&schema.UintFunctions{
					ValidateFunc: func(val uint) bool {
						return val >= 8 && val <= 500000
					},
				},
			)
		},
		func() error {
			return props.SetBool(
				mkCPUPropertyNUMA,
				&obj.NUMA,
				DefaultCPUPropertyNUMA,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkCPUPropertyFreezeAtStartup,
				&obj.FreezeAtStartup,
				DefaultCPUPropertyFreezeAtStartup,
				nil,
			)
		},
	)
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
