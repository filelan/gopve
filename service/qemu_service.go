package service

import (
	"github.com/xabinapal/gopve/internal"
)

type QEMUServiceProvider interface {
	Create() error
	Update() error
	Delete() error
	Clone() error
}

type QEMUService struct {
	client *internal.Client
}

type QEMU struct {
	VMID             int
	Name             string
	CPUSockets       int
	CPUCores         int
	CPULimit         int
	CPUUnits         int
	MemoryTotal      int
	MemoryMinimum    int
	MemoryBallooning bool
}

func (qemu *QEMUService) Create() error {
}

func (qemu *QEMUService) Update() error {
}

func (qemu *QEMUService) Delete() error {
}

func (qemu *QEMUService) Clone() error {
}
