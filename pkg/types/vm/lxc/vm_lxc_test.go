package lxc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/lxc"
	"github.com/xabinapal/gopve/test"
)

func TestGlobalProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"ostype":     "archlinux",
		"protection": 1,
		"onboot":     1,
	})

	requiredProps := []string{"ostype"}

	defaultProps := []string{"protection", "onboot"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := lxc.NewGlobalProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			globalProps, err := lxc.NewGlobalProperties(props)
			require.NoError(t, err)

			assert.Equal(t, lxc.OSTypeArchLinux, globalProps.OSType)
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
				require.IsType(t, (*lxc.GlobalProperties)(nil), obj)

				globalProps := obj.(*lxc.GlobalProperties)

				assert.Equal(
					t,
					lxc.DefaultGlobalPropertyProtected,
					globalProps.Protected,
				)
				assert.Equal(
					t,
					lxc.DefaultGlobalPropertyStartAtBoot,
					globalProps.StartAtBoot,
				)
			},
		),
	)
}
