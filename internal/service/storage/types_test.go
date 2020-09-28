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
	allProps := map[string](interface{}){
		"path":          "test_path",
		"mkdir":         1,
		"is_mountpoint": 1,
	}

	requiredProps := []string{"path"}

	optionalProps := []string{"mkdir", "is_mountpoint"}

	newStorage := func(props map[string](interface{})) (types.Storage, error) {
		return storage.NewDynamicStorage(
			nil,
			"test_storage",
			types.KindDir,
			types.Properties{
				ExtraProperties: props,
			},
		)
	}

	t.Run("Create", func(t *testing.T) {
		obj, err := newStorage(allProps)

		require.NoError(t, err)
		require.Implements(t, (*types.StorageDir)(nil), obj)

		concreteStorage := obj.(types.StorageDir)

		assert.Equal(t, "test_path", concreteStorage.LocalPath())
		assert.Equal(t, true, concreteStorage.LocalPathCreate())
		assert.Equal(t, true, concreteStorage.LocalPathIsManaged())
	})

	t.Run("RequiredProperties", func(t *testing.T) {
		for _, prop := range requiredProps {
			finalProps := make(map[string]interface{}, len(allProps))
			for k, v := range allProps {
				finalProps[k] = v
			}
			delete(finalProps, prop)

			_, err := newStorage(finalProps)

			expectedError := types.ErrMissingProperty
			expectedError.AddKey("name", prop)

			assert.EqualError(t, err, expectedError.Error())
		}
	})

	t.Run("DefaultProperties", func(t *testing.T) {
		finalProps := make(map[string]interface{}, len(allProps))
		for k, v := range allProps {
			finalProps[k] = v
		}
		for _, prop := range optionalProps {
			delete(finalProps, prop)
		}

		obj, err := newStorage(finalProps)

		require.NoError(t, err)

		concreteStorage := obj.(types.StorageDir)

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
	allProps := map[string](interface{}){
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

	newStorage := func(props map[string](interface{})) (types.Storage, error) {
		return storage.NewDynamicStorage(
			nil,
			"test_storage",
			types.KindLVM,
			types.Properties{
				ExtraProperties: props,
			},
		)
	}

	t.Run("Create", func(t *testing.T) {
		obj, err := newStorage(allProps)

		require.NoError(t, err)
		require.Implements(t, (*types.StorageLVM)(nil), obj)

		concreteStorage := obj.(types.StorageLVM)

		assert.Equal(t, "test_base", concreteStorage.BaseStorage())
		assert.Equal(t, "test_vg", concreteStorage.VolumeGroup())
		assert.Equal(t, true, concreteStorage.SafeRemove())
		assert.Equal(t, 1024, concreteStorage.SafeRemoveThroughput())
		assert.Equal(t, true, concreteStorage.TaggedOnly())
	})

	t.Run("RequiredProperties", func(t *testing.T) {
		for _, prop := range requiredProps {
			finalProps := make(map[string]interface{}, len(allProps))
			for k, v := range allProps {
				finalProps[k] = v
			}
			delete(finalProps, prop)

			_, err := newStorage(finalProps)

			expectedError := types.ErrMissingProperty
			expectedError.AddKey("name", prop)

			assert.EqualError(t, err, expectedError.Error())
		}
	})

	t.Run("DefaultProperties", func(t *testing.T) {
		finalProps := make(map[string]interface{}, len(allProps))
		for k, v := range allProps {
			finalProps[k] = v
		}
		for _, prop := range optionalProps {
			delete(finalProps, prop)
		}

		obj, err := newStorage(finalProps)

		require.NoError(t, err)

		concreteStorage := obj.(types.StorageLVM)

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
	allProps := map[string](interface{}){
		"thinpool": "test_pool",
		"vgname":   "test_vg",
	}

	requiredProps := []string{"thinpool", "vgname"}

	newStorage := func(props map[string](interface{})) (types.Storage, error) {
		return storage.NewDynamicStorage(
			nil,
			"test_storage",
			types.KindLVMThin,
			types.Properties{
				ExtraProperties: props,
			},
		)
	}

	t.Run("Create", func(t *testing.T) {
		obj, err := newStorage(allProps)

		require.NoError(t, err)
		require.Implements(t, (*types.StorageLVMThin)(nil), obj)

		concreteStorage := obj.(types.StorageLVMThin)

		assert.Equal(t, "test_pool", concreteStorage.ThinPool())
		assert.Equal(t, "test_vg", concreteStorage.VolumeGroup())
	})

	t.Run("RequiredProperties", func(t *testing.T) {
		for _, prop := range requiredProps {
			finalProps := make(map[string]interface{}, len(allProps))
			for k, v := range allProps {
				finalProps[k] = v
			}
			delete(finalProps, prop)

			_, err := newStorage(finalProps)

			expectedError := types.ErrMissingProperty
			expectedError.AddKey("name", prop)

			assert.EqualError(t, err, expectedError.Error())
		}
	})
}

func TestStorageNewZFS(t *testing.T) {
	allProps := map[string](interface{}){
		"pool":       "test_pool",
		"blocksize":  1024,
		"sparse":     1,
		"mountpoint": "test_mountpoint",
	}

	requiredProps := []string{"pool"}

	optionalProps := []string{"blocksize", "sparse", "mountpoint"}

	newStorage := func(props map[string](interface{})) (types.Storage, error) {
		return storage.NewDynamicStorage(
			nil,
			"test_storage",
			types.KindZFS,
			types.Properties{
				ExtraProperties: props,
			},
		)
	}

	t.Run("Create", func(t *testing.T) {
		obj, err := newStorage(allProps)

		require.NoError(t, err)
		require.Implements(t, (*types.StorageZFS)(nil), obj)

		concreteStorage := obj.(types.StorageZFS)

		assert.Equal(t, "test_pool", concreteStorage.PoolName())
		assert.Equal(t, uint(1024), concreteStorage.BlockSize())
		assert.Equal(t, true, concreteStorage.UseSparse())
		assert.Equal(t, "test_mountpoint", concreteStorage.LocalPath())
	})

	t.Run("RequiredProperties", func(t *testing.T) {
		for _, prop := range requiredProps {
			finalProps := make(map[string]interface{}, len(allProps))
			for k, v := range allProps {
				finalProps[k] = v
			}
			delete(finalProps, prop)

			_, err := newStorage(finalProps)

			expectedError := types.ErrMissingProperty
			expectedError.AddKey("name", prop)

			assert.EqualError(t, err, expectedError.Error())
		}
	})

	t.Run("DefaultProperties", func(t *testing.T) {
		finalProps := make(map[string]interface{}, len(allProps))
		for k, v := range allProps {
			finalProps[k] = v
		}
		for _, prop := range optionalProps {
			delete(finalProps, prop)
		}

		obj, err := newStorage(finalProps)

		require.NoError(t, err)

		concreteStorage := obj.(types.StorageZFS)

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
