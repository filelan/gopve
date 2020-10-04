package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/test"
)

func TestStorageISCSIUserProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"portal": "test_portal",
		"target": "test_target",
	})

	requiredProps := []string{"portal", "target"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := storage.NewStorageISCSIUserProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			storageProps, err := storage.NewStorageISCSIUserProperties(props)
			require.NoError(t, err)

			assert.Equal(t, "test_portal", storageProps.Portal)
			assert.Equal(t, "test_target", storageProps.Target)
		})

	t.Run(
		"RequiredProperties",
		test.HelperTestRequiredProperties(t, props, requiredProps, factoryFunc),
	)
}
