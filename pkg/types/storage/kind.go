package storage

import (
	"encoding/json"
)

type Kind string

const (
	KindDir          Kind = "dir"
	KindLVM          Kind = "lvm"
	KindLVMThin      Kind = "lvmthin"
	KindZFS          Kind = "zfspool"
	KindNFS          Kind = "nfs"
	KindCIFS         Kind = "cifs"
	KindGlusterFS    Kind = "glusterfs"
	KindISCSIKernel  Kind = "iscsi"
	KindISCSIUser    Kind = "iscsidirect"
	KindCephFS       Kind = "cephfs"
	KindRBD          Kind = "rbd"
	KindDRBD         Kind = "drbd"
	KindZFSOverISCSI Kind = "zfs"
)

func (obj Kind) String() string {
	if v, err := obj.Marshal(); err == nil {
		return v
	}

	return ""
}

func (obj Kind) IsValid() bool {
	switch obj {
	case KindDir,
		KindLVM,
		KindLVMThin,
		KindZFS,
		KindNFS,
		KindCIFS,
		KindGlusterFS,
		KindISCSIKernel,
		KindISCSIUser,
		KindCephFS,
		KindRBD,
		KindDRBD,
		KindZFSOverISCSI:
		return true
	default:
		return false
	}
}

func (obj Kind) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj Kind) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *Kind) Unmarshal(s string) error {
	*obj = Kind(s)
	return nil
}

func (obj *Kind) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
