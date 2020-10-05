package vm

import (
	"encoding/json"
	"fmt"
)

type QEMUCPUArchitecture uint

const (
	QEMUCPUArchitectureHost QEMUCPUArchitecture = iota
	QEMUCPUArchitectureX86_64
	QEMUCPUArchitectureAArch64
)

func (obj QEMUCPUArchitecture) Marshal() (string, error) {
	switch obj {
	case QEMUCPUArchitectureX86_64:
		return "x86_64", nil
	case QEMUCPUArchitectureAArch64:
		return "aarch64", nil
	default:
		return "", fmt.Errorf("unknown qemu cpu architecture")
	}
}

func (obj *QEMUCPUArchitecture) Unmarshal(s string) error {
	switch s {
	case "x86_64":
		*obj = QEMUCPUArchitectureX86_64
	case "aarch64":
		*obj = QEMUCPUArchitectureAArch64
	default:
		return fmt.Errorf("can't unmarshal qemu cpu architecture %s", s)
	}

	return nil
}

func (obj *QEMUCPUArchitecture) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}

type LXCCPUArchitecture uint

const (
	LXCCPUArchitectureOther LXCCPUArchitecture = iota
	LXCCPUArchitectureAMD64
	LXCCPUArchitectureI386
	LXCCPUArchitectureARM64
	LXCCPUArchitectureARMHF
)

func (obj LXCCPUArchitecture) Marshal() (string, error) {
	switch obj {
	case LXCCPUArchitectureAMD64:
		return "amd64", nil
	case LXCCPUArchitectureI386:
		return "i386", nil
	case LXCCPUArchitectureARM64:
		return "arm64", nil
	case LXCCPUArchitectureARMHF:
		return "armhf", nil
	default:
		return "", fmt.Errorf("unknown lxc cpu architecture")
	}
}

func (obj *LXCCPUArchitecture) Unmarshal(s string) error {
	switch s {
	case "amd64":
		*obj = LXCCPUArchitectureAMD64
	case "i386":
		*obj = LXCCPUArchitectureI386
	case "arm64":
		*obj = LXCCPUArchitectureARM64
	case "armhf":
		*obj = LXCCPUArchitectureARMHF
	default:
		return fmt.Errorf("can't unmarshal lxc cpu architecture %s", s)
	}

	return nil
}

func (obj *LXCCPUArchitecture) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
