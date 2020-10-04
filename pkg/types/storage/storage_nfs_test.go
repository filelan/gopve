package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/test"
)

func TestStorageNFSProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"server":  "test_server",
		"options": "vers=4.2",
		"export":  "/test_export",
		"path":    "/test_path",
		"mkdir":   1,
	})

	requiredProps := []string{"server", "export", "path"}

	defaultProps := []string{"options", "mkdir"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := storage.NewStorageNFSProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			storageProps, err := storage.NewStorageNFSProperties(props)
			require.NoError(t, err)

			assert.Equal(t, "test_server", storageProps.Server)
			assert.Equal(t, storage.NFSVersion42, storageProps.NFSVersion)
			assert.Equal(t, "/test_export", storageProps.ServerPath)
			assert.Equal(t, "/test_path", storageProps.LocalPath)
			assert.Equal(t, true, storageProps.LocalPathCreate)
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
				require.IsType(t, (*storage.StorageNFSProperties)(nil), obj)

				storageProps := obj.(*storage.StorageNFSProperties)

				assert.Equal(
					t,
					storage.DefaultStorageNFSVersion,
					storageProps.NFSVersion,
				)

				assert.Equal(
					t,
					storage.DefaultStorageNFSLocalPathCreate,
					storageProps.LocalPathCreate,
				)
			},
		),
	)
}

func TestNFSVersion(t *testing.T) {
	NFSVersionCases := map[string](struct {
		Object storage.NFSVersion
		Value  string
	}){
		"3": {
			Object: storage.NFSVersion30,
			Value:  "3",
		},
		"4": {
			Object: storage.NFSVersion40,
			Value:  "4",
		},
		"4.1": {
			Object: storage.NFSVersion41,
			Value:  "4.1",
		},
		"4.2": {
			Object: storage.NFSVersion42,
			Value:  "4.2",
		},
	}

	t.Run("Marshal", func(t *testing.T) {
		for n, tt := range NFSVersionCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				var receivedTransport storage.NFSVersion
				err := (&receivedTransport).Unmarshal(tt.Value)
				require.NoError(t, err)
				assert.Equal(t, tt.Object, receivedTransport)
			})
		}
	})

	t.Run("Unmarshal", func(t *testing.T) {
		for n, tt := range NFSVersionCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				receivedValue, err := tt.Object.Marshal()
				require.NoError(t, err)
				assert.Equal(t, tt.Value, receivedValue)
			})
		}
	})
}
