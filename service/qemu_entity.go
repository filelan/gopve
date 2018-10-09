package service

type QEMU struct {
	provider QEMUServiceProvider

	VMID   int
	Name   string
	Status string
	QEMUConfig
}

type QEMUConfig struct {
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

const (
	QEMUDefaultCPULimit = 0
	QEMUDefaultCPUUnits = 1000
)

func (e QEMU) Start() error {
	return e.provider.Start(e.VMID)
}

func (e *QEMU) Stop() error {
	return e.provider.Stop(e.VMID)
}

func (e *QEMU) Reset() error {
	return e.provider.Reset(e.VMID)
}

func (e *QEMU) Shutdown() error {
	return e.provider.Shutdown(e.VMID)
}

func (e *QEMU) Suspend() error {
	return e.provider.Suspend(e.VMID)
}

func (e *QEMU) Resume() error {
	return e.provider.Resume(e.VMID)
}
