package vm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
	"github.com/xabinapal/gopve/test"
)

func TestStorageLXCCPUProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"cores":    16,
		"cpulimit": 2,
		"cpuunits": 2048,
	})

	requiredProps := []string{"cores"}

	defaultProps := []string{"cpulimit", "cpuunits"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := vm.NewLXCCPUProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			cpuProps, err := vm.NewLXCCPUProperties(props)
			require.NoError(t, err)

			assert.Equal(t, uint(16), cpuProps.Cores)
			assert.Equal(t, uint(2), cpuProps.Limit)
			assert.Equal(t, uint(2048), cpuProps.Units)
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
				require.IsType(t, (*vm.LXCCPUProperties)(nil), obj)

				cpuProps := obj.(*vm.LXCCPUProperties)

				assert.Equal(
					t,
					vm.DefaultLXCCPUPropertyLimit,
					cpuProps.Limit,
				)
				assert.Equal(
					t,
					vm.DefaultLXCCPUPropertyUnits,
					cpuProps.Units,
				)
			},
		),
	)
}

func TestStorageLXCMemoryProperties(t *testing.T) {
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
