package vm

import (
	"fmt"
)

type Kind string

const (
	KindQEMU Kind = "qemu"
	KindLXC  Kind = "lxc"
)

func (obj Kind) IsValid() error {
	switch obj {
	case KindQEMU, KindLXC:
		return nil
	default:
		return fmt.Errorf("invalid virtual machine kind")
	}
}

func (obj Kind) String() string {
	return string(obj)
}
