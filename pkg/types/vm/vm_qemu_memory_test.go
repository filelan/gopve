package vm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
	"github.com/xabinapal/gopve/test"
)

func TestQEMUMemoryProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"memory":  4096,
		"balloon": 2048,
		"shares":  512,
	})

	requiredProps := []string{"memory"}

	defaultProps := []string{"shares"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := vm.NewQEMUMemoryProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			memoryProps, err := vm.NewQEMUMemoryProperties(props)
			require.NoError(t, err)

			assert.Equal(t, uint(4096), memoryProps.Memory)
			assert.Equal(t, true, memoryProps.Ballooning)
			assert.Equal(t, uint(2048), memoryProps.MinimumMemory)
			assert.Equal(t, uint(512), memoryProps.Shares)
		})

	t.Run(
		"RequiredProperties",
		test.HelperTestRequiredProperties(t, props, requiredProps, factoryFunc),
	)

	t.Run(
		"DefaultProperties",
		test.HelperTestOptionalProperties(
			t,
			props,
			defaultProps,
			factoryFunc,
			func(obj interface{}) {
				require.IsType(t, vm.QEMUMemoryProperties{}, obj)

				memoryProps := obj.(vm.QEMUMemoryProperties)

				assert.Equal(
					t,
					vm.DefaultQEMUMemoryShares,
					memoryProps.Shares,
				)
			},
		),
	)
}
