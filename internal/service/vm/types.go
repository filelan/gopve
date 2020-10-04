package vm

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/vm"
)

type VirtualMachine struct {
	svc *Service

	vmid uint
	kind vm.Kind
	node string

	isTemplate bool

	props *vm.Properties
}

func NewVirtualMachine(
	svc *Service,
	vmid uint,
	kind vm.Kind,
	node string,
	isTemplate bool,
	props *vm.Properties,
) *VirtualMachine {
	return &VirtualMachine{
		svc:        svc,
		vmid:       vmid,
		kind:       kind,
		node:       node,
		isTemplate: isTemplate,
		props:      props,
	}
}

func NewDynamicVirtualMachine(
	svc *Service,
	vmid uint,
	kind vm.Kind,
	node string,
	isTemplate bool,
	props *vm.Properties,
	extraProps types.Properties,
) (vm.VirtualMachine, error) {
	obj := NewVirtualMachine(svc, vmid, kind, node, isTemplate, props)
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

func (obj *VirtualMachine) IsTemplate() bool {
	return obj.isTemplate
}

func (this *VirtualMachine) GetProperties() (vm.Properties, error) {
	if err := this.Load(); err != nil {
		return vm.Properties{}, err
	}

	return *this.props, nil
}

func (this *VirtualMachine) Name() (string, error) {
	props, err := this.GetProperties()
	if err != nil {
		return "", err
	}

	return props.Name, nil
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

	obj.isTemplate = true
	return nil
}

type QEMUVirtualMachine struct {
	VirtualMachine
	props *vm.QEMUProperties
}

func NewQEMU(
	obj VirtualMachine,
	extraProps types.Properties,
) (*QEMUVirtualMachine, error) {
	qemuObj := &QEMUVirtualMachine{
		VirtualMachine: obj,
	}

	if extraProps != nil {
		props, err := vm.NewQEMUProperties(extraProps)
		if err != nil {
			return nil, err
		}

		qemuObj.props = props
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

func (obj *QEMUVirtualMachine) GetQEMUProperties() (vm.QEMUProperties, error) {
	if err := obj.Load(); err != nil {
		return vm.QEMUProperties{}, err
	}

	return *obj.props, nil
}

func (obj *QEMUVirtualMachine) SetQEMUProperties(
	props vm.QEMUProperties,
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

func (obj *QEMUVirtualMachine) CPU() (vm.QEMUCPUProperties, error) {
	if err := obj.Load(); err != nil {
		return vm.QEMUCPUProperties{}, err
	}

	return obj.props.CPU, nil
}

func (obj *QEMUVirtualMachine) Memory() (vm.QEMUMemoryProperties, error) {
	if err := obj.Load(); err != nil {
		return vm.QEMUMemoryProperties{}, err
	}

	return obj.props.Memory, nil
}

type LXCVirtualMachine struct {
	VirtualMachine
	props *vm.LXCProperties

	cpu    vm.LXCCPUProperties
	memory vm.LXCMemoryProperties
}

func NewLXC(
	obj VirtualMachine,
	extraProps types.Properties,
) (*LXCVirtualMachine, error) {
	lxcObj := &LXCVirtualMachine{
		VirtualMachine: obj,
	}

	if extraProps != nil {
		props, err := vm.NewLXCProperties(extraProps)
		if err != nil {
			return nil, err
		}

		lxcObj.props = props
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

func (obj *LXCVirtualMachine) GetLXCProperties() (vm.LXCProperties, error) {
	if err := obj.Load(); err != nil {
		return vm.LXCProperties{}, err
	}

	return *obj.props, nil
}

func (obj *LXCVirtualMachine) SetLXCProperties(props vm.LXCProperties) error {
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

func (obj *LXCVirtualMachine) CPU() (vm.LXCCPUProperties, error) {
	if err := obj.Load(); err != nil {
		return vm.LXCCPUProperties{}, err
	}

	return obj.cpu, nil
}

func (obj *LXCVirtualMachine) Memory() (vm.LXCMemoryProperties, error) {
	if err := obj.Load(); err != nil {
		return vm.LXCMemoryProperties{}, err
	}

	return obj.memory, nil
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
