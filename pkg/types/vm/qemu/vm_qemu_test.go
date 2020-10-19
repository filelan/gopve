package qemu_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
	"github.com/xabinapal/gopve/test"
)

func TestGlobalProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"ostype": "l26",
		"acpi":   0,
		"kvm":    0,
		"tablet": 0,
	})

	requiredProps := []string{"ostype"}

	defaultProps := []string{"acpi", "kvm", "tablet"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := qemu.NewGlobalProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			globalProps, err := qemu.NewGlobalProperties(props)
			require.NoError(t, err)

			assert.Equal(t, qemu.OSTypeLinux26, globalProps.OSType)
			assert.Equal(t, false, globalProps.ACPI)
			assert.Equal(t, false, globalProps.KVMVirtualization)
			assert.Equal(t, false, globalProps.USBTabletDevice)
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
				require.IsType(t, qemu.GlobalProperties{}, obj)

				globalProps := obj.(qemu.GlobalProperties)

				assert.Equal(
					t,
					qemu.DefaultGlobalPropertiesACPI,
					globalProps.ACPI,
				)
				assert.Equal(
					t,
					qemu.DefaultGlobalPropertiesKVMVirtualization,
					globalProps.KVMVirtualization,
				)
				assert.Equal(
					t,
					qemu.DefaultGlobalPropertiesUSBTabletDevice,
					globalProps.USBTabletDevice,
				)
			},
		),
	)
}
