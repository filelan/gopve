package vm

import (
	"encoding/json"
)

type QEMUCPUArchitecture string

const (
	QEMUCPUArchitectureHost    QEMUCPUArchitecture = ""
	QEMUCPUArchitectureX86_64  QEMUCPUArchitecture = "x86_64"
	QEMUCPUArchitectureAArch64 QEMUCPUArchitecture = "aarch64"
)

func (obj QEMUCPUArchitecture) IsValid() bool {
	switch obj {
	case QEMUCPUArchitectureHost, QEMUCPUArchitectureX86_64, QEMUCPUArchitectureAArch64:
		return true
	default:
		return false
	}
}

func (obj QEMUCPUArchitecture) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj QEMUCPUArchitecture) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *QEMUCPUArchitecture) Unmarshal(s string) error {
	*obj = QEMUCPUArchitecture(s)
	return nil
}

func (obj *QEMUCPUArchitecture) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}

type LXCCPUArchitecture string

const (
	LXCCPUArchitectureAMD64 LXCCPUArchitecture = "amd64"
	LXCCPUArchitectureI386  LXCCPUArchitecture = "i386"
	LXCCPUArchitectureARM64 LXCCPUArchitecture = "arm64"
	LXCCPUArchitectureARMHF LXCCPUArchitecture = "armhf"
)

func (obj LXCCPUArchitecture) IsValid() bool {
	switch obj {
	case LXCCPUArchitectureAMD64, LXCCPUArchitectureI386, LXCCPUArchitectureARM64, LXCCPUArchitectureARMHF:
		return true
	default:
		return false
	}
}

func (obj LXCCPUArchitecture) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj LXCCPUArchitecture) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *LXCCPUArchitecture) Unmarshal(s string) error {
	*obj = LXCCPUArchitecture(s)
	return nil
}

func (obj *LXCCPUArchitecture) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
