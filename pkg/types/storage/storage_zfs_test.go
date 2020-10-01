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

	factoryFunc := func(props storage.ExtraProperties) (interface{}, error) {
		obj, err := storage.NewStorageZFSProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			storageProps, err := storage.NewStorageZFSProperties(props)
			require.NoError(t, err)

			assert.Equal(t, "test_pool", storageProps.PoolName)
			assert.Equal(t, "1024", storageProps.BlockSize)
			assert.Equal(t, true, storageProps.UseSparse)
			assert.Equal(t, "test_mountpoint", storageProps.LocalPath)
		})

	t.Run(
		"RequiredProperties", helperTestRequiredProperties(t, props, requiredProps, factoryFunc))

	t.Run("DefaultProperties", helperTestOptionalProperties(t, props, defaultProps, factoryFunc, func(obj interface{}) {
		require.IsType(t, (*storage.StorageZFSProperties)(nil), obj)

		storageProps := obj.(*storage.StorageZFSProperties)

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
	},
	))
}
