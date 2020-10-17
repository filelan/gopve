package qemu

import (
	"encoding/json"
)

type ConsoleMode string

const (
	ConsoleModeShell   ConsoleMode = "shell"
	ConsoleModeConsole ConsoleMode = "console"
	ConsoleModeTTY     ConsoleMode = "tty"
)

func (obj ConsoleMode) IsValid() bool {
	switch obj {
	case ConsoleModeShell, ConsoleModeConsole, ConsoleModeTTY:
		return true
	default:
		return false
	}
}

func (obj ConsoleMode) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj ConsoleMode) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *ConsoleMode) Unmarshal(s string) error {
	*obj = ConsoleMode(s)
	return nil
}

func (obj *ConsoleMode) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
