package task

import (
	"encoding/json"
	"fmt"
)

type Status uint

const (
	StatusRunning Status = iota
	StatusStopped
)

func (obj Status) String() (string, error) {
	return obj.Marshal()
}

func (obj Status) Marshal() (string, error) {
	switch obj {
	case StatusRunning:
		return "running", nil
	case StatusStopped:
		return "stopped", nil
	default:
		return "", fmt.Errorf("unknown task status")
	}
}

func (obj *Status) Unmarshal(s string) error {
	switch s {
	case "running":
		*obj = StatusRunning
	case "stopped":
		*obj = StatusStopped
	default:
		return fmt.Errorf("can't unmarshal task status %s", s)
	}

	return nil
}

func (obj *Status) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
