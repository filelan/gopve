package qemu

import (
	"fmt"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
	"github.com/xabinapal/gopve/pkg/types/schema"
)

type MemoryProperties struct {
	Memory uint

	Ballooning    bool
	MinimumMemory uint
	Shares        uint
}

const (
	mkMemoryPropertyMemory  = "memory"
	mkMemoryPropertyBalloon = "balloon"
	mkMemoryPropertyShares  = "shares"

	DefaultMemoryShares uint = 1000
)

func NewMemoryProperties(
	props types.Properties,
) (MemoryProperties, error) {
	obj := MemoryProperties{}

	return obj, errors.ChainUntilFail(
		func() error {
			return props.SetRequiredUint(
				mkMemoryPropertyMemory,
				&obj.Memory,
				&schema.UintFunctions{
					ValidateFunc: func(val uint) bool {
						return val <= 4178944
					},
				},
			)
		},
		func() error {
			var balloon uint

			err := props.SetRequiredUint(
				mkMemoryPropertyBalloon,
				&balloon,
				&schema.UintFunctions{
					ValidateFunc: func(v uint) bool {
						return v <= obj.Memory
					},
				},
			)

			if errors.ErrMissingProperty.IsBase(err) {
				obj.Ballooning = true
				obj.MinimumMemory = obj.Memory
				obj.Shares = 0
				return nil
			}

			if err != nil {
				return err
			}

			switch balloon {
			case 0:
				obj.Ballooning = false
				obj.MinimumMemory = obj.Memory
				obj.Shares = 0
			case obj.Memory:
				obj.Ballooning = true
				obj.MinimumMemory = obj.Memory
				obj.Shares = 0
			default:
				obj.Ballooning = true
				obj.MinimumMemory = balloon

				return props.SetUint(
					mkMemoryPropertyShares,
					&obj.Shares,
					DefaultMemoryShares,
					&schema.UintFunctions{
						ValidateFunc: func(v uint) bool {
							return v <= 50000
						},
					},
				)
			}

			return nil
		},
	)
}

func (obj MemoryProperties) MapToValues() (request.Values, error) {
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
