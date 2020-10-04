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
