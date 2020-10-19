package storage

import (
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageISCSIUser interface {
	Storage

	// iSCSI target
	Portal() string

	// iSCSI portal (IP or DNS name with optional port).
	Target() string
}

const (
	StorageISCSIUserContents    = ContentQEMUData
	StorageISCSIUserImageFormat = ImageFormatRaw
	StorageISCSIUserShared      = AllowShareForced
	StorageISCSIUserSnapshots   = AllowSnapshotNever
	StorageISCSIUserClones      = AllowCloneNever
)

type StorageISCSIUserProperties struct {
	Portal string
	Target string
}

const (
	mkISCSIUserPortal = "portal"
	mkISCSIUserTarget = "target"
)

func NewStorageISCSIUserProperties(
	props types.Properties,
) (*StorageISCSIUserProperties, error) {
	obj := new(StorageISCSIUserProperties)

	return obj, errors.ChainUntilFail(
		func() error {
			return props.SetRequiredString(
				mkISCSIUserPortal,
				&obj.Portal,
				nil,
			)
		},
		func() error {
			return props.SetRequiredString(
				mkISCSIUserTarget,
				&obj.Target,
				nil,
			)
		},
	)
}
