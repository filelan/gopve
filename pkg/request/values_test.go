package request_test

import (
	"strings"
	"testing"
	"time"

	"github.com/xabinapal/gopve/pkg/request"
)

func helpValuesContainsKVPair(t *testing.T, values request.Values, key, value string) {
	for k, v := range values {
		if k == key {
			if len(v) != 1 {
				t.Errorf("Got multi-valued key '%s', value '%s'", key, strings.Join(v, ","))
			} else if v[0] != value {
				t.Errorf("Got key '%s', value '%s', expected '%s'", key, v[0], value)
			}

			return
		}
	}

	t.Errorf("Key '%s' not found", key)
}

func helpValuesNotContainsKVPair(t *testing.T, values request.Values, key string) {
	for k := range values {
		if k == key {
			t.Errorf("Key '%s' found", key)
		}
	}
}

func TestValuesAddValue(t *testing.T) {
	values := request.Values{}

	values.AddString("stringKey", "stringValue")
	helpValuesContainsKVPair(t, values, "stringKey", "stringValue")

	values.AddInt("intKey", -1)
	helpValuesContainsKVPair(t, values, "intKey", "-1")

	values.AddUint("uintKey", 1)
	helpValuesContainsKVPair(t, values, "uintKey", "1")

	values.AddBool("boolTrueKey", true)
	helpValuesContainsKVPair(t, values, "boolTrueKey", "1")

	values.AddBool("boolFalseKey", false)
	helpValuesContainsKVPair(t, values, "boolFalseKey", "0")

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		t.Fatalf("Unexpected time.LoadLocation error: %s", err.Error())
	}

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

	values.ConditionalAddBool("boolTrueKey", true, true)
	helpValuesContainsKVPair(t, values, "boolTrueKey", "1")

	values.ConditionalAddBool("boolFalseKey", false, true)
	helpValuesContainsKVPair(t, values, "boolFalseKey", "0")

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		t.Fatalf("Unexpected time.LoadLocation error: %s", err.Error())
	}

	values.ConditionalAddTime("timeKey", time.Unix(1609458356, 0).In(loc), true)
	helpValuesContainsKVPair(t, values, "timeKey", "2020-12-31 23:45:56")
}

func TestValuesConditionalNotAddValue(t *testing.T) {
	values := request.Values{}

	values.ConditionalAddString("stringKey", "stringValue", false)
	helpValuesNotContainsKVPair(t, values, "stringKey")

	values.ConditionalAddInt("intKey", -1, false)
	helpValuesNotContainsKVPair(t, values, "intKey")

	values.ConditionalAddUint("uintKey", 1, false)
	helpValuesNotContainsKVPair(t, values, "uintKey")

	values.ConditionalAddBool("boolTrueKey", true, false)
	helpValuesNotContainsKVPair(t, values, "boolTrueKey")

	values.ConditionalAddBool("boolFalseKey", false, false)
	helpValuesNotContainsKVPair(t, values, "boolFalseKey")

	values.ConditionalAddTime("timeKey", time.Unix(1609458356, 0), false)
	helpValuesNotContainsKVPair(t, values, "timeKey")
}
