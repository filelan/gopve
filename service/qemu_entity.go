package service

type QEMU struct {
	provider QEMUServiceProvider

	VMID             int
	Name             string
	Status           string
	CPU              int
	CPUSockets       int
	CPUCores         int
	CPULimit         int
	CPUUnits         int
	MemoryTotal      int
	MemoryMinimum    int
	MemoryBallooning bool
}

type QEMUList []*QEMU
