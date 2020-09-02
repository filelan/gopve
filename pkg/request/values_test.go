package request_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/request"
)

func helpValuesContainsKVPair(t *testing.T, values request.Values, key, value string) {
	assert.Contains(t, values, key)
	assert.Len(t, values[key], 1)
	assert.Equal(t, []string{value}, values[key])
}

func TestValuesAddValue(t *testing.T) {
	values := request.Values{}

	values.AddString("stringKey", "stringValue")
	helpValuesContainsKVPair(t, values, "stringKey", "stringValue")

	values.AddInt("intKey", -1)
	helpValuesContainsKVPair(t, values, "intKey", "-1")

	values.AddUint("uintKey", 1)
	helpValuesContainsKVPair(t, values, "uintKey", "1")

	values.AddBool("tueKey", true)
	helpValuesContainsKVPair(t, values, "tueKey", "1")

	values.AddBool("falseKey", false)
	helpValuesContainsKVPair(t, values, "falseKey", "0")

	loc, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	values.AddTime("timeKey", time.Unix(1609458356, 0).In(loc))
	helpValuesContainsKVPair(t, values, "timeKey", "2020-12-31 23:45:56")
}

func TestValuesConditionalAddValue(t *testing.T) {
	values := request.Values{}

	values.ConditionalAddString("stringKey", "stringValue", true)
	helpValuesContainsKVPair(t, values, "stringKey", "stringValue")

	values.ConditionalAddInt("intKey", -1, true)
	helpValuesContainsKVPair(t, values, "intKey", "-1")

	values.ConditionalAddUint("uintKey", 1, true)
	helpValuesContainsKVPair(t, values, "uintKey", "1")

	values.ConditionalAddBool("trueKey", true, true)
	helpValuesContainsKVPair(t, values, "trueKey", "1")

	values.ConditionalAddBool("falseKey", false, true)
	helpValuesContainsKVPair(t, values, "falseKey", "0")

	loc, err := time.LoadLocation("UTC")
	require.NoError(t, err)

	values.ConditionalAddTime("timeKey", time.Unix(1609458356, 0).In(loc), true)
	helpValuesContainsKVPair(t, values, "timeKey", "2020-12-31 23:45:56")
}

func TestValuesConditionalNotAddValue(t *testing.T) {
	values := request.Values{}

	values.ConditionalAddString("stringKey", "stringValue", false)
	assert.NotContains(t, values, "stringKey")

	values.ConditionalAddInt("intKey", -1, false)
	assert.NotContains(t, values, "intKey")

	values.ConditionalAddUint("uintKey", 1, false)
	assert.NotContains(t, values, "uintKey")

	values.ConditionalAddBool("trueKey", true, false)
	assert.NotContains(t, values, "trueKey")

	values.ConditionalAddBool("falseKey", false, false)
	assert.NotContains(t, values, "falseKey")

	values.ConditionalAddTime("timeKey", time.Unix(1609458356, 0), false)
	assert.NotContains(t, values, "timeKey")
}
