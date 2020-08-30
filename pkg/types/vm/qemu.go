package vm

import (
	"fmt"
)

type QEMUVirtualMachine interface {
	VirtualMachine
}

type QEMUImageFormat string

const (
	QEMUImageFormatRaw   QEMUImageFormat = "raw"
	QEMUImageFormatQcow2 QEMUImageFormat = "qcow2"
	QEMUImageFormatVMDK  QEMUImageFormat = "vmdk"
)

func (obj QEMUImageFormat) IsValid() error {
	switch obj {
	case QEMUImageFormatRaw, QEMUImageFormatQcow2, QEMUImageFormatVMDK:
		return nil
	default:
		return fmt.Errorf("invalid virtual machine status")
	}
}

type QEMUCreateOptions struct {
	VMID uint
	Node string

	Name        string
	Description string
}
