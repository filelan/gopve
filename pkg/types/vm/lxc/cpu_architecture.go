package lxc

import (
	"encoding/json"
)

type CPUArchitecture string

const (
	CPUArchitectureI386  CPUArchitecture = "i386"
	CPUArchitectureAMD64 CPUArchitecture = "amd64"
	CPUArchitectureARM64 CPUArchitecture = "arm64"
	CPUArchitectureARMHF CPUArchitecture = "armhf"
)

func (obj CPUArchitecture) IsValid() bool {
	switch obj {
	case CPUArchitectureI386,
		CPUArchitectureAMD64,
		CPUArchitectureARM64,
		CPUArchitectureARMHF:
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
