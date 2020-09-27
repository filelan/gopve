package storage

import (
	"encoding/json"
	"fmt"
)

type StorageZFSOverISCSI interface {
	Storage

	Portal() string
	Target() string

	PoolName() string

	BlockSize() uint
	UseSparse() bool
	WriteCache() bool

	ISCSIProvider() ISCSIProvider

	ComstarHostGroup() string
	ComstarTargetGroup() string
	LIOTargetPortalGroup() string
}

const (
	StorageZFSOverISCSIKernelContents    = ContentQEMUData
	StorageZFSOverISCSIKernelImageFormat = ImageFormatRaw
	StorageZFSOverISCSIKernelShared      = AllowShareForced
	StorageZFSOverISCSIKernelSnapshots   = AllowSnapshotNever
	StorageZFSOverISCSIKernelClones      = AllowCloneNever
)

const (
	DefaultStorageZFSOverISCSIUseSparse = false
	DefaultStorageZFSOverISCSIWriteCache  = true
	DefaultStorageZFSOverISCSIComstarHostGroup = ""
	DefaultStorageZFSOverISCSIComstarTargetGroup = ""
	DefaultStorageZFSOverISCSILIOTargetPortalGroup = ""
)


type ISCSIProvider uint

const (
	ISCSIProviderComstar ISCSIProvider = iota
	ISCSIProviderISTGT
	ISCSIProviderIET
	ISCSIProviderLIO
)

func (obj ISCSIProvider) Marshal() (string, error) {
	switch obj {
	case ISCSIProviderComstar:
		return "comstar", nil
	case ISCSIProviderISTGT:
		return "istgt", nil
	case ISCSIProviderIET:
		return "iet", nil
	case ISCSIProviderLIO:
		return "LIO", nil
	default:
		return "", fmt.Errorf("unknown iscsi provider")
	}
}

func (obj *ISCSIProvider) Unmarshal(s string) error {
	switch s {
	case "comstar":
		*obj = ISCSIProviderComstar
	case "istgt":
		*obj = ISCSIProviderISTGT
	case "iet":
		*obj = ISCSIProviderIET
	case "LIO":
		*obj = ISCSIProviderLIO
	default:
		return fmt.Errorf("can't unmarshal iscsi provider %s", s)
	}

	return nil
}

func (obj *ISCSIProvider) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
