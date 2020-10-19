package lxc

import (
	"fmt"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
	"github.com/xabinapal/gopve/pkg/types/schema"
)

type CPUProperties struct {
	Architecture CPUArchitecture

	Cores uint

	Limit uint
	Units uint
}

const (
	mkCPUPropertyArchitecture = "arch"

	mkCPUPropertyCores = "cores"
	mkCPUPropertyLimit = "cpulimit"
	mkCPUPropertyUnits = "cpuunits"

	DefaultCPUPropertyLimit uint = 0
	DefaultCPUPropertyUnits uint = 1024
)

func NewCPUProperties(props types.Properties) (obj CPUProperties, err error) {
	return obj, errors.ChainUntilFail(
		func() error {
			return props.SetRequiredFixedValue(
				mkCPUPropertyArchitecture,
				&obj.Architecture,
				nil,
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
	)
}

func (obj CPUProperties) MapToValues() (request.Values, error) {
	values := request.Values{}

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
		return nil, fmt.Errorf(
			"Invalid CPU units, must be between 2 and 500000",
		)
	} else if obj.Units != 0 {
		values.AddUint("cpuunits", obj.Units)
	}

	return values, nil
}
