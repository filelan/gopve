package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/internal/service/storage"
	types "github.com/xabinapal/gopve/pkg/types/storage"
)

func helperCreateStorage(kind types.Kind, props storage.ExtraProperties) (types.Storage, error) {
	return storage.NewDynamicStorage(
		nil,
		"test_storage",
		kind,
		types.Properties{
			ExtraProperties: props,
		},
	)
}

func helperFilterOptionalProperties(props storage.ExtraProperties, optionalProps []string) storage.ExtraProperties {
	finalProps := make(storage.ExtraProperties, len(props))
	for k, v := range props {
		finalProps[k] = v
	}

	for _, prop := range optionalProps {
		delete(finalProps, prop)
	}

	return finalProps
}

func helperTestRequiredProperties(t *testing.T, kind types.Kind, props storage.ExtraProperties, requiredProps []string) func(t *testing.T) {
	t.Helper()

	return func(t *testing.T) {
		for _, prop := range requiredProps {
			finalProps := make(storage.ExtraProperties, len(props))
			for k, v := range props {
				finalProps[k] = v
			}

			delete(finalProps, prop)

			_, err := helperCreateStorage(kind, finalProps)

			expectedError := types.ErrMissingProperty
			expectedError.AddKey("name", prop)

			assert.EqualError(t, err, expectedError.Error())
		}
	}
}

func TestStorageNew(t *testing.T) {
}

func TestStorageNewDir(t *testing.T) {
	kind := types.KindDir

	props := storage.ExtraProperties{
		"path":          "test_path",
		"mkdir":         1,
		"is_mountpoint": 1,
	}

	requiredProps := []string{"path"}

	optionalProps := []string{"mkdir", "is_mountpoint"}

	t.Run("Create", func(t *testing.T) {
		obj, err := helperCreateStorage(kind, props)
		require.NoError(t, err)

		require.Implements(t, (*types.StorageDir)(nil), obj)

		concreteStorage, ok := obj.(types.StorageDir)
		require.Equal(t, true, ok)

		assert.Equal(t, "test_path", concreteStorage.LocalPath())
		assert.Equal(t, true, concreteStorage.LocalPathCreate())
		assert.Equal(t, true, concreteStorage.LocalPathIsManaged())
	})

	t.Run("RequiredProperties", helperTestRequiredProperties(t, kind, props, requiredProps))

	t.Run("DefaultProperties", func(t *testing.T) {
		finalProps := helperFilterOptionalProperties(props, optionalProps)

		obj, err := helperCreateStorage(kind, finalProps)
		require.NoError(t, err)

		concreteStorage, ok := obj.(types.StorageDir)
		require.Equal(t, true, ok)

		assert.Equal(
			t,
			types.DefaultStorageDirLocalPathCreate,
			concreteStorage.LocalPathCreate(),
		)
		assert.Equal(
			t,
			types.DefaultStorageDirLocalIsManaged,
			concreteStorage.LocalPathIsManaged(),
		)
	})
}

func TestStorageNewLVM(t *testing.T) {
	kind := types.KindLVM

	props := storage.ExtraProperties{
		"base":                  "test_base",
		"vgname":                "test_vg",
		"saferemove":            1,
		"saferemove_throughput": 1024,
		"tagged_only":           1,
	}

	requiredProps := []string{"vgname"}

	optionalProps := []string{
		"base",
		"saferemove",
		"saferemove_throughput",
		"tagged_only",
	}

	t.Run("Create", func(t *testing.T) {
		obj, err := helperCreateStorage(kind, props)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageLVM)(nil), obj)

		concreteStorage, ok := obj.(types.StorageLVM)
		require.Equal(t, true, ok)

		assert.Equal(t, "test_base", concreteStorage.BaseStorage())
		assert.Equal(t, "test_vg", concreteStorage.VolumeGroup())
		assert.Equal(t, true, concreteStorage.SafeRemove())
		assert.Equal(t, 1024, concreteStorage.SafeRemoveThroughput())
		assert.Equal(t, true, concreteStorage.TaggedOnly())
	})

	t.Run("RequiredProperties", helperTestRequiredProperties(t, kind, props, requiredProps))

	t.Run("DefaultProperties", func(t *testing.T) {
		finalProps := helperFilterOptionalProperties(props, optionalProps)

		obj, err := helperCreateStorage(kind, finalProps)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageLVM)(nil), obj)

		concreteStorage, ok := obj.(types.StorageLVM)
		require.Equal(t, true, ok)

		assert.Equal(
			t,
			types.DefaultStorageLVMBaseStorage,
			concreteStorage.BaseStorage(),
		)
		assert.Equal(
			t,
			types.DefaultStorageLVMSafeRemove,
			concreteStorage.SafeRemove(),
		)
		assert.Equal(
			t,
			types.DefaultStorageLVMSafeRemoveThroughput,
			concreteStorage.SafeRemoveThroughput(),
		)
		assert.Equal(
			t,
			types.DefaultStorageLVMTaggedOnly,
			concreteStorage.TaggedOnly(),
		)
	})
}

func TestStorageNewLVMThin(t *testing.T) {
	kind := types.KindLVMThin

	props := storage.ExtraProperties{
		"thinpool": "test_pool",
		"vgname":   "test_vg",
	}

	requiredProps := []string{"thinpool", "vgname"}

	t.Run("Create", func(t *testing.T) {
		obj, err := helperCreateStorage(kind, props)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageLVMThin)(nil), obj)

		concreteStorage, ok := obj.(types.StorageLVMThin)
		require.Equal(t, true, ok)

		assert.Equal(t, "test_pool", concreteStorage.ThinPool())
		assert.Equal(t, "test_vg", concreteStorage.VolumeGroup())
	})

	t.Run("RequiredProperties", helperTestRequiredProperties(t, kind, props, requiredProps))
}

func TestStorageNewZFS(t *testing.T) {
	kind := types.KindZFS

	props := storage.ExtraProperties{
		"pool":       "test_pool",
		"blocksize":  1024,
		"sparse":     1,
		"mountpoint": "test_mountpoint",
	}

	requiredProps := []string{"pool"}

	optionalProps := []string{"blocksize", "sparse", "mountpoint"}

	t.Run("Create", func(t *testing.T) {
		obj, err := helperCreateStorage(kind, props)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageZFS)(nil), obj)

		concreteStorage, ok := obj.(types.StorageZFS)
		require.Equal(t, true, ok)

		assert.Equal(t, "test_pool", concreteStorage.PoolName())
		assert.Equal(t, uint(1024), concreteStorage.BlockSize())
		assert.Equal(t, true, concreteStorage.UseSparse())
		assert.Equal(t, "test_mountpoint", concreteStorage.LocalPath())
	})

	t.Run("RequiredProperties", helperTestRequiredProperties(t, kind, props, requiredProps))

	t.Run("DefaultProperties", func(t *testing.T) {
		finalProps := helperFilterOptionalProperties(props, optionalProps)

		obj, err := helperCreateStorage(kind, finalProps)
		require.NoError(t, err)
		require.Implements(t, (*types.StorageZFS)(nil), obj)

		concreteStorage, ok := obj.(types.StorageZFS)
		require.Equal(t, true, ok)

		assert.Equal(
			t,
			types.DefaultStorageZFSBlockSize,
			concreteStorage.BlockSize(),
		)
		assert.Equal(
			t,
			types.DefaultStorageZFSUseSparse,
			concreteStorage.UseSparse(),
		)
		assert.Equal(
			t,
			types.DefaultStorageZFSMountPoint,
			concreteStorage.LocalPath(),
		)
	})
}
