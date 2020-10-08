package test

import (
	"encoding/json"
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/types"
)

func HelperTestFixedValue(t *testing.T, value types.FixedValuePtr, cases map[string]struct {
	Object types.FixedValue
	Value  string
}) {
	t.Helper()

	r := reflect.ValueOf(value)
	require.True(t, r.Kind() == reflect.Ptr)

	require.Implements(t, (*types.Marshaler)(nil), value)
	require.Implements(t, (*types.Unmarshaler)(nil), value)
	require.Implements(t, (*json.Unmarshaler)(nil), value)

	val := reflect.New(r.Type().Elem()).Elem()
	require.Implements(t, (*types.Marshaler)(nil), val.Interface())

	for n, tt := range cases {
		tt := tt

		t.Run(n, func(t *testing.T) {
			require.IsType(t, val.Interface(), tt.Object)

			t.Run("Valid", func(t *testing.T) {
				valid := tt.Object.IsValid()
				assert.True(t, valid)

				unknown := tt.Object.IsUnknown()
				assert.False(t, unknown)
			})

			t.Run("Marshal", func(t *testing.T) {
				receivedValue, err := tt.Object.Marshal()
				require.NoError(t, err)
				assert.Equal(t, tt.Value, receivedValue)
			})

			t.Run("Unmarshal", func(t *testing.T) {
				ptr := reflect.New(val.Type()).Interface().(types.FixedValuePtr)

				err := ptr.Unmarshal(tt.Value)
				require.NoError(t, err)

				v := reflect.ValueOf(ptr).Elem().Interface()
				assert.Equal(t, tt.Object, v)
			})

			t.Run("UnmarshalJSON", func(t *testing.T) {
				ptr := reflect.New(val.Type()).Interface().(types.FixedValuePtr)

				err := ptr.UnmarshalJSON([]byte(fmt.Sprintf("\"%s\"", tt.Value)))
				require.NoError(t, err)

				v := reflect.ValueOf(ptr).Elem().Interface()
				assert.Equal(t, tt.Object, v)
			})
		})
	}

	t.Run("Unknown", func(t *testing.T) {
		t.Run("Valid", func(t *testing.T) {
			v := reflect.New(val.Type()).Elem()
			v.SetString("test_unknown_value")

			vv := v.Interface().(types.FixedValue)

			valid := vv.IsValid()
			assert.False(t, valid)

			unknown := vv.IsUnknown()
			assert.True(t, unknown)
		})
	})

	t.Run("UnmarshalJSONError", func(t *testing.T) {
		v := reflect.New(val.Type())

		vv := v.Interface().(types.FixedValuePtr)

		err := (vv).UnmarshalJSON([]byte{})
		assert.Error(t, err)
	})
}
