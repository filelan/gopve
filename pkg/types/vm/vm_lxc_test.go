package vm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
	"github.com/xabinapal/gopve/pkg/types/vm/lxc"
	"github.com/xabinapal/gopve/test"
)

func TestLXCGlobalProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"ostype":     "archlinux",
		"protection": 1,
		"onboot":     1,
	})

	requiredProps := []string{"ostype"}

	defaultProps := []string{"protection", "onboot"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := vm.NewLXCGlobalProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			globalProps, err := vm.NewLXCGlobalProperties(props)
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
				require.IsType(t, (*vm.LXCGlobalProperties)(nil), obj)

				globalProps := obj.(*vm.LXCGlobalProperties)

				assert.Equal(
					t,
					vm.DefaultLXCGlobalPropertyProtected,
					globalProps.Protected,
				)
				assert.Equal(
					t,
					vm.DefaultLXCGlobalPropertyStartAtBoot,
					globalProps.StartAtBoot,
				)
			},
		),
	)
}
