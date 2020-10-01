package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

func TestStorageDRBD(t *testing.T) {
	props := map[string]interface{}{"redundancy": 16}

	defaultProps := []string{"redundancy"}

	factoryFunc := func(props storage.ExtraProperties) (interface{}, error) {
		obj, err := storage.NewStorageDRBDProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			storageProps, err := storage.NewStorageDRBDProperties(props)
			require.NoError(t, err)

			assert.Equal(t, uint(16), storageProps.Redundancy)
		})

	t.Run("DefaultProperties", helperTestOptionalProperties(t, props, defaultProps, factoryFunc, func(obj interface{}) {
		require.IsType(t, (*storage.StorageDRBDProperties)(nil), obj)

		storageProps := obj.(*storage.StorageDRBDProperties)

		assert.Equal(
			t,
			storage.DefaultStorageDRBDRedundancy,
			storageProps.Redundancy,
		)
	},
	))
}
