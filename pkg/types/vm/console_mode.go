package vm

import (
	"encoding/json"
)

type QEMUConsoleMode string

const (
	QEMUConsoleModeShell   QEMUConsoleMode = "shell"
	QEMUConsoleModeConsole QEMUConsoleMode = "console"
	QEMUConsoleModeTTY     QEMUConsoleMode = "tty"
)

func (obj QEMUConsoleMode) IsValid() bool {
	switch obj {
	case QEMUConsoleModeShell, QEMUConsoleModeConsole, QEMUConsoleModeTTY:
		return true
	default:
		return false
	}
}

func (obj QEMUConsoleMode) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj QEMUConsoleMode) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *QEMUConsoleMode) Unmarshal(s string) error {
	*obj = QEMUConsoleMode(s)
	return nil
}

func (obj *QEMUConsoleMode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
