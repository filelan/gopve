package types

import (
	"fmt"
	"time"
)

type NodeService interface {
	GetAll() ([]Node, error)
	Get(node string) (Node, error)
}

type Node interface {
	Name() string
	Status() NodeStatus

	Shutdown() error
	Reboot() error
	WakeOnLAN() (Task, error)

	GetSyslog(opts NodeGetSyslogOptions) (LogEntries, error)

	GetTime(local bool) (*time.Time, error)
	GetTimezone() (*time.Location, error)
	SetTimezone(timezone *time.Location) error
}

type NodeStatus string

const (
	NodeUnknown NodeStatus = "unknown"
	NodeOnline  NodeStatus = "online"
	NodeOffline NodeStatus = "offline"
)

func (obj NodeStatus) IsValid() error {
	switch obj {
	case NodeUnknown, NodeOnline, NodeOffline:
		return nil
	default:
		return fmt.Errorf("invalid node status")
	}
}

type NodeGetSyslogOptions struct {
	LineStart uint
	LineLimit uint

	Since time.Time
	Until time.Time

	Service string
}
