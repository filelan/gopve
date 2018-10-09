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
	for _, node := range data.([]interface{}) {
		val := node.(map[string]interface{})
		row := &Node{
			provider: s,

			Node:          val["node"].(string),
			Status:        val["status"].(string),
			Uptime:        int(val["uptime"].(float64)),
			CPUTotal:      int(val["maxcpu"].(float64)),
			CPUPercentage: val["cpu"].(float64),
			MemTotal:      int(val["maxmem"].(float64)),
			MemUsed:       int(val["mem"].(float64)),
			DiskTotal:     int(val["maxdisk"].(float64)),
			DiskUsed:      int(val["disk"].(float64)),
		}

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

	val := data.(map[string]interface{})
	cpu := val["cpuinfo"].(map[string]interface{})
	mem := val["memory"].(map[string]interface{})
	disk := val["rootfs"].(map[string]interface{})

	res := &Node{
		provider: s,

		Node:          node,
		Status:        "",
		Uptime:        int(val["uptime"].(float64)),
		CPUTotal:      int(cpu["cpus"].(float64)),
		CPUPercentage: val["cpu"].(float64),
		MemTotal:      int(mem["total"].(float64)),
		MemUsed:       int(mem["used"].(float64)),
		DiskTotal:     int(disk["total"].(float64)),
		DiskUsed:      int(disk["used"].(float64)),
	}

	res.QEMU = s.qemuFactory.Create(res)
	res.LXC = s.lxcFactory.Create(res)
	return res, nil
}

func (s *NodeService) power(node string, command string) error {
	params := url.Values{"command": {command}}
	_, err := s.client.Post("nodes/"+node+"/status", params)
	return err
}

func (s *NodeService) Reboot(node string) error {
	return s.power(node, "reboot")
}

func (s *NodeService) Shutdown(node string) error {
	return s.power(node, "shutdown")
}
