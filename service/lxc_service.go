package service

import (
	"errors"

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
	node   NodeServiceProvider
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

func NewLXCService(c *internal.Client, n NodeServiceProvider) *LXCService {
	lxc := &LXCService{
		client: c,
		node:   n,
	}

	return lxc
}

func (lxc *LXCService) Create() error {
	return errors.New("Not yet implemented")
}

func (lxc *LXCService) Update() error {
	return errors.New("Not yet implemented")
}

func (lxc *LXCService) Delete() error {
	return errors.New("Not yet implemented")
}

func (lxc *LXCService) Clone() error {
	return errors.New("Not yet implemented")
}
