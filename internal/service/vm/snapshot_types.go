package vm

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

type Snapshot struct {
	vm *VirtualMachine

	name        string
	description string
	timestamp   time.Time

	withRAM bool

	parent string
}

func NewSnapshot(vm *VirtualMachine, name string, description string, timestamp time.Time, withRAM bool, parent string) *Snapshot {
	return &Snapshot{
		vm: vm,

		name:        name,
		description: description,
		timestamp:   timestamp,

		withRAM: withRAM,

		parent: parent,
	}
}

func NewCurrentSnapshot(vm *VirtualMachine, parent string) *Snapshot {
	return &Snapshot{
		vm: vm,

		name: "current",

		parent: parent,
	}
}

func (obj *Snapshot) Name() string {
	return obj.name
}

func (obj *Snapshot) Description() string {
	if obj.name == "current" {
		return "You are here!"
	}

	return obj.description
}

func (obj *Snapshot) Timestamp() time.Time {
	if obj.name == "current" {
		return time.Now()
	}

	return obj.timestamp
}

func (obj *Snapshot) WithRAM() bool {
	if obj.name == "current" {
		return true
	}

	return obj.withRAM
}

func (obj *Snapshot) Parent() string {
	return obj.parent
}

func (obj *Snapshot) GetParent() (vm.Snapshot, error) {
	if obj.parent == "" {
		return nil, vm.ErrRootParentSnapshot
	} else {
		return obj.vm.GetSnapshot(obj.parent)
	}
}

func (obj *Snapshot) GetProperties() (vm.SnapshotProperties, error) {
	if obj.name == "current" {
		return vm.SnapshotProperties{
			Description: "You are here!",
		}, nil
	}

	return vm.SnapshotProperties{
		Description: obj.description,
	}, nil
}

func (obj *Snapshot) SetProperties(props vm.SnapshotProperties) error {
	if obj.name == "current" {
		return vm.ErrUpdateCurrentSnapshot
	}

	if err := obj.vm.svc.client.Request(http.MethodPut, fmt.Sprintf("nodes/%s/%s/%d/snapshot/%s/config", obj.vm.node, obj.vm.kind.String(), obj.vm.vmid, obj.name), request.Values{
		"description": {props.Description},
	}, nil); err != nil {
		return err
	}

	obj.description = props.Description

	return nil
}

func (obj *Snapshot) Delete() error {
	if obj.name == "current" {
		return vm.ErrDeleteCurrentSnapshot
	}

	return obj.vm.svc.client.Request(http.MethodDelete, fmt.Sprintf("nodes/%s/%s/%d/snapshot/%s", obj.vm.node, obj.vm.kind.String(), obj.vm.vmid, obj.name), nil, nil)
}
