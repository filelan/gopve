package task

import (
	"strconv"
	"strings"

	"github.com/xabinapal/gopve/pkg/types/task"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

type Task struct {
	svc *Service

	upid string

	node   string
	action task.Action
	id     string
	user   string
}

func NewTask(svc *Service, node, uuid, action, id, user, extra string) *Task {
	upid := strings.Join([]string{"UPID", node, uuid, action, id, user, extra}, ":")

	t := &Task{
		svc:  svc,
		upid: upid,
		node: node,
		id:   id,
		user: user,
	}

	if err := (&t.action).Unmarshal(action); err != nil {
		t.action = task.ActionUnknown
	}

	return t
}

func (t *Task) UPID() string {
	return t.upid
}

func (t *Task) Node() string {
	return t.node
}

func (t *Task) Action() task.Action {
	return t.action
}

func (t *Task) ID() string {
	return t.id
}

func (t *Task) User() string {
	return t.user
}

type VirtualMachineTask struct {
	Task

	vmid uint
	kind vm.Kind
}

func NewVirtualMachineTask(t *Task) (*VirtualMachineTask, error) {
	vmid, err := strconv.Atoi(t.id)
	if err != nil {
		return nil, task.ErrInvalidID
	}

	var kind vm.Kind

	switch t.action {
	case task.ActionQMCreate:
		kind = vm.KindQEMU

	case task.ActionVZCreate:
		kind = vm.KindLXC

	default:
		return nil, task.ErrInvalidKind
	}

	return &VirtualMachineTask{
		Task: *t,
		vmid: uint(vmid),
		kind: kind,
	}, nil
}

func (obj *VirtualMachineTask) VMID() uint {
	return obj.vmid
}

func (obj *VirtualMachineTask) IsQEMU() bool {
	return obj.kind == vm.KindQEMU
}

func (obj *VirtualMachineTask) IsLXC() bool {
	return obj.kind == vm.KindLXC
}
