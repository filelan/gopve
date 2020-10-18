package qemu

import (
	"fmt"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
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

	if err := props.SetRequiredUint(mkMemoryPropertyMemory, &obj.Memory, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= 4178944
		},
	}); err != nil {
		return obj, err
	}

	if v, ok := props[mkMemoryPropertyBalloon].(float64); ok {
		if v != float64(int(v)) || v < 0 || uint(v) > obj.Memory {
			err := errors.ErrInvalidProperty
			err.AddKey("name", mkMemoryPropertyBalloon)
			err.AddKey("value", v)
			return obj, err
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

			if v, ok := props[mkMemoryPropertyShares].(float64); ok {
				if v != float64(int(v)) || v < 0 || v > 50000 {
					err := errors.ErrInvalidProperty
					err.AddKey("name", mkMemoryPropertyShares)
					err.AddKey("value", v)
					return obj, err
				}

				obj.Shares = uint(v)
			} else {
				obj.Shares = DefaultMemoryShares
			}
		}
	} else {
		obj.Ballooning = true
		obj.MinimumMemory = obj.Memory
		obj.Shares = 0
	}

	return obj, nil
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
