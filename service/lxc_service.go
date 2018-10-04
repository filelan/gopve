package service

import (
	"github.com/xabinapal/gopve/internal"
)

type LXCServiceProvider interface {
	Create() error
	Update() error
	Delete() error
	Clone() error
}

type LXCService struct {
	client *internal.Client
}

type LXC struct {
	CTID        int
	Name        string
	CPUCores    int
	CPULimit    int
	CPUUnits    int
	MemoryTotal int
	MemorySwap  int
}

func (lxc *LXCService) Create() error {
}

func (lxc *LXCService) Update() error {
}

func (lxc *LXCService) Delete() error {
}

func (lxc *LXCService) Clone() error {
}
