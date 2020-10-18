package lxc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/lxc"
	"github.com/xabinapal/gopve/test"
)

func TestCPUProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"arch":     "arm64",
		"cores":    16,
		"cpulimit": 2,
		"cpuunits": 2048,
	})

	requiredProps := []string{"arch", "cores"}

	defaultProps := []string{"cpulimit", "cpuunits"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := lxc.NewCPUProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			cpuProps, err := lxc.NewCPUProperties(props)
			require.NoError(t, err)

			assert.Equal(t, lxc.CPUArchitectureARM64, cpuProps.Architecture)
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
				require.IsType(t, (*lxc.CPUProperties)(nil), obj)

				cpuProps := obj.(*lxc.CPUProperties)

				assert.Equal(
					t,
					lxc.DefaultCPUPropertyLimit,
					cpuProps.Limit,
				)
				assert.Equal(
					t,
					lxc.DefaultCPUPropertyUnits,
					cpuProps.Units,
				)
			},
		),
	)
}
