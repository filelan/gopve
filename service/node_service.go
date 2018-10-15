package service

import (
	"net/url"

	"github.com/xabinapal/gopve/internal"
)

type NodeServiceProvider interface {
	List() (*NodeList, error)
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

func (s *NodeService) List() (*NodeList, error) {
	data, err := s.client.Get("nodes")
	if err != nil {
		return nil, err
	}

	var res NodeList
	for _, node := range data.(internal.JArray) {
		val := node.(internal.JObject)
		row := &Node{provider: s}
		internal.JSONToStruct(val, row)
		row.QEMU = s.qemuFactory.Create(row)
		row.LXC = s.lxcFactory.Create(row)
		row.Task = s.taskFactory.Create(row)
		res = append(res, row)
	}

	return &res, nil
}

func (s *NodeService) Get(node string) (*Node, error) {
	data, err := s.client.Get("nodes/" + node + "/status")
	if err != nil {
		return nil, err
	}

	val := data.(internal.JObject)

	res := &Node{ provider: s, Node: node}
	internal.JSONToStruct(val, res)
	res.QEMU = s.qemuFactory.Create(res)
	res.LXC = s.lxcFactory.Create(res)
	res.Task = s.taskFactory.Create(res)
	return res, nil
}

func (s *NodeService) power(node string, command string) error {
	params := &url.Values{"command": {command}}
	_, err := s.client.Post("nodes/"+node+"/status", params)
	return err
}

func (s *NodeService) Reboot(node string) error {
	return s.power(node, "reboot")
}

func (s *NodeService) Shutdown(node string) error {
	return s.power(node, "shutdown")
}
