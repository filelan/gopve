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
