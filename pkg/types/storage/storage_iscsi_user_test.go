package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

func TestStorageISCSIUser(t *testing.T) {
	props := map[string]interface{}{
		"portal": "test_portal",
		"target": "test_target",
	}

	requiredProps := []string{"portal", "target"}

	factoryFunc := func(props storage.ExtraProperties) (interface{}, error) {
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
		"RequiredProperties", helperTestRequiredProperties(t, props, requiredProps, factoryFunc))
}
