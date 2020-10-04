package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/test"
)

func TestStorageCephFSProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"monhost":  "test_host_1 test_host_2 test_host_3",
		"username": "test_username",
		"fuse":     1,
		"subdir":   "/test_subdir",
		"path":     "/test_path",
	})

	requiredProps := []string{"path"}

	defaultProps := []string{"monhost", "username", "fuse", "subdir"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := storage.NewStorageCephFSProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			storageProps, err := storage.NewStorageCephFSProperties(props)
			require.NoError(t, err)

			assert.ElementsMatch(
				t,
				[]string{"test_host_1", "test_host_2", "test_host_3"},
				storageProps.MonitorHosts,
			)
			assert.Equal(t, "test_username", storageProps.Username)
			assert.Equal(t, true, storageProps.UseFUSE)
			assert.Equal(t, "/test_subdir", storageProps.ServerPath)
			assert.Equal(t, "/test_path", storageProps.LocalPath)
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
				require.IsType(t, (*storage.StorageCephFSProperties)(nil), obj)

				storageProps := obj.(*storage.StorageCephFSProperties)

				assert.ElementsMatch(
					t,
					storage.DefaultStorageCephFSMonitorHosts,
					storageProps.MonitorHosts,
				)

				assert.Equal(
					t,
					storage.DefaultStorageCephFSUsername,
					storageProps.Username,
				)

				assert.Equal(
					t,
					storage.DefaultStorageCephFSUseFUSE,
					storageProps.UseFUSE,
				)

				assert.Equal(
					t,
					storage.DefaultStorageCephFSServerPath,
					storageProps.ServerPath,
				)
			},
		),
	)
}
