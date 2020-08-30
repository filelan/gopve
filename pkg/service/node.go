package service

import (
	"github.com/xabinapal/gopve/pkg/types/node"
)

type Node interface {
	List() ([]node.Node, error)
	Get(node string) (node.Node, error)
}
