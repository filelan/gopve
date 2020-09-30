package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

func TestStorageZFS(t *testing.T) {
	props := map[string]interface{}{
		"pool":       "test_pool",
		"blocksize":  "1024",
		"sparse":     1,
		"mountpoint": "test_mountpoint",
	}

	requiredProps := []string{"pool"}

	defaultProps := []string{"blocksize", "sparse", "mountpoint"}

	t.Run(
		"Unmarshal", func(t *testing.T) {
			var storageProps storage.StorageZFSProperties

			err := (&storageProps).Unmarshal(props)
			require.NoError(t, err)

			assert.Equal(t, "test_pool", storageProps.PoolName)
			assert.Equal(t, "1024", storageProps.BlockSize)
			assert.Equal(t, true, storageProps.UseSparse)
			assert.Equal(t, "test_mountpoint", storageProps.LocalPath)
		})

	t.Run(
		"RequiredProperties", func(t *testing.T) {
			var storageProps storage.StorageZFSProperties

			helperTestRequiredProperties(t, props, requiredProps, (&storageProps).Unmarshal)
		})

	t.Run("DefaultProperties", func(t *testing.T) {
		var storageProps storage.StorageZFSProperties

		helperTestOptionalProperties(t, props, defaultProps, (&storageProps).Unmarshal)

		assert.Equal(
			t,
			storage.DefaultStorageZFSBlockSize,
			storageProps.BlockSize,
		)
		assert.Equal(
			t,
			storage.DefaultStorageZFSUseSparse,
			storageProps.UseSparse,
		)
		assert.Equal(
			t,
			storage.DefaultStorageZFSMountPoint,
			storageProps.LocalPath,
		)
	})
}
