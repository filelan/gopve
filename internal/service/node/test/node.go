package test

import (
	"github.com/xabinapal/gopve/internal/service/node"
	"github.com/xabinapal/gopve/pkg/request/mocks"
	types "github.com/xabinapal/gopve/pkg/types/node"
)

func NewNode() (*node.Node, *mocks.Executor) {
	svc, exc := NewService()
	return node.NewNode(svc, "test_node", types.StatusOnline), exc
}
