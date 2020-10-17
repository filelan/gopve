package vm_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
	"github.com/xabinapal/gopve/test"
)

func TestStorageQEMUCPUProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"cpu":      "pentium,flags=+md-clear;-spec-ctrl",
		"arch":     "aarch64",
		"sockets":  2,
		"cores":    16,
		"vcpus":    8,
		"cpulimit": "2",
		"cpuunits": 2048,
		"numa":     1,
		"freeze":   1,
	})

	requiredProps := []string{"sockets", "cores"}

	defaultProps := []string{
		"cpu",
		"arch",
		"vcpus",
		"cpulimit",
		"cpuunits",
		"numa",
		"freeze",
	}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := vm.NewQEMUCPUProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			cpuProps, err := vm.NewQEMUCPUProperties(props)
			require.NoError(t, err)

			assert.Equal(t, qemu.CPUTypeIntelPentium, cpuProps.Kind)
			assert.Equal(
				t,
				qemu.CPUArchitectureARM64,
				cpuProps.Architecture,
			)
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
				require.IsType(t, vm.QEMUCPUProperties{}, obj)

				cpuProps := obj.(vm.QEMUCPUProperties)
				fmt.Printf("%+v\n", cpuProps)

				assert.Equal(
					t,
					vm.DefaultQEMUCPUPropertyKind,
					cpuProps.Kind,
				)
				assert.Equal(
					t,
					vm.DefaultQEMUCPUPropertyArchitecture,
					cpuProps.Architecture,
				)

				fmt.Printf("%+v\n", cpuProps)
				assert.Equal(
					t,
					cpuProps.Sockets*cpuProps.Cores,
					cpuProps.VCPUs,
				)
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
