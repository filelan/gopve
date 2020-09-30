package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

func TestStorageLVMThin(t *testing.T) {

	props := map[string]interface{}{
		"vgname":   "test_vg",
		"thinpool": "test_pool",
	}

	requiredProps := []string{"vgname", "thinpool"}

	t.Run(
		"Unmarshal", func(t *testing.T) {
			var storageProps storage.StorageLVMThinProperties

			err := (&storageProps).Unmarshal(props)
			require.NoError(t, err)

			assert.Equal(t, "test_vg", storageProps.VolumeGroup)
			assert.Equal(t, "test_pool", storageProps.ThinPool)
		})

	t.Run(
		"RequiredProperties", func(t *testing.T) {
			var storageProps storage.StorageLVMThinProperties

			helperTestRequiredProperties(t, props, requiredProps, (&storageProps).Unmarshal)
		})
}
