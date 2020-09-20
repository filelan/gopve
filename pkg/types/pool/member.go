package pool

import (
	"encoding/json"
	"fmt"

	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

type MemberKind int

const (
	MemberKindVirtualMachine MemberKind = iota
	MemberKindStorage
)

func (obj *MemberKind) Unmarshal(s string) error {
	switch s {
	case "qemu", "lxc":
		*obj = MemberKindVirtualMachine

	case "storage":
		*obj = MemberKindStorage

	default:
		return fmt.Errorf("unknown pool member kind %s", s)
	}

	return nil
}

func (obj *MemberKind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}

type PoolMember interface {
	ID() string
	Kind() MemberKind
}

type PoolMemberVirtualMachine interface {
	PoolMember

	Get() (vm.VirtualMachine, error)
}

type PoolMemberStorage interface {
	PoolMember

	Get() (storage.Storage, error)
}
