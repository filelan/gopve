package node

import (
	"github.com/xabinapal/gopve/internal/client"
	"github.com/xabinapal/gopve/pkg/types/node"
)

type Service struct {
	client client.Client
	api    client.API
}

func NewService(cli client.Client, api client.API) *Service {
	return &Service{
		client: cli,
		api:    api,
	}
}

type Node struct {
	svc    *Service
	name   string
	status node.Status
}

func NewNode(svc *Service, name string) *Node {
	return &Node{
		svc:  svc,
		name: name,
	}
}

func (n *Node) Name() string {
	return n.name
}

func (n *Node) Status() node.Status {
	return n.status
}
