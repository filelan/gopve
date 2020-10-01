package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

func TestStorageRBD(t *testing.T) {
	props := map[string]interface{}{
		"monhost":  "test_host_1 test_host_2 test_host_3",
		"username": "test_username",
		"krbd":     1,
		"pool":     "test_pool",
	}

	defaultProps := []string{"monhost", "username", "krbd", "pool"}

	factoryFunc := func(props storage.ExtraProperties) (interface{}, error) {
		obj, err := storage.NewStorageRBDProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			storageProps, err := storage.NewStorageRBDProperties(props)
			require.NoError(t, err)

			assert.ElementsMatch(
				t,
				[]string{"test_host_1", "test_host_2", "test_host_3"},
				storageProps.MonitorHosts,
			)
			assert.Equal(t, "test_username", storageProps.Username)
			assert.Equal(t, true, storageProps.UseKRBD)
			assert.Equal(t, "test_pool", storageProps.PoolName)
		})

	t.Run("DefaultProperties", helperTestOptionalProperties(t, props, defaultProps, factoryFunc, func(obj interface{}) {
		require.IsType(t, (*storage.StorageRBDProperties)(nil), obj)

		storageProps := obj.(*storage.StorageRBDProperties)

		assert.ElementsMatch(
			t,
			storage.DefaultStorageRBDMonitorHosts,
			storageProps.MonitorHosts,
		)

		assert.Equal(
			t,
			storage.DefaultStorageRBDUsername,
			storageProps.Username,
		)

		assert.Equal(
			t,
			storage.DefaultStorageRBDUseKRBD,
			storageProps.UseKRBD,
		)

		assert.Equal(
			t,
			storage.DefaultStorageRBDPoolName,
			storageProps.PoolName,
		)
	},
	))
}
