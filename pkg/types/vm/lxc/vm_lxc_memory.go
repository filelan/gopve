package lxc

import (
	"fmt"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
)

type MemoryProperties struct {
	Memory uint
	Swap   uint
}

const (
	mkMemoryPropertyMemory = "memory"
	mkMemoryPropertySwap   = "swap"
)

func NewMemoryProperties(
	props types.Properties,
) (*MemoryProperties, error) {
	obj := new(MemoryProperties)

	if err := props.SetRequiredUint(mkMemoryPropertyMemory, &obj.Memory, nil); err != nil {
		return nil, err
	}

	if err := props.SetRequiredUint(mkMemoryPropertySwap, &obj.Swap, nil); err != nil {
		return nil, err
	}

	return obj, nil
}

func (obj MemoryProperties) MapToValues() (request.Values, error) {
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
