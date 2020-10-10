package types

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type PVEKeyValue struct {
	Separator    string
	AllowNoValue bool

	key      string
	value    string
	hasValue bool
}

func (obj PVEKeyValue) Key() string {
	return obj.key
}

func (obj PVEKeyValue) Value() string {
	return obj.value
}

func (obj PVEKeyValue) ValueAsInt() (int, error) {
	if obj.HasValue() {
		return strconv.Atoi(obj.Value())
	}

	return 0, fmt.Errorf("no value")
}

func (obj PVEKeyValue) ValueAsBool() (bool, error) {
	if obj.HasValue() {
		if b, err := NewPVEBoolFromString(obj.Value()); err == nil {
			return b.Bool(), nil
		} else {
			return false, err
		}
	}

	return false, fmt.Errorf("no value")
}

func (obj PVEKeyValue) HasValue() bool {
	return obj.hasValue
}

func (obj PVEKeyValue) Marshal() (string, error) {
	return fmt.Sprintf("%s%s%s", obj.key, obj.Separator, obj.value), nil
}

func (obj *PVEKeyValue) Unmarshal(s string) error {
	if obj.Separator == "" {
		return fmt.Errorf("can't unmarshal kv, no separator defined")
	}

	content := strings.Split(s, obj.Separator)

	switch len(content) {
	case 1:
		if !obj.AllowNoValue {
			return fmt.Errorf("can't unmarshal key %s", s)
		}

		obj.key = content[0]
		obj.hasValue = false

	case 2:
		obj.key = content[0]
		obj.value = content[1]
		obj.hasValue = true

	default:
		return fmt.Errorf("can't unmarshal key %s", s)
	}

	return nil
}

func (obj *PVEKeyValue) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
