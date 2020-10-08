package vm

import (
	"encoding/json"
)

type QEMUBIOS string

const (
	QEMUBIOSSeaBIOS QEMUBIOS = "seabios"
	QEMUBIOSOVMF    QEMUBIOS = "ovmf"
)

func (obj QEMUBIOS) IsValid() bool {
	switch obj {
	case QEMUBIOSSeaBIOS, QEMUBIOSOVMF:
		return true
	default:
		return false
	}
}

func (obj QEMUBIOS) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj QEMUBIOS) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *QEMUBIOS) Unmarshal(s string) error {
	*obj = QEMUBIOS(s)
	return nil
}

func (obj *QEMUBIOS) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
