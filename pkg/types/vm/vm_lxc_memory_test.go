package vm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
	"github.com/xabinapal/gopve/test"
)

func TestLXCMemoryProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"memory": 4096,
		"swap":   2048,
	})

	requiredProps := []string{"memory", "swap"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := vm.NewLXCMemoryProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			memoryProps, err := vm.NewLXCMemoryProperties(props)
			require.NoError(t, err)

			assert.Equal(t, uint(4096), memoryProps.Memory)
			assert.Equal(t, uint(2048), memoryProps.Swap)
		})

	t.Run(
		"RequiredProperties",
		test.HelperTestRequiredProperties(t, props, requiredProps, factoryFunc),
	)
}
