package types

import (
	"fmt"
)

type NodeService interface {
	GetAll() ([]Node, error)
	Get(node string) (Node, error)
}

type Node interface {
	Name() string
	Status() NodeStatus
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
