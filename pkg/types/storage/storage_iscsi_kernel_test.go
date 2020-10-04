package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/test"
)

func TestStorageISCSIKernelProperties(t *testing.T) {
	props := test.HelperCreatePropertiesMap(types.Properties{
		"portal": "test_portal",
		"target": "test_target",
	})

	requiredProps := []string{"portal", "target"}

	factoryFunc := func(props types.Properties) (interface{}, error) {
		obj, err := storage.NewStorageISCSIKernelProperties(props)
		return obj, err
	}

	t.Run(
		"Create", func(t *testing.T) {
			storageProps, err := storage.NewStorageISCSIKernelProperties(props)
			require.NoError(t, err)

			assert.Equal(t, "test_portal", storageProps.Portal)
			assert.Equal(t, "test_target", storageProps.Target)
		})

	t.Run(
		"RequiredProperties",
		test.HelperTestRequiredProperties(t, props, requiredProps, factoryFunc),
	)
}
