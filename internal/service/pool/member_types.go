package pool

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xabinapal/gopve/pkg/types/pool"
	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

type PoolMember struct {
	svc *Service

	id string
}

func (obj *PoolMember) ID() string {
	return obj.id
}

type PoolMemberVirtualMachine struct {
	PoolMember
}

func NewPoolMemberVirtualMachine(svc *Service, id string) pool.PoolMember {
	return &PoolMemberVirtualMachine{
		PoolMember: PoolMember{
			svc: svc,
			id:  id,
		},
	}
}

func (obj *PoolMemberVirtualMachine) Get() (vm.VirtualMachine, error) {
	id := strings.Split(obj.id, "/")
	if len(id) != 2 {
		return nil, fmt.Errorf("unknown virtual machine pool member id")
	}

	vmid, err := strconv.Atoi(id[1])
	if err != nil {
		return nil, err
	}

	return obj.svc.api.VirtualMachine().Get(uint(vmid))
}

type PoolMemberStorage struct {
	PoolMember
}

func NewPoolMemberStorage(svc *Service, id string) pool.PoolMember {
	return &PoolMemberStorage{
		PoolMember: PoolMember{
			svc: svc,
			id:  id,
		},
	}
}

func (obj *PoolMemberStorage) ID() string {
	return obj.id
}

func (obj *PoolMemberStorage) Get() (storage.Storage, error) {
	id := strings.Split(obj.id, "/")
	if len(id) < 2 {
		return nil, fmt.Errorf("unknown storage pool member id")
	}

	return obj.svc.api.Storage().Get(id[len(id)-1])
}
