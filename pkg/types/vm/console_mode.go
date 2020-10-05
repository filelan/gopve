package vm

import (
	"encoding/json"
	"fmt"
)

type QEMUConsoleMode uint

const (
	QEMUConsoleModeUnknown QEMUConsoleMode = iota
	QEMUConsoleModeShell
	QEMUConsoleModeConsole
	QEMUConsoleModeTTY
)

func (obj QEMUConsoleMode) Marshal() (string, error) {
	switch obj {
	case QEMUConsoleModeShell:
		return "shell", nil
	case QEMUConsoleModeConsole:
		return "console", nil
	case QEMUConsoleModeTTY:
		return "tty", nil
	default:
		return "", fmt.Errorf("unknown qemu console mode")
	}
}

func (obj *QEMUConsoleMode) Unmarshal(s string) error {
	switch s {
	case "shell":
		*obj = QEMUConsoleModeShell
	case "console":
		*obj = QEMUConsoleModeConsole
	case "tty":
		*obj = QEMUConsoleModeTTY
	default:
		return fmt.Errorf("can't unmarshal qemu console mode %s", s)
	}

	return nil
}

func (obj *QEMUConsoleMode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
