package vm

import (
	"fmt"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/lxc"
)

type LXCCPUProperties struct {
	Architecture lxc.CPUArchitecture

	Cores uint

	Limit uint
	Units uint
}

const (
	mkLXCCPUPropertyArchitecture = "arch"

	mkLXCCPUPropertyCores = "cores"
	mkLXCCPUPropertyLimit = "cpulimit"
	mkLXCCPUPropertyUnits = "cpuunits"

	DefaultLXCCPUPropertyLimit uint = 0
	DefaultLXCCPUPropertyUnits uint = 1024
)

func NewLXCCPUProperties(props types.Properties) (*LXCCPUProperties, error) {
	obj := new(LXCCPUProperties)

	if err := props.SetRequiredFixedValue(mkLXCCPUPropertyArchitecture, &obj.Architecture, nil); err != nil {
		return nil, err
	}

	if err := props.SetRequiredUint(mkLXCCPUPropertyCores, &obj.Cores, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= 128
		},
	}); err != nil {
		return nil, err
	}

	if err := props.SetUint(mkLXCCPUPropertyLimit, &obj.Limit, DefaultLXCCPUPropertyLimit, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= 128
		},
	}); err != nil {
		return nil, err
	}

	if err := props.SetUint(mkLXCCPUPropertyUnits, &obj.Units, DefaultLXCCPUPropertyUnits, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val >= 8 && val <= 500000
		},
	}); err != nil {
		return nil, err
	}

	return obj, nil
}

func (obj LXCCPUProperties) MapToValues() (request.Values, error) {
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
