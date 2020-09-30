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

	t.Run(
		"Unmarshal", func(t *testing.T) {
			var storageProps storage.StorageDirProperties

			err := (&storageProps).Unmarshal(props)
			require.NoError(t, err)

			assert.Equal(t, "test_path", storageProps.LocalPath)
			assert.Equal(t, true, storageProps.LocalPathCreate)
			assert.Equal(t, true, storageProps.LocalPathIsManaged)
		})

	t.Run(
		"RequiredProperties", func(t *testing.T) {
			var storageProps storage.StorageDirProperties

			helperTestRequiredProperties(t, props, requiredProps, (&storageProps).Unmarshal)
		})

	t.Run("DefaultProperties", func(t *testing.T) {
		var storageProps storage.StorageDirProperties

		helperTestOptionalProperties(t, props, defaultProps, (&storageProps).Unmarshal)

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
	})
}
