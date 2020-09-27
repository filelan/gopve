package storage

import (
	"encoding/json"
	"fmt"
)

type Kind int

const (
	KindUnknown Kind = iota
	KindDir
	KindLVM
	KindLVMThin
	KindZFS
	KindNFS
	KindCIFS
	KindGlusterFS
	KindISCSIKernelMode
	KindISCSIUserMode
	KindCephFS
	KindRBD
	KindZFSOverISCSI
)

func (obj Kind) Marshal() (string, error) {
	switch obj {
	case KindDir:
		return "dir", nil
	case KindLVM:
		return "lvm", nil
	case KindLVMThin:
		return "lvmthin", nil
	case KindZFS:
		return "zfspool", nil
	case KindNFS:
		return "nfs", nil
	case KindCIFS:
		return "cifs", nil
	case KindGlusterFS:
		return "glusterfs", nil
	case KindISCSIKernelMode:
		return "iscsi", nil
	case KindISCSIUserMode:
		return "iscsidirect", nil
	case KindCephFS:
		return "cephfs", nil
	case KindRBD:
		return "rbd", nil
	case KindZFSOverISCSI:
		return "zfs", nil

	default:
		return "", fmt.Errorf("unknown storage kind")
	}
}

func (obj *Kind) Unmarshal(s string) error {
	switch s {
	case "dir":
		*obj = KindDir
	case "lvm":
		*obj = KindLVM
	case "lvmthin":
		*obj = KindLVMThin
	case "zfspool":
		*obj = KindZFS
	case "nfs":
		*obj = KindNFS
	case "cifs":
		*obj = KindCIFS
	case "glusterfs":
		*obj = KindGlusterFS
	case "iscsi":
		*obj = KindISCSIKernelMode
	case "iscsidirect":
		*obj = KindISCSIUserMode
	case "cephfs":
		*obj = KindCephFS
	case "rbd":
		*obj = KindRBD
	case "zfs":
		*obj = KindZFSOverISCSI

	default:
		*obj = KindUnknown
		return fmt.Errorf("unknown storage kind %s", s)
	}

	return nil
}

func (obj *Kind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
