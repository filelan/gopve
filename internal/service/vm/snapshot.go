package vm

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types/task"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

type getSnapshotResponseJSON struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	SnapTime    uint64 `json:"snaptime"`

	VMState types.PVEBool `json:"vmstate"`

	Parent string `json:"parent"`
}

func (obj getSnapshotResponseJSON) Map(virtualMachine *VirtualMachine) (vm.Snapshot, error) {
	if obj.Name == "current" {
		return NewCurrentSnapshot(virtualMachine, obj.Parent), nil
	}

	loc, err := time.LoadLocation("UTC")
	if err != nil {
		return nil, err
	}
	timestamp := time.Unix(int64(obj.SnapTime), 0).In(loc)

	return NewSnapshot(virtualMachine, obj.Name, obj.Description, timestamp, obj.VMState.Bool(), obj.Parent), nil
}

func (obj *VirtualMachine) ListSnapshots() ([]vm.Snapshot, error) {
	var res []getSnapshotResponseJSON
	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/%s/%d/snapshot", obj.node, obj.kind.String(), obj.vmid), nil, &res); err != nil {
		return nil, err
	}

	snapshots := make([]vm.Snapshot, len(res))
	for i, snapshot := range res {
		s, err := snapshot.Map(obj)
		if err != nil {
			return nil, err
		}

		snapshots[i] = s
	}

	return snapshots, nil
}

func (obj *VirtualMachine) GetSnapshot(name string) (vm.Snapshot, error) {
	snapshots, err := obj.ListSnapshots()
	if err != nil {
		return nil, err
	}

	for _, snapshot := range snapshots {
		if snapshot.Name() == name {
			return snapshot, nil
		}
	}

	return nil, vm.ErrNoSnapshot
}

func (obj *VirtualMachine) CreateSnapshot(name string, props vm.SnapshotProperties) (task.Task, error) {
	form, err := props.MapToValues()
	if err != nil {
		return nil, err
	}

	form.AddString("snapname", name)

	var task string
	if err := obj.svc.client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/%s/%d/snapshot", obj.node, obj.kind.String(), obj.vmid), form, &task); err != nil {
		return nil, err
	}

	return obj.svc.api.Task().Get(task)
}

func (obj *VirtualMachine) RollbackToSnapshot(name string) (task.Task, error) {
	var task string
	if err := obj.svc.client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/%s/%d/snapshot/%s/rollback", obj.node, obj.kind.String(), obj.vmid, name), nil, &task); err != nil {
		return nil, err
	}

	return obj.svc.api.Task().Get(task)
}
