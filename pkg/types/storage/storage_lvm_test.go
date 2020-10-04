package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/test"
)

func TestStorageLVMProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"base":                  "test_base",
		"vgname":                "test_vg",
		"saferemove":            1,
		"saferemove_throughput": 1024,
		"tagged_only":           1,
	})

	requiredProps := []string{"vgname"}

	defaultProps := []string{
		"base",
		"saferemove",
		"saferemove_throughput",
		"tagged_only",
	}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := storage.NewStorageLVMProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			storageProps, err := storage.NewStorageLVMProperties(props)
			require.NoError(t, err)

			assert.Equal(t, "test_base", storageProps.BaseStorage)
			assert.Equal(t, "test_vg", storageProps.VolumeGroup)
			assert.Equal(t, true, storageProps.SafeRemove)
			assert.Equal(t, 1024, storageProps.SafeRemoveThroughput)
			assert.Equal(t, true, storageProps.TaggedOnly)
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
				require.IsType(t, (*storage.StorageLVMProperties)(nil), obj)

				storageProps := obj.(*storage.StorageLVMProperties)

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
			},
		),
	)
}
