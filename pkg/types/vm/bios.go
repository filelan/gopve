package vm

import (
	"encoding/json"
	"fmt"
)

type QEMUBIOS uint

const (
	QEMUBIOSSeaBIOS QEMUBIOS = iota
	QEMUBIOSOVMF
)

func (obj QEMUBIOS) Marshal() (string, error) {
	switch obj {
	case QEMUBIOSSeaBIOS:
		return "seabios", nil
	case QEMUBIOSOVMF:
		return "ovmf", nil
	default:
		return "", fmt.Errorf("unknown bios")
	}
}

func (obj *QEMUBIOS) Unmarshal(s string) error {
	switch s {
	case "seabios":
		*obj = QEMUBIOSSeaBIOS
	case "ovmf":
		*obj = QEMUBIOSOVMF
	default:
		return fmt.Errorf("can't unmarshal bios %s", s)
	}

	return nil
}

func (obj *QEMUBIOS) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
