package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/storage"
	types "github.com/xabinapal/gopve/pkg/types/storage"
)

func TestStorageNew(t *testing.T) {
}

func TestStorageNewDir(t *testing.T) {
	t.Run("Create", func(t *testing.T) {
		obj, err := storage.NewDynamicStorage(nil, "test_storage", types.KindDir, types.Properties{
			ExtraProperties: map[string](interface{}){
				"path":          "test_path",
				"mkdir":         1,
				"is_mountpoint": 1,
			},
		})

		require.NoError(t, err)
		require.Implements(t, (*types.StorageDir)(nil), obj)
		assert.Equal(t, "test_path", obj.(types.StorageDir).LocalPath())
		assert.Equal(t, true, obj.(types.StorageDir).LocalPathCreate())
		assert.Equal(t, true, obj.(types.StorageDir).LocalPathIsManaged())
	})

	t.Run("RequirePathProperty", func(t *testing.T) {
		_, err := storage.NewDynamicStorage(nil, "test_storage", types.KindDir, types.Properties{
			ExtraProperties: map[string](interface{}){
				"mkdir":         "1",
				"is_mountpoint": "1",
			},
		})

		expectedError := types.ErrMissingProperty
		expectedError.AddKey("name", "path")

		assert.EqualError(t, err, expectedError.Error())
	})

	t.Run("DefaultProperties", func(t *testing.T) {
		obj, err := storage.NewDynamicStorage(nil, "test_storage", types.KindDir, types.Properties{
			ExtraProperties: map[string](interface{}){
				"path":          "test_path",
			},
		})

		require.NoError(t, err)
		assert.Equal(t, types.DefaultStorageLocalPathCreate, obj.(types.StorageDir).LocalPathCreate())
		assert.Equal(t, types.DefaultStorageLocalIsManaged, obj.(types.StorageDir).LocalPathIsManaged())
	})
}
