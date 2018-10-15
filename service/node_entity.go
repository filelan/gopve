package service

type Node struct {
	provider NodeServiceProvider
	QEMU     QEMUServiceProvider
	LXC      LXCServiceProvider
	Task     TaskServiceProvider

	Node          string  `n:"node"`
	Status        string  `n:"status"`
	Uptime        int     `n:"uptime"`
	CPUTotal      int     `n:"maxcpu,cpuinfo.cpus"`
	CPUPercentage float64 `n:"cpu"`
	MemTotal      int     `n:"maxmem,memory.total"`
	MemUsed       int     `n:"mem,memory.used"`
	DiskTotal     int     `n:"maxdisk,rootfs.total"`
	DiskUsed      int     `n:"disk,rootfs.used"`
}

type NodeList []*Node

func (e *Node) Reboot() error {
	return e.provider.Reboot(e.Node)
}

func (e *Node) Shutdown() error {
	return e.provider.Shutdown(e.Node)
}
