package service

import (
	"github.com/xabinapal/gopve/internal"
)

type NodeServiceProvider interface {
	List() (NodeList, error)
}

type NodeService struct {
	client *internal.Client

	QEMU QEMUServiceProvider
	LXC  LXCServiceProvider
}

type Node struct {
	Node string
}

type NodeList []Node

func NewNodeService(c *internal.Client) *NodeService {
	node := &NodeService{client: c}
	node.QEMU = NewQEMUService(c, node)
	node.LXC = NewLXCService(c, node)
	return node
}

func (s *NodeService) List() (NodeList, error) {
	data, err := s.client.Get("nodes")
	if err != nil {
		return nil, err
	}

	var nodes NodeList
	for _, node := range data {
		value := node.(map[string]interface{})
		nodes = append(nodes, Node{
			Node: value["name"].(string),
		})
	}

	return nodes, nil
}
