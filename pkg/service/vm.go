package service

import "github.com/xabinapal/gopve/pkg/types/vm"

type VirtualMachine interface {
	List() ([]vm.VirtualMachine, error)
}
