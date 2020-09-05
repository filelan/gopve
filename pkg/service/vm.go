package service

import (
	"github.com/xabinapal/gopve/pkg/types/vm"
)

//go:generate mockery --case snake --name VirtualMachine --filename vm.go

type VirtualMachine interface {
	List() ([]vm.VirtualMachine, error)
}
