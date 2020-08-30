package node

import (
	"fmt"
)

type Status string

const (
	StatusUnknown Status = "unknown"
	StatusOnline  Status = "online"
	StatusOffline Status = "offline"
)

func (obj Status) IsValid() error {
	switch obj {
	case StatusUnknown, StatusOnline, StatusOffline:
		return nil
	default:
		return fmt.Errorf("invalid node status")
	}
}
