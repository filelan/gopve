package service

import (
	"github.com/xabinapal/gopve/pkg/types/node"
)

//go:generate mockery --case snake --name Node

type Node interface {
	List() ([]node.Node, error)
	Get(node string) (node.Node, error)
}
