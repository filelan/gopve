package qemu

import (
	"encoding/json"
)

type CPUArchitecture string

const (
	CPUArchitectureHost  CPUArchitecture = ""
	CPUArchitectureAMD64 CPUArchitecture = "x86_64"
	CPUArchitectureARM64 CPUArchitecture = "aarch64"
)

func (obj CPUArchitecture) IsValid() bool {
	switch obj {
	case CPUArchitectureHost,
		CPUArchitectureAMD64,
		CPUArchitectureARM64:
		return true
	default:
		return false
	}
}

func (obj CPUArchitecture) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj CPUArchitecture) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *CPUArchitecture) Unmarshal(s string) error {
	*obj = CPUArchitecture(s)
	return nil
}

func (obj *CPUArchitecture) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
