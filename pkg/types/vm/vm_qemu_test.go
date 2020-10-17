package vm_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
	"github.com/xabinapal/gopve/test"
)

func TestQEMUGlobalProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"ostype":     "l26",
		"protection": 1,
		"onboot":     1,
	})

	requiredProps := []string{"ostype"}

	defaultProps := []string{"protection", "onboot"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := vm.NewQEMUGlobalProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			globalProps, err := vm.NewQEMUGlobalProperties(props)
			require.NoError(t, err)

			assert.Equal(t, qemu.OSTypeLinux26, globalProps.OSType)
			assert.Equal(t, true, globalProps.Protected)
			assert.Equal(t, true, globalProps.StartOnBoot)
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
				require.IsType(t, vm.QEMUGlobalProperties{}, obj)

				globalProps := obj.(vm.QEMUGlobalProperties)

				assert.Equal(
					t,
					vm.DefaultQEMUGlobalPropertyProtected,
					globalProps.Protected,
				)
				assert.Equal(
					t,
					vm.DefaultQEMUGlobalPropertyProtected,
					globalProps.Protected,
				)
				assert.Equal(
					t,
					vm.DefaultQEMUGlobalPropertyStartOnBoot,
					globalProps.StartOnBoot,
				)
			},
		),
	)
}
