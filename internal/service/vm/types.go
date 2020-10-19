package vm

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
	"github.com/xabinapal/gopve/pkg/types/vm/lxc"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
)

type VirtualMachine struct {
	svc *Service

	vmid     uint
	kind     vm.Kind
	node     string
	name     string
	template bool

	props *vm.Properties
}

func NewVirtualMachine(
	svc *Service,
	vmid uint,
	kind vm.Kind,
	node string,
	name string,
	template bool,
	props *vm.Properties,
) *VirtualMachine {
	return &VirtualMachine{
		svc:      svc,
		vmid:     vmid,
		kind:     kind,
		node:     node,
		name:     name,
		template: template,

		props: props,
	}
}

func NewDynamicVirtualMachine(
	svc *Service,
	vmid uint,
	kind vm.Kind,
	node string,
	name string,
	template bool,
	props *vm.Properties,
	extraProps types.Properties,
) (vm.VirtualMachine, error) {
	obj := NewVirtualMachine(svc, vmid, kind, node, name, template, props)
	switch kind {
	case vm.KindQEMU:
		return NewQEMU(*obj, extraProps)
	case vm.KindLXC:
		return NewLXC(*obj, extraProps)
	default:
		return nil, vm.ErrInvalidKind
	}
}

func (this *VirtualMachine) Load() error {
	if this.props != nil {
		return nil
	}

	obj, err := this.svc.Get(this.vmid)
	if err != nil {
		return err
	}

	props, err := obj.GetProperties()
	if err != nil {
		panic("this should never happen")
	}

	this.props = &props

	return nil
}

func (obj *VirtualMachine) VMID() uint {
	return obj.vmid
}

func (obj *VirtualMachine) Kind() vm.Kind {
	return obj.kind
}

func (obj *VirtualMachine) Node() string {
	return obj.node
}

func (obj *VirtualMachine) Name() string {
	return obj.name
}

func (obj *VirtualMachine) Template() bool {
	return obj.template
}

func (this *VirtualMachine) GetProperties() (vm.Properties, error) {
	if err := this.Load(); err != nil {
		return vm.Properties{}, err
	}

	return *this.props, nil
}

func (this *VirtualMachine) Description() (string, error) {
	props, err := this.GetProperties()
	if err != nil {
		return "", err
	}

	return props.Description, nil
}

func (this *VirtualMachine) Digest() (string, error) {
	props, err := this.GetProperties()
	if err != nil {
		return "", err
	}

	return props.Digest, nil
}

func (obj *VirtualMachine) GetStatus() (vm.Status, error) {
	var res struct {
		Status vm.Status `json:"status"`
	}

	if err := obj.svc.client.Request(http.MethodPost, fmt.Sprintf("node/%s/qemu/%d/status/current", obj.node, obj.vmid), nil, &res); err != nil {
		return vm.StatusStopped, err
	}

	if err := res.Status.IsValid(); err != nil {
		return vm.StatusStopped, err
	}

	return res.Status, nil
}

func (obj *VirtualMachine) ConvertToTemplate() error {
	if err := obj.svc.client.Request(
		http.MethodPost,
		fmt.Sprintf(
			"node/%s/%s/%d/template",
			obj.node,
			obj.kind.String(),
			obj.vmid,
		),
		nil,
		nil,
	); err != nil {
		return err
	}

	obj.template = true
	return nil
}

type QEMUVirtualMachine struct {
	VirtualMachine
	props *qemu.Properties
}

func NewQEMU(
	obj VirtualMachine,
	extraProps types.Properties,
) (*QEMUVirtualMachine, error) {
	qemuObj := &QEMUVirtualMachine{
		VirtualMachine: obj,
	}

	if extraProps != nil {
		props, err := qemu.NewProperties(extraProps)
		if err != nil {
			return nil, err
		}

		qemuObj.props = &props
	}

	return qemuObj, nil
}

