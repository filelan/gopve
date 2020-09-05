package vm

type QEMUVirtualMachine interface {
	VirtualMachine
}

type QEMUCreateOptions struct {
	VMID uint
	Node string

	Name        string
	Description string

	CPU    QEMUCPUProperties
	Memory QEMUMemoryProperties
}

type QEMUCPUProperties struct {
	Type    string
	Sockets uint
	Cores   uint
	VCPUs   uint

	Limit uint
	Units uint

	NUMA bool

	FreezeAtStartup bool
}

type QEMUMemoryProperties struct {
	Memory uint

	Ballooning    bool
	MinimumMemory uint
	Shares        uint
}
