package types

import (
	"bytes"
	"fmt"
)

type PVEBool bool

func NewPVEBoolFromInt(x int) PVEBool {
	if x == 0 {
		return PVEBool(false)
	}

	return PVEBool(true)
}

func NewPVEBoolFromFloat64(x float64) PVEBool {
	if x == 0 {
		return PVEBool(false)
	}

	return PVEBool(true)
}

func NewPVEBoolFromString(x string) (PVEBool, error) {
	switch x {
	case "0", "no", "off":
		return PVEBool(false), nil
	case "1", "yes", "on":
		return PVEBool(true), nil
	default:
		return PVEBool(false), fmt.Errorf("unknown value")
	}
}

func (obj PVEBool) String() string {
	if obj {
		return "1"
	}

	return "0"
}

func (obj PVEBool) Bool() bool {
	return bool(obj)
}

func (obj *PVEBool) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("1")) {
		*obj = PVEBool(true)
	} else if bytes.Equal(b, []byte("0")) || bytes.Equal(b, []byte("\"\"")) {
		*obj = PVEBool(false)
	} else {
		return fmt.Errorf("unknown boolean value %s", string(b))
	}

	return nil
}
