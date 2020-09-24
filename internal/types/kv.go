package types

import (
	"encoding/json"
	"fmt"
	"strings"
)

const defaultSeparator = "="

type PVEStringKV struct {
	Separator    string
	AllowNoValue bool

	key      string
	value    string
	hasValue bool
}

func (obj PVEStringKV) Key() string {
	return obj.key
}

func (obj PVEStringKV) Value() string {
	return obj.value
}

func (obj PVEStringKV) HasValue() bool {
	return obj.hasValue
}

func (obj PVEStringKV) String() string {
	return fmt.Sprintf("%s%s%s", obj.key, obj.Separator, obj.value)
}

func (obj *PVEStringKV) Unmarshal(s string) error {
	separator := obj.Separator
	if separator == "" {
		separator = defaultSeparator
	}

	content := strings.Split(s, separator)

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

func (obj *PVEStringKV) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
