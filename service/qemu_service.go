package service

import (
	"errors"

	"github.com/xabinapal/gopve/internal"
)

type QEMUServiceProvider interface {
	Start() error
	Create() error
	Update() error
	Delete() error
	Clone() error
}

type QEMUService struct {
	client *internal.Client
	node   NodeServiceProvider
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

func NewQEMUService(c *internal.Client, n NodeServiceProvider) *QEMUService {
	qemu := &QEMUService{
		client: c,
		node:   n,
	}

	return qemu
}

func (qemu *QEMUService) Start() error {
	return errors.New("Not yet implemented")
}

func (qemu *QEMUService) Create() error {
	return errors.New("Not yet implemented")
}

func (qemu *QEMUService) Update() error {
	return errors.New("Not yet implemented")
}

func (qemu *QEMUService) Delete() error {
	return errors.New("Not yet implemented")
}

func (qemu *QEMUService) Clone() error {
	return errors.New("Not yet implemented")
}