func (this *QEMUVirtualMachine) Load() error {
	if this.props != nil {
		return nil
	}

	obj, err := this.svc.Get(this.vmid)
	if err != nil {
		return err
	}

	qemuObj, ok := obj.(*QEMUVirtualMachine)
	if !ok {
		return fmt.Errorf("invalid kind")
	}

	props, err := qemuObj.GetProperties()
	if err != nil {
		panic("this should never happen")
	}

	qemuProps, err := qemuObj.GetQEMUProperties()
	if err != nil {
		panic("this should never happen")
	}

	this.VirtualMachine.props = &props
	this.props = &qemuProps

	return nil
}

func (obj *QEMUVirtualMachine) GetQEMUProperties() (qemu.Properties, error) {
	if err := obj.Load(); err != nil {
		return qemu.Properties{}, err
	}

	return *obj.props, nil
}

func (obj *QEMUVirtualMachine) SetQEMUProperties(
	props qemu.Properties,
) error {
	form, err := props.MapToValues()
	if err != nil {
		return err
	}

	if err := obj.svc.client.Request(http.MethodPost, fmt.Sprintf("node/%s/qemu/%d/config", obj.node, obj.vmid), form, nil); err != nil {
		return err
	}

	obj.props = &props

	return nil
}

func (obj *QEMUVirtualMachine) CPU() (qemu.CPUProperties, error) {
	if err := obj.Load(); err != nil {
		return qemu.CPUProperties{}, err
	}

	return obj.props.CPU, nil
}

func (obj *QEMUVirtualMachine) Memory() (qemu.MemoryProperties, error) {
	if err := obj.Load(); err != nil {
		return qemu.MemoryProperties{}, err
	}

	return obj.props.Memory, nil
}

type LXCVirtualMachine struct {
	VirtualMachine
	props *lxc.Properties
}

func NewLXC(
	obj VirtualMachine,
	extraProps types.Properties,
) (*LXCVirtualMachine, error) {
	lxcObj := &LXCVirtualMachine{
		VirtualMachine: obj,
	}

	if extraProps != nil {
		props, err := lxc.NewProperties(extraProps)
		if err != nil {
			return nil, err
		}

		lxcObj.props = &props
	}

	return lxcObj, nil
}

func (this *LXCVirtualMachine) Load() error {
	if this.props != nil {
		return nil
	}

	obj, err := this.svc.Get(this.vmid)
	if err != nil {
		return err
	}

	lxcObj, ok := obj.(*LXCVirtualMachine)
	if !ok {
		return fmt.Errorf("invalid kind")
	}

	props, err := lxcObj.GetProperties()
	if err != nil {
		panic("this should never happen")
	}

	lxcProps, err := lxcObj.GetLXCProperties()
	if err != nil {
		panic("this should never happen")
	}

	this.VirtualMachine.props = &props
	this.props = &lxcProps

	return nil
}

func (obj *LXCVirtualMachine) GetLXCProperties() (lxc.Properties, error) {
	if err := obj.Load(); err != nil {
		return lxc.Properties{}, err
	}

	return *obj.props, nil
}

func (obj *LXCVirtualMachine) SetLXCProperties(props lxc.Properties) error {
	form, err := props.MapToValues()
	if err != nil {
		return err
	}

	if err := obj.svc.client.Request(http.MethodPost, fmt.Sprintf("node/%s/lxcqemu/%d/config", obj.node, obj.vmid), form, nil); err != nil {
		return err
	}

	obj.props = &props

	return nil
}

func (obj *LXCVirtualMachine) CPU() (lxc.CPUProperties, error) {
	if err := obj.Load(); err != nil {
		return lxc.CPUProperties{}, err
	}

	return obj.props.CPU, nil
}

func (obj *LXCVirtualMachine) Memory() (lxc.MemoryProperties, error) {
	if err := obj.Load(); err != nil {
		return lxc.MemoryProperties{}, err
	}

	return obj.props.Memory, nil
}

func (virtualMachine *VirtualMachine) getHighAvailabilitySID() string {
	switch virtualMachine.kind {
	case vm.KindQEMU:
		return fmt.Sprintf("vm:%d", virtualMachine.vmid)
	case vm.KindLXC:
		return fmt.Sprintf("ct:%d", virtualMachine.vmid)
	default:
		return ""
	}
}
