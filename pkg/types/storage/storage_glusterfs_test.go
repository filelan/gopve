package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/test"
)

func TestStorageGlusterFSProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"server":    "test_server",
		"server2":   "test_backup",
		"transport": "unix",
		"volume":    "test_volume",
	})

	requiredProps := []string{"server", "volume"}

	defaultProps := []string{"server2", "transport"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
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
					(*storage.StorageGlusterFSProperties)(nil),
					obj,
				)

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
		),
	)
}

func TestGlusterFSTransport(t *testing.T) {
	test.HelperTestFixedValue(
		t,
		(*storage.GlusterFSTransport)(nil),
		map[string](struct {
			Object types.FixedValue
			Value  string
		}){
			"None": {
				Object: storage.GlusterFSTransportNone,
				Value:  "",
			},
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
		},
	)
}
