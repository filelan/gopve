package service

type Node struct {
	provider NodeServiceProvider
	QEMU     QEMUServiceProvider
	LXC      LXCServiceProvider

	Node          string
	Status        string
	Uptime        int
	CPUTotal      int
	CPUPercentage float64
	MemTotal      int
	MemUsed       int
	DiskTotal     int
	DiskUsed      int
}

type NodeList []*Node

func (e *Node) Reboot() error {
	return e.provider.Reboot(e.Node)
}

func (e *Node) Shutdown() error {
	return e.provider.Shutdown(e.Node)
}
