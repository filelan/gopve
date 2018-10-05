package service

import (
	"github.com/xabinapal/gopve/internal"
)

type NodeServiceProvider interface {
	List() (*NodeList, error)
}

type NodeService struct {
	client *internal.Client

	QEMU QEMUServiceProvider
	LXC  LXCServiceProvider
}

type Node struct {
	Node        string
	Status      string
	Uptime      int
	CPUTotal    int
	CPUCurrent  float64
	MemTotal    int
	MemCurrent  int
	DiskTotal   int
	DiskCurrent int
}

type NodeList []Node

func NewNodeService(c *internal.Client) *NodeService {
	node := &NodeService{client: c}
	node.QEMU = NewQEMUService(c, node)
	node.LXC = NewLXCService(c, node)
	return node
}

func (s *NodeService) List() (*NodeList, error) {
	data, err := s.client.Get("nodes")
	if err != nil {
		return nil, err
	}

	var nodes NodeList
	for _, node := range data.([]interface{}) {
		value := node.(map[string]interface{})
		nodes = append(nodes, Node{
			Node:        value["node"].(string),
			Status:      value["status"].(string),
			Uptime:      int(value["uptime"].(float64)),
			CPUTotal:    int(value["maxcpu"].(float64)),
			CPUCurrent:  value["cpu"].(float64),
			MemTotal:    int(value["maxmem"].(float64)),
			MemCurrent:  int(value["mem"].(float64)),
			DiskTotal:   int(value["maxdisk"].(float64)),
			DiskCurrent: int(value["disk"].(float64)),
		})
	}

	return &nodes, nil
}