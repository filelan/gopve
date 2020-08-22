package types

import "fmt"

type VMService interface {
	GetAll() ([]VirtualMachine, error)
}

type VirtualMachineCategory string

const (
	QEMU VirtualMachineCategory = "qemu"
	LXC  VirtualMachineCategory = "lxc"
)

func (obj VirtualMachineCategory) IsValid() error {
	switch obj {
	case QEMU, LXC:
		return nil
	default:
		return fmt.Errorf("invalid virtual machine category")
	}
}

type VirtualMachineStatus string

const (
	VMRunning VirtualMachineStatus = "running"
	VMStopped VirtualMachineStatus = "stopped"
)

func (obj VirtualMachineStatus) IsValid() error {
	switch obj {
	case VMRunning, VMStopped:
		return nil
	default:
		return fmt.Errorf("invalid virtual machine status")
	}
}

type QEMUImageFormat string

const (
	ImageRaw   QEMUImageFormat = "raw"
	ImageQcow2 QEMUImageFormat = "qcow2"
	ImageVMDK  QEMUImageFormat = "vmdk"
)

func (obj QEMUImageFormat) IsValid() error {
	switch obj {
	case ImageRaw, ImageQcow2, ImageVMDK:
		return nil
	default:
		return fmt.Errorf("invalid virtual machine status")
	}
}

type VirtualMachine interface {
	Node() (Node, error)
	Category() VirtualMachineCategory

	VMID() uint
	Name() string
	Status() VirtualMachineStatus
	IsTemplate() bool

	Clone(options VMCloneOptions) (Task, error)

	Start() (Task, error)
	Stop() (Task, error)
	Reset() (Task, error)
	Shutdown() (Task, error)
	Reboot() (Task, error)
	Suspend() (Task, error)
	Resume() (Task, error)
}

type QEMUVirtualMachine interface {
	VirtualMachine
}

type LXCVirtualMachine interface {
	VirtualMachine
}

type QEMUCreateOptions struct {
	VMID uint
	Node string

	Name        string
	Description string
}

type VMCloneOptions struct {
	VMID         uint
	Name         string
	Description  string
	PoolName     string
	SnapshotName string

	BandwidthLimit    uint
	TemplateFullClone bool
	ImageFormat       QEMUImageFormat

	TargetNode    string
	TargetStorage string
}
