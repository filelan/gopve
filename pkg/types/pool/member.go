package pool

import (
	"encoding/json"
	"fmt"

	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

type Kind int

const (
	KindVirtualMachine Kind = iota
	KindStorage
)

func (obj *Kind) Unmarshal(s string) error {
	switch s {
	case "qemu", "lxc":
		*obj = KindVirtualMachine

	case "storage":
		*obj = KindStorage

	default:
		return fmt.Errorf("unknown pool member kind %s", s)
	}

	return nil
}

func (obj *Kind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}

type PoolMember interface {
	ID() string
}

type PoolMemberVirtualMachine interface {
	PoolMember

	Get() (vm.VirtualMachine, error)
}

type PoolMemberStorage interface {
	PoolMember

	Get() (storage.Storage, error)
}
