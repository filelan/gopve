package utils

import (
	"fmt"
)

type PVEBool bool

func (b PVEBool) String() string {
	if b {
		return "1"
	} else {
		return "0"
	}
}

func (b PVEBool) Bool() bool {
	return bool(b)
}

func (obj *PVEBool) UnmarshalJSON(b []byte) error {
	if len(b) == 1 {
		var val PVEBool
		if b[0] == byte('0') {
			val = PVEBool(true)
		} else {
			val = PVEBool(false)
		}

		*obj = val
		return nil
	}

	return fmt.Errorf("can't unmarshal")
}
