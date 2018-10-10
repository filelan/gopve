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
	for _, node := range internal.NewJArray(data) {
		val := internal.NewJObject(node)
		row := &Node{
			provider: s,

			Node:          val.GetString("node"),
			Status:        val.GetString("status"),
			Uptime:        val.GetInt("uptime"),
			CPUTotal:      val.GetInt("maxcpu"),
			CPUPercentage: val.GetFloat("cpu"),
			MemTotal:      val.GetInt("maxmem"),
			MemUsed:       val.GetInt("mem"),
			DiskTotal:     val.GetInt("maxdisk"),
			DiskUsed:      val.GetInt("disk"),
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

	val := internal.NewJObject(data)
	cpu := val.GetJObject("cpuinfo")
	mem := val.GetJObject("memory")
	disk := val.GetJObject("rootfs")

	res := &Node{
		provider: s,

		Node:          node,
		Status:        "",
		Uptime:        val.GetInt("uptime"),
		CPUTotal:      cpu.GetInt("cpus"),
		CPUPercentage: val.GetFloat("cpu"),
		MemTotal:      mem.GetInt("total"),
		MemUsed:       mem.GetInt("used"),
		DiskTotal:     disk.GetInt("total"),
		DiskUsed:      disk.GetInt("used"),
	}

	res.QEMU = s.qemuFactory.Create(res)
	res.LXC = s.lxcFactory.Create(res)
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
