package service

type Node struct {
	provider NodeServiceProvider

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

type NodeList []Node

func (n *Node) Reboot() error {
	return n.provider.Reboot(n.Node)
}

func (n *Node) Shutdown() error {
	return n.provider.Shutdown(n.Node)
}
