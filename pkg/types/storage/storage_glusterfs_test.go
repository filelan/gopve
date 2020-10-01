package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

func TestStorageGlusterFS(t *testing.T) {
	props := map[string]interface{}{
		"server":    "test_server",
		"server2":   "test_backup",
		"transport": "unix",
		"volume":    "test_volume",
	}

	requiredProps := []string{"server", "volume"}

	defaultProps := []string{"server2", "transport"}

	factoryFunc := func(props storage.ExtraProperties) (interface{}, error) {
		obj, err := storage.NewStorageGlusterFSProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			storageProps, err := storage.NewStorageGlusterFSProperties(props)
			require.NoError(t, err)

			assert.Equal(t, "test_server", storageProps.MainServer)
			assert.Equal(t, "test_backup", storageProps.BackupServer)
			assert.Equal(
				t,
				storage.GlusterFSTransportUNIX,
				storageProps.Transport,
			)
			assert.Equal(t, "test_volume", storageProps.Volume)
		})

	t.Run(
		"RequiredProperties", helperTestRequiredProperties(t, props, requiredProps, factoryFunc))

	t.Run("DefaultProperties", helperTestOptionalProperties(t, props, defaultProps, factoryFunc, func(obj interface{}) {
		require.IsType(t, (*storage.StorageGlusterFSProperties)(nil), obj)

		storageProps := obj.(*storage.StorageGlusterFSProperties)

		assert.Equal(
			t,
			storage.DefaultStorageGlusterFSBackupServer,
			storageProps.BackupServer,
		)

		assert.Equal(
			t,
			storage.DefaultStorageGlusterFSTransport,
			storageProps.Transport,
		)
	},
	))
}

func TestGlusterFSTransport(t *testing.T) {
	var GlusterFSTransportCases = map[string](struct {
		Object storage.GlusterFSTransport
		Value  string
	}){
		"TCP": {
			Object: storage.GlusterFSTransportTCP,
			Value:  "tcp",
		},
		"UNIX": {
			Object: storage.GlusterFSTransportUNIX,
			Value:  "unix",
		},
		"RDMA": {
			Object: storage.GlusterFSTransportRDMA,
			Value:  "rdma",
		},
	}

	t.Run("Marshal", func(t *testing.T) {

		for n, tt := range GlusterFSTransportCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				var receivedTransport storage.GlusterFSTransport
				err := (&receivedTransport).Unmarshal(tt.Value)
				require.NoError(t, err)
				assert.Equal(t, tt.Object, receivedTransport)
			})
		}
	})
	t.Run("Unarshal", func(t *testing.T) {
		for n, tt := range GlusterFSTransportCases {
			tt := tt
			t.Run(n, func(t *testing.T) {
				receivedValue, err := tt.Object.Marshal()
				require.NoError(t, err)
				assert.Equal(t, tt.Value, receivedValue)
			})
		}
	})
}
