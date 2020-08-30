package vm

import (
	"fmt"
)

type Status string

const (
	StatusRunning Status = "running"
	StatusStopped Status = "stopped"
)

func (obj Status) IsValid() error {
	switch obj {
	case StatusRunning, StatusStopped:
		return nil
	default:
		return fmt.Errorf("invalid virtual machine status")
	}
}
