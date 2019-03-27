package service

import "fmt"

type NodeError struct {
	Node string
}

func (e *NodeError) Error() string {
	return fmt.Sprintf("Node %s does not exist", e.Node)
}

type Node struct {
	provider NodeServiceProvider
	qemu     QEMUServiceProvider
	lxc      LXCServiceProvider
	task     TaskServiceProvider

	Node          string  `n:"node"`
	Status        string  `n:"status"`
	Uptime        int     `n:"uptime"`
	CPUTotal      int     `n:"maxcpu"`
	CPUPercentage float64 `n:"cpu"`
	MemTotal      int     `n:"maxmem"`
	MemUsed       int     `n:"mem"`
	DiskTotal     int     `n:"maxdisk"`
	DiskUsed      int     `n:"disk"`
}

type NodeList map[string]*Node

func (e *Node) QEMU() QEMUServiceProvider {
	return e.qemu
}

func (e *Node) LXC() LXCServiceProvider {
	return e.lxc
}

func (e *Node) Task() TaskServiceProvider {
	return e.task
}

func (e *Node) Reboot() error {
	return e.provider.Reboot(e.Node)
}

func (e *Node) Shutdown() error {
	return e.provider.Shutdown(e.Node)
}
