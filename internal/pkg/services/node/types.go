package node

import (
	"github.com/xabinapal/gopve/internal/pkg/utils"
	"github.com/xabinapal/gopve/pkg/types"
)

type Service struct {
	utils.Service
}

func NewService(api utils.API, client utils.Client) *Service {
	return &Service{utils.Service{api, client}}
}

type Node struct {
	svc *Service

	name   string
	status types.NodeStatus
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) Status() types.NodeStatus {
	return n.status
}
