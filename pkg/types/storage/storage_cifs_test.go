package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/test"
)

func TestStorageCIFSProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"server":     "test_server",
		"smbversion": "2.1",
		"domain":     "test_domain",
		"username":   "test_username",
		"password":   "test_password",
		"share":      "test_share",
		"path":       "/test_path",
		"mkdir":      1,
	})

	requiredProps := []string{"server", "share", "path"}

	defaultProps := []string{
		"smbversion",
		"domain",
		"username",
		"password",
		"mkdir",
	}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := storage.NewStorageCIFSProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			storageProps, err := storage.NewStorageCIFSProperties(props)
			require.NoError(t, err)

			assert.Equal(t, "test_server", storageProps.Server)
			assert.Equal(t, storage.SMBVersion21, storageProps.SMBVersion)
			assert.Equal(t, "test_domain", storageProps.Domain)
			assert.Equal(t, "test_username", storageProps.Username)
			assert.Equal(t, "test_password", storageProps.Password)
			assert.Equal(t, "test_share", storageProps.ServerShare)
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
				require.IsType(t, (*storage.StorageCIFSProperties)(nil), obj)

				storageProps := obj.(*storage.StorageCIFSProperties)

				assert.Equal(
					t,
					storage.DefaultStorageCIFSSMBVersion,
					storageProps.SMBVersion,
				)

				assert.Equal(
					t,
					storage.DefaultStorageCIFSDomain,
					storageProps.Domain,
				)

				assert.Equal(
					t,
					storage.DefaultStorageCIFSUsername,
					storageProps.Username,
				)

				assert.Equal(
					t,
					storage.DefaultStorageCIFSPassword,
					storageProps.Password,
				)

				assert.Equal(
					t,
					storage.DefaultStorageCIFSLocalPathCreate,
					storageProps.LocalPathCreate,
				)
			},
		),
	)
}

func TestSMBVersion(t *testing.T) {
	test.HelperTestFixedValue(
		t,
		(*storage.SMBVersion)(nil),
		map[string](struct {
			Object types.FixedValue
			Value  string
		}){
			"2.0": {
				Object: storage.SMBVersion20,
				Value:  "2.0",
			},
			"2.1": {
				Object: storage.SMBVersion21,
				Value:  "2.1",
			},
			"3.0": {
				Object: storage.SMBVersion30,
				Value:  "3.0",
			},
		},
	)
}
