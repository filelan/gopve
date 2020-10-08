package storage

import (
	"github.com/xabinapal/gopve/pkg/types"
)

type StorageISCSIKernel interface {
	Storage

	Portal() string
	Target() string
}

const (
	StorageISCSIKernelContents    = ContentQEMUData
	StorageISCSIKernelImageFormat = ImageFormatRaw
	StorageISCSIKernelShared      = AllowShareForced
	StorageISCSIKernelSnapshots   = AllowSnapshotNever
	StorageISCSIKernelClones      = AllowCloneNever
)

type StorageISCSIKernelProperties struct {
	Portal string
	Target string
}

const (
	mkISCSIKernelPortal = "portal"
	mkISCSIKernelTarget = "target"
)

func NewStorageISCSIKernelProperties(
	props types.Properties,
) (*StorageISCSIKernelProperties, error) {
	obj := new(StorageISCSIKernelProperties)

	if err := props.SetRequiredString(mkISCSIKernelPortal, &obj.Portal, nil); err != nil {
		return nil, err
	}

	if err := props.SetRequiredString(mkISCSIKernelTarget, &obj.Target, nil); err != nil {
		return nil, err
	}

	return obj, nil
}
