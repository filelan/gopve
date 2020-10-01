package storage_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types/storage"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func helperTestRequiredProperties(
	t *testing.T,
	props storage.ExtraProperties,
	requiredProps []string,
	factoryFunc func(storage.ExtraProperties) (interface{}, error),
) func(t *testing.T) {
	t.Helper()

	return func(t *testing.T) {
		for _, prop := range requiredProps {
			finalProps := make(storage.ExtraProperties, len(props))
			for k, v := range props {
				if k != prop {
					finalProps[k] = v
				}
			}

			expectedError := storage.ErrMissingProperty
			expectedError.AddKey("name", prop)

			obj, err := factoryFunc(finalProps)

			assert.Nil(t, obj)
			assert.EqualError(t, err, expectedError.Error())
		}
	}
}

func helperTestOptionalProperties(
	t *testing.T,
	props storage.ExtraProperties,
	optionalProps []string,
	factoryFunc func(storage.ExtraProperties) (interface{}, error),
	testFunc func(interface{}),
) func(t *testing.T) {
	t.Helper()

	return func(t *testing.T) {
		finalProps := make(storage.ExtraProperties)
		for k, v := range props {
			finalProps[k] = v
		}

		for _, v := range optionalProps {
			delete(finalProps, v)
		}

		obj, err := factoryFunc(finalProps)
		require.NoError(t, err)

		testFunc(obj)
	}
}
