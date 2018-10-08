package service

type LXC struct {
	provider LXCServiceProvider

	CTID        int
	Name        string
	Status      string
	CPU         int
	CPULimit    int
	CPUUnits    int
	MemoryTotal int
	MemorySwap  int
}

type LXCList []*LXC

func (e *LXC) Start() error {
	return e.provider.Start(e.CTID)
}

func (e *LXC) Stop() error {
	return e.provider.Stop(e.CTID)
}

func (e *LXC) Reset() error {
	return e.provider.Reset(e.CTID)
}

func (e *LXC) Shutdown() error {
	return e.provider.Shutdown(e.CTID)
}

func (e *LXC) Suspend() error {
	return e.provider.Suspend(e.CTID)
}

func (e *LXC) Resume() error {
	return e.provider.Resume(e.CTID)
}
