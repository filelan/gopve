package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

func TestStorageDir(t *testing.T) {
	props := map[string]interface{}{
		"path":          "test_path",
		"mkdir":         1,
		"is_mountpoint": 1,
	}

	requiredProps := []string{"path"}

	defaultProps := []string{"mkdir", "is_mountpoint"}

	factoryFunc := func(props storage.ExtraProperties) (interface{}, error) {
		obj, err := storage.NewStorageDirProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			storageProps, err := storage.NewStorageDirProperties(props)
			require.NoError(t, err)

			assert.Equal(t, "test_path", storageProps.LocalPath)
			assert.Equal(t, true, storageProps.LocalPathCreate)
			assert.Equal(t, true, storageProps.LocalPathIsManaged)
		})

	t.Run(
		"RequiredProperties", helperTestRequiredProperties(t, props, requiredProps, factoryFunc))

	t.Run("DefaultProperties", helperTestOptionalProperties(t, props, defaultProps, factoryFunc, func(obj interface{}) {
		require.IsType(t, (*storage.StorageDirProperties)(nil), obj)

		storageProps := obj.(*storage.StorageDirProperties)

		assert.Equal(
			t,
			storage.DefaultStorageDirLocalPathCreate,
			storageProps.LocalPathCreate,
		)

		assert.Equal(
			t,
			storage.DefaultStorageDirLocalIsManaged,
			storageProps.LocalPathIsManaged,
		)
	},
	))
}
