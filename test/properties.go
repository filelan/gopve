package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

func HelperCreatePropertiesMap(props types.Properties) types.Properties {
	obj := make(types.Properties, len(props))

	for k, v := range props {
		switch x := v.(type) {
		case string:
			obj[k] = x
		case int:
			obj[k] = float64(x)
		case int8:
			obj[k] = float64(x)
		case int16:
			obj[k] = float64(x)
		case int32:
			obj[k] = float64(x)
		case int64:
			obj[k] = float64(x)
		case uint:
			obj[k] = float64(x)
		case uint8:
			obj[k] = float64(x)
		case uint16:
			obj[k] = float64(x)
		case uint32:
			obj[k] = float64(x)
		case uint64:
			obj[k] = float64(x)
		case float32:
			obj[k] = float64(x)
		case float64:
			obj[k] = x
		default:
			panic("unsupported type")
		}
	}

	return obj
}

func HelperTestRequiredProperties(
	t *testing.T,
	props types.Properties,
	requiredProps []string,
	factoryFunc func(types.Properties) (interface{}, error),
) func(t *testing.T) {
	t.Helper()

	return func(t *testing.T) {
		for _, prop := range requiredProps {
			finalProps := make(types.Properties, len(props))
			for k, v := range props {
				if k != prop {
					finalProps[k] = v
				}
			}

			expectedError := errors.ErrMissingProperty
			expectedError.AddKey("name", prop)

			_, err := factoryFunc(finalProps)
			assert.EqualError(t, err, expectedError.Error())
		}
	}
}

func HelperTestOptionalProperties(
	t *testing.T,
	props types.Properties,
	optionalProps []string,
	factoryFunc func(types.Properties) (interface{}, error),
	testFunc func(interface{}),
) func(t *testing.T) {
	t.Helper()

	return func(t *testing.T) {
		finalProps := make(types.Properties)
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
