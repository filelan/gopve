package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/test"
)

func TestStorageLVMThinProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"vgname":   "test_vg",
		"thinpool": "test_pool",
	})

	requiredProps := []string{"vgname", "thinpool"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := storage.NewStorageLVMThinProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			storageProps, err := storage.NewStorageLVMThinProperties(props)
			require.NoError(t, err)

			assert.Equal(t, "test_vg", storageProps.VolumeGroup)
			assert.Equal(t, "test_pool", storageProps.ThinPool)
		})

	t.Run(
		"RequiredProperties",
		test.HelperTestRequiredProperties(t, props, requiredProps, factoryFunc),
	)
}
