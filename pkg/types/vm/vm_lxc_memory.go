package vm

import (
	"fmt"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
)

type LXCMemoryProperties struct {
	Memory uint
	Swap   uint
}

const (
	mkLXCMemoryPropertyMemory = "memory"
	mkLXCMemoryPropertySwap   = "swap"
)

func NewLXCMemoryProperties(
	props types.Properties,
) (*LXCMemoryProperties, error) {
	obj := new(LXCMemoryProperties)

	if err := props.SetRequiredUint(mkQEMUMemoryPropertyMemory, &obj.Memory, nil); err != nil {
		return nil, err
	}

	if err := props.SetRequiredUint(mkLXCMemoryPropertySwap, &obj.Swap, nil); err != nil {
		return nil, err
	}

	return obj, nil
}

func (obj LXCMemoryProperties) MapToValues() (request.Values, error) {
	values := request.Values{}

	memory := obj.Memory
	if memory == 0 {
		memory = 512
	} else if memory < 16 {
		return nil, fmt.Errorf("Invalid memory, must at least 16")
	}
	values.AddUint("memory", memory)

	values.AddUint("swap", obj.Swap)

	return values, nil
}
