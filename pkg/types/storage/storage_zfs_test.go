package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/test"
)

func TestStorageZFSProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"pool":       "test_pool",
		"blocksize":  "1024",
		"sparse":     1,
		"mountpoint": "test_mountpoint",
	})

	requiredProps := []string{"pool"}

	defaultProps := []string{"blocksize", "sparse", "mountpoint"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
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
		),
	)
}
