package service

import (
	"net/url"

	"github.com/xabinapal/gopve/internal"
)

type NodeServiceProvider interface {
	List() (NodeList, error)
	Get(string) (*Node, error)
	Reboot(node string) error
	Shutdown(node string) error
}

type NodeService struct {
	client *internal.Client

	qemuFactory QEMUServiceFactoryProvider
	lxcFactory  LXCServiceFactoryProvider
	taskFactory TaskServiceFactoryProvider
}

func NewNodeService(c *internal.Client) *NodeService {
	node := &NodeService{client: c}
	node.qemuFactory = NewQEMUServiceFactoryProvider(c)
	node.lxcFactory = NewLXCServiceFactoryProvider(c)
	node.taskFactory = NewTaskServiceFactoryProvider(c)
	return node
}

func (s *NodeService) List() (NodeList, error) {
	data, err := s.client.Get("nodes")
	if err != nil {
		return nil, err
	}

	res := make(NodeList)
	for _, node := range data.(internal.JArray) {
		val := node.(internal.JObject)
		row := &Node{provider: s}
		internal.JSONToStruct(val, row)
		row.qemu = s.qemuFactory.Create(row)
		row.lxc = s.lxcFactory.Create(row)
		row.task = s.taskFactory.Create(row)
		res[row.Node] = row
	}

	return res, nil
}

func (s *NodeService) Get(node string) (*Node, error) {
	list, err := s.List()
	if err != nil {
		return nil, err
	}

	res, exist := list[node]
	if !exist {
		return nil, &NodeError{Node: node}
	}

	return res, nil
}

func (s *NodeService) power(node string, command string) error {
	params := &url.Values{"command": {command}}
	_, err := s.client.Post("nodes/" + node + "/status", params)
	return err
}

func (s *NodeService) Reboot(node string) error {
	return s.power(node, "reboot")
}

func (s *NodeService) Shutdown(node string) error {
	return s.power(node, "shutdown")
}
