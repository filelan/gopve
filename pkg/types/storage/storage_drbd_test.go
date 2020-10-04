package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/test"
)

func TestStorageDRBDProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{"redundancy": 16})

	defaultProps := []string{"redundancy"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := storage.NewStorageDRBDProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			storageProps, err := storage.NewStorageDRBDProperties(props)
			require.NoError(t, err)

			assert.Equal(t, uint(16), storageProps.Redundancy)
		})

	t.Run(
		"DefaultProperties",
		test.HelperTestOptionalProperties(
			t,
			props,
			defaultProps,
			factoryFunc,
			func(obj interface{}) {
				require.IsType(t, (*storage.StorageDRBDProperties)(nil), obj)

				storageProps := obj.(*storage.StorageDRBDProperties)

				assert.Equal(
					t,
					storage.DefaultStorageDRBDRedundancy,
					storageProps.Redundancy,
				)
			},
		),
	)
}
