package storage

import (
	"github.com/xabinapal/gopve/pkg/types"
)

type StorageLVMThin interface {
	Storage

	VolumeGroup() string
	ThinPool() string
}

const (
	StorageLVMThinContents    = ContentQEMUData & ContentContainerData
	StorageLVMThinImageFormat = ImageFormatRaw
	StorageLVMThinShared      = AllowShareNever
	StorageLVMThinSnapshots   = AllowSnapshotAll
	StorageLVMThinClones      = AllowCloneAll
)

type StorageLVMThinProperties struct {
	VolumeGroup string
	ThinPool    string
}

const (
	mkLVMThinVolumeGroup = "vgname"
	mkLVMThinThinPool    = "thinpool"
)

func NewStorageLVMThinProperties(
	props types.Properties,
) (*StorageLVMThinProperties, error) {
	obj := new(StorageLVMThinProperties)

	if err := props.SetRequiredString(mkLVMThinVolumeGroup, &obj.VolumeGroup, nil); err != nil {
		return nil, err
	}

	if err := props.SetRequiredString(mkLVMThinThinPool, &obj.ThinPool, nil); err != nil {
		return nil, err
	}

	return obj, nil
}
