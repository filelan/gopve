package service

type QEMU struct {
	provider QEMUServiceProvider

	VMID   int
	Name   string
	Status string
	QEMUConfig
}

type QEMUConfig struct {
	OSType           string `n:"ostype"`
	CPU              int    `i:"always"`
	CPUSockets       int    `i:"always"`
	CPUCores         int    `n:"cores"`
	CPULimit         int    `n:"cpulimit"`
	CPUUnits         int    `n:"cpuunits"`
	MemoryTotal      int    `n:"memory"`
	MemoryMinimum    int    `n:"balloon"`
	MemoryBallooning bool   `i:"always"`
	IsNUMAAware      bool   `n:"numa"`
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
