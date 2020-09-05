package test

import (
	"github.com/xabinapal/gopve/internal/service/node"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func NewNode() (*node.Node, *mocks.Executor) {
	svc, exc := NewService()
	return node.NewNode(svc, "test_node"), exc
}
