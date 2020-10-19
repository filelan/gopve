package storage

import (
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
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

	return obj, errors.ChainUntilFail(
		func() error {
			return props.SetRequiredString(
				mkISCSIKernelPortal,
				&obj.Portal,
				nil)
		},
		func() error {
			return props.SetRequiredString(
				mkISCSIKernelTarget,
				&obj.Target,
				nil,
			)
		},
	)
}
