package qemu

import (
	"encoding/json"
)

type BIOSType string

const (
	BIOSTypeSeaBIOS BIOSType = "seabios"
	BIOSTypeOVMF    BIOSType = "ovmf"
)

func (obj BIOSType) IsValid() bool {
	switch obj {
	case BIOSTypeSeaBIOS, BIOSTypeOVMF:
		return true
	default:
		return false
	}
}

func (obj BIOSType) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj BIOSType) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *BIOSType) Unmarshal(s string) error {
	*obj = BIOSType(s)
	return nil
}

func (obj *BIOSType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
