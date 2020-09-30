package storage_test

import (
	"testing"

	"github.com/xabinapal/gopve/pkg/types/storage"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func helperTestOptionalProperties(
	t *testing.T,
	props storage.ExtraProperties,
	optionalProps []string,
	unmarshalFunc func(props storage.ExtraProperties) error,
) {
	t.Helper()

	finalProps := make(storage.ExtraProperties)
	for k, v := range props {
		finalProps[k] = v
	}

	for _, v := range optionalProps {
		delete(finalProps, v)
	}

	err := unmarshalFunc(finalProps)
	require.NoError(t, err)
}

func helperTestRequiredProperties(
	t *testing.T,
	props storage.ExtraProperties,
	requiredProps []string,
	unmarshalFunc func(props storage.ExtraProperties) error,
) {
	t.Helper()

	for _, prop := range requiredProps {
		finalProps := make(storage.ExtraProperties, len(props))
		for k, v := range props {
			if k != prop {
				finalProps[k] = v
			}
		}

		err := unmarshalFunc(finalProps)

		expectedError := storage.ErrMissingProperty
		expectedError.AddKey("name", prop)

		assert.EqualError(t, err, expectedError.Error())
	}
}
