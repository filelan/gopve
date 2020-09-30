package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

func TestStorageLVM(t *testing.T) {
	props := map[string]interface{}{
		"base":                  "test_base",
		"vgname":                "test_vg",
		"saferemove":            1,
		"saferemove_throughput": 1024,
		"tagged_only":           1,
	}

	requiredProps := []string{"vgname"}

	defaultProps := []string{"base", "saferemove", "saferemove_throughput", "tagged_only"}

	t.Run(
		"Unmarshal", func(t *testing.T) {
			var storageProps storage.StorageLVMProperties

			err := (&storageProps).Unmarshal(props)
			require.NoError(t, err)

			assert.Equal(t, "test_base", storageProps.BaseStorage)
			assert.Equal(t, "test_vg", storageProps.VolumeGroup)
			assert.Equal(t, true, storageProps.SafeRemove)
			assert.Equal(t, 1024, storageProps.SafeRemoveThroughput)
			assert.Equal(t, true, storageProps.TaggedOnly)
		})

	t.Run(
		"RequiredProperties", func(t *testing.T) {
			var storageProps storage.StorageLVMProperties

			helperTestRequiredProperties(t, props, requiredProps, (&storageProps).Unmarshal)
		})

	t.Run("DefaultProperties", func(t *testing.T) {
		var storageProps storage.StorageLVMProperties

		helperTestOptionalProperties(t, props, defaultProps, (&storageProps).Unmarshal)

		assert.Equal(
			t,
			storage.DefaultStorageLVMBaseStorage,
			storageProps.BaseStorage,
		)
		assert.Equal(
			t,
			storage.DefaultStorageLVMSafeRemove,
			storageProps.SafeRemove,
		)
		assert.Equal(
			t,
			storage.DefaultStorageLVMSafeRemoveThroughput,
			storageProps.SafeRemoveThroughput,
		)
		assert.Equal(
			t,
			storage.DefaultStorageLVMTaggedOnly,
			storageProps.TaggedOnly,
		)
	})
}
