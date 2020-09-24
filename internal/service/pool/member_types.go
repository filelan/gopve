package pool

import (
	"fmt"
	"strconv"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types/pool"
	"github.com/xabinapal/gopve/pkg/types/storage"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

type PoolMember struct {
	svc *Service

	id   string
	kind pool.MemberKind
}

func (obj *PoolMember) ID() string {
	return obj.id
}

func (obj *PoolMember) Kind() pool.MemberKind {
	return obj.kind
}

type PoolMemberVirtualMachine struct {
	PoolMember

	vmid uint
}

func NewPoolMemberVirtualMachine(
	svc *Service,
	id string,
) (pool.PoolMember, error) {
	memberID := types.PVEStringKV{Separator: "/", AllowNoValue: false}
	if err := memberID.Unmarshal(id); err != nil {
		return nil, err
	}

	vmid, err := strconv.Atoi(memberID.Value())
	if err != nil {
		return nil, err
	}

	return &PoolMemberVirtualMachine{
		PoolMember: PoolMember{
			svc:  svc,
			id:   id,
			kind: pool.MemberKindVirtualMachine,
		},

		vmid: uint(vmid),
	}, nil
}

func (obj *PoolMemberVirtualMachine) MemberID() string {
	return strconv.Itoa(int(obj.vmid))
}

func (obj *PoolMemberVirtualMachine) Get() (vm.VirtualMachine, error) {
	return obj.svc.api.VirtualMachine().Get(obj.vmid)
}

type PoolMemberStorage struct {
	PoolMember

	name string
}

func NewPoolMemberStorage(svc *Service, id string) (pool.PoolMember, error) {
	memberID := types.PVEStringList{Separator: "/"}
	if err := memberID.Unmarshal(id); err != nil {
		return nil, err
	}

	if memberID.Len() < 2 {
		return nil, fmt.Errorf("unknown storage pool member id")
	}

	name := memberID.Elem(memberID.Len() - 1)

	return &PoolMemberStorage{
		PoolMember: PoolMember{
			svc:  svc,
			id:   fmt.Sprintf("storage/%s", name),
			kind: pool.MemberKindStorage,
		},

		name: name,
	}, nil
}

func (obj *PoolMemberStorage) MemberID() string {
	return obj.name
}

func (obj *PoolMemberStorage) Get() (storage.Storage, error) {
	return obj.svc.api.Storage().Get(obj.name)
}
