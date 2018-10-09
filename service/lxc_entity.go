package service

type LXC struct {
	provider LXCServiceProvider

	VMID   int
	Name   string
	Status string
	LXCConfig
}

type LXCConfig struct {
	CPU         int
	CPULimit    int
	CPUUnits    int
	MemoryTotal int
	MemorySwap  int
}

type LXCList []*LXC

const (
	LXC_DEFAULT_CPU_LIMIT = 0
	LXC_DEFAULT_CPU_UNITS = 1000
)

func (e *LXC) Start() error {
	return e.provider.Start(e.VMID)
}

func (e *LXC) Stop() error {
	return e.provider.Stop(e.VMID)
}

func (e *LXC) Reset() error {
	return e.provider.Reset(e.VMID)
}

func (e *LXC) Shutdown() error {
	return e.provider.Shutdown(e.VMID)
}

func (e *LXC) Suspend() error {
	return e.provider.Suspend(e.VMID)
}

func (e *LXC) Resume() error {
	return e.provider.Resume(e.VMID)
}
