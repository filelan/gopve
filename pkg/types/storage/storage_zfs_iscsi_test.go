package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/test"
)

func TestStorageZFSOverISCSIProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"portal":        "test_portal",
		"target":        "test_target",
		"pool":          "test_pool",
		"blocksize":     "1024",
		"sparse":        1,
		"nowritecache":  1,
		"iscsiprovider": "iet",
		"comstar_hg":    "test_comstar_hg",
		"comstar_tg":    "test_comstar_tg",
		"lio_tpg":       "test_lio_tpg",
	})

	requiredProps := []string{
		"portal",
		"target",
		"pool",
		"blocksize",
		"iscsiprovider",
	}

	defaultProps := []string{
		"sparse",
		"nowritecache",
		"comstar_hg",
		"comstar_tg",
		"lio_tpg",
	}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := storage.NewStorageZFSOverISCSIProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			storageProps, err := storage.NewStorageZFSOverISCSIProperties(props)
			require.NoError(t, err)

			assert.Equal(t, "test_portal", storageProps.Portal)
			assert.Equal(t, "test_target", storageProps.Target)
			assert.Equal(t, "test_pool", storageProps.PoolName)
			assert.Equal(t, "1024", storageProps.BlockSize)
			assert.Equal(t, true, storageProps.UseSparse)
			assert.Equal(t, false, storageProps.WriteCache)
			assert.Equal(
				t,
				storage.ISCSIProviderIET,
				storageProps.ISCSIProvider,
			)
			assert.Equal(t, "test_comstar_hg", storageProps.ComstarHostGroup)
			assert.Equal(t, "test_comstar_tg", storageProps.ComstarTargetGroup)
			assert.Equal(t, "test_lio_tpg", storageProps.LIOTargetPortalGroup)
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
				require.IsType(
					t,
					(*storage.StorageZFSOverISCSIProperties)(nil),
					obj,
				)

				storageProps := obj.(*storage.StorageZFSOverISCSIProperties)

				assert.ElementsMatch(
					t,
					storage.DefaultStorageZFSOverISCSIUseSparse,
					storageProps.UseSparse,
				)

				assert.Equal(
					t,
					storage.DefaultStorageZFSOverISCSIWriteCache,
					storageProps.WriteCache,
				)

				assert.Equal(
					t,
					storage.DefaultStorageZFSOverISCSIComstarHostGroup,
					storageProps.ComstarHostGroup,
				)

				assert.Equal(
					t,
					storage.DefaultStorageZFSOverISCSIComstarTargetGroup,
					storageProps.ComstarTargetGroup,
				)

				assert.Equal(
					t,
					storage.DefaultStorageZFSOverISCSILIOTargetPortalGroup,
					storageProps.LIOTargetPortalGroup,
				)
			},
		),
	)
}
