package vm

type LXCVirtualMachine interface {
	VirtualMachine
}

type LXCCreateOptions struct {
	VMID uint
	Node string

	Name        string
	Description string

	OSTemplateStorage string
	OSTemplate        string

	CPU    LXCCPUProperties
	Memory LXCMemoryProperties

	RootFSStorage string
	RootFSSize    uint
}

type LXCCPUProperties struct {
	Cores uint

	Limit uint
	Units uint
}

type LXCMemoryProperties struct {
	Memory uint
	Swap   uint
}
