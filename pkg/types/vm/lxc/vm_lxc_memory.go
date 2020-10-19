package lxc

import (
	"fmt"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
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
) (obj MemoryProperties, err error) {
	return obj, errors.ChainUntilFail(
		func() error {
			return props.SetRequiredUint(
				mkMemoryPropertyMemory,
				&obj.Memory,
				nil,
			)
		},
		func() error {
			return props.SetRequiredUint(mkMemoryPropertySwap, &obj.Swap, nil)
		},
	)
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
