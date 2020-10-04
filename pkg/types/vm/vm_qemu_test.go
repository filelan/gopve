package vm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
	"github.com/xabinapal/gopve/test"
)

func TestStorageQEMUCPUProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"sockets":  2,
		"cores":    16,
		"vcpus":    8,
		"cpulimit": "2",
		"cpuunits": 2048,
		"numa":     1,
		"freeze":   1,
	})

	requiredProps := []string{"sockets", "cores", "vcpus"}

	defaultProps := []string{"cpulimit", "cpuunits", "numa", "freeze"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := vm.NewQEMUCPUProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			cpuProps, err := vm.NewQEMUCPUProperties(props)
			require.NoError(t, err)

			assert.Equal(t, uint(2), cpuProps.Sockets)
			assert.Equal(t, uint(16), cpuProps.Cores)
			assert.Equal(t, uint(8), cpuProps.VCPUs)
			assert.Equal(t, uint(2), cpuProps.Limit)
			assert.Equal(t, uint(2048), cpuProps.Units)
			assert.Equal(t, true, cpuProps.NUMA)
			assert.Equal(t, true, cpuProps.FreezeAtStartup)
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
				require.IsType(t, (*vm.QEMUCPUProperties)(nil), obj)

				cpuProps := obj.(*vm.QEMUCPUProperties)

				assert.Equal(
					t,
					vm.DefaultQEMUCPUPropertyLimit,
					cpuProps.Limit,
				)
				assert.Equal(
					t,
					vm.DefaultQEMUCPUPropertyUnits,
					cpuProps.Units,
				)
				assert.Equal(
					t,
					vm.DefaultQEMUCPUPropertyNUMA,
					cpuProps.NUMA,
				)
				assert.Equal(
					t,
					vm.DefaultQEMUCPUPropertyFreezeAtStartup,
					cpuProps.FreezeAtStartup,
				)
			},
		),
	)
}

func TestStorageQEMUMemoryProperties(t *testing.T) {
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
				require.IsType(t, (*vm.QEMUMemoryProperties)(nil), obj)

				memoryProps := obj.(*vm.QEMUMemoryProperties)

				assert.Equal(
					t,
					vm.DefaultQEMUMemoryShares,
					memoryProps.Shares,
				)
			},
		),
	)
}
