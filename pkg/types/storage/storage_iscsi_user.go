package storage

import (
	"github.com/xabinapal/gopve/pkg/types"
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

	if err := props.SetRequiredString(mkISCSIUserPortal, &obj.Portal, nil); err != nil {
		return nil, err
	}

	if err := props.SetRequiredString(mkISCSIUserTarget, &obj.Target, nil); err != nil {
		return nil, err
	}

	return obj, nil
}
