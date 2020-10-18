package qemu_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
	"github.com/xabinapal/gopve/test"
)

func TestStorageCPUProperties(t *testing.T) {
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
		obj, err := qemu.NewCPUProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			cpuProps, err := qemu.NewCPUProperties(props)
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
				require.IsType(t, qemu.CPUProperties{}, obj)

				cpuProps := obj.(qemu.CPUProperties)
				fmt.Printf("%+v\n", cpuProps)

				assert.Equal(
					t,
					qemu.DefaultCPUPropertyKind,
					cpuProps.Kind,
				)
				assert.Equal(
					t,
					qemu.DefaultCPUPropertyArchitecture,
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
					qemu.DefaultCPUPropertyLimit,
					cpuProps.Limit,
				)
				assert.Equal(
					t,
					qemu.DefaultCPUPropertyUnits,
					cpuProps.Units,
				)
				assert.Equal(
					t,
					qemu.DefaultCPUPropertyNUMA,
					cpuProps.NUMA,
				)
				assert.Equal(
					t,
					qemu.DefaultCPUPropertyFreezeAtStartup,
					cpuProps.FreezeAtStartup,
				)
			},
		),
	)
}
