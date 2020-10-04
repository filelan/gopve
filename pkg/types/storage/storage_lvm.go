package storage

import (
	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageLVM interface {
	Storage

	// Base volume. This volume is automatically activated.
	BaseStorage() string

	// Volume group name.
	VolumeGroup() string

	// Zero-out data when removing LVs.
	SafeRemove() bool

	// Limit the thoughput of the data stream, in bytes per second. If the value is positive, tries to keep the overall rate at the specified value for the whole session. If the value is negative, it is an upper limit for each read/write system call pair. In other words, the negative number will never exceed that limit, the positive number will exceed it to make good for previous underutilization
	SafeRemoveThroughput() int

	// Only use logical volumes tagged with 'pve-vm-ID'.
	TaggedOnly() bool
}

type StorageLVMProperties struct {
	BaseStorage          string
	VolumeGroup          string
	SafeRemove           bool
	SafeRemoveThroughput int
	TaggedOnly           bool
}

func NewStorageLVMProperties(
	props types.Properties,
) (*StorageLVMProperties, error) {
	obj := new(StorageLVMProperties)

	if v, ok := props["base"].(string); ok {
		obj.BaseStorage = v
	} else {
		obj.BaseStorage = DefaultStorageLVMBaseStorage
	}

	if v, ok := props["vgname"].(string); ok {
		obj.VolumeGroup = v
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", "vgname")
		return nil, err
	}

	if v, ok := props["saferemove"].(float64); ok {
		obj.SafeRemove = internal_types.NewPVEBoolFromFloat64(v).Bool()
	} else {
		obj.SafeRemove = DefaultStorageLVMSafeRemove
	}

	if v, ok := props["saferemove_throughput"].(float64); ok {
		if v == float64(int(v)) {
			obj.SafeRemoveThroughput = int(v)
		} else {
			err := errors.ErrInvalidProperty
			err.AddKey("name", "saferemove_throughput")
			err.AddKey("value", v)
			return nil, err
		}
	} else {
		obj.SafeRemoveThroughput = DefaultStorageLVMSafeRemoveThroughput
	}

	if v, ok := props["tagged_only"].(float64); ok {
		obj.TaggedOnly = internal_types.NewPVEBoolFromFloat64(v).Bool()
	} else {
		obj.SafeRemove = DefaultStorageLVMTaggedOnly
	}

	return obj, nil
}

const (
	StorageLVMContent       = ContentQEMUData & ContentContainerData
	StorageLVMImageFormat   = ImageFormatRaw
	StorageLVMAllowShare    = AllowSharePossible
	StorageLVMAllowSnapshot = AllowSnapshotNever
	StorageLVMAllowClone    = AllowCloneNever
)

const (
	DefaultStorageLVMBaseStorage          = ""
	DefaultStorageLVMSafeRemove           = false
	DefaultStorageLVMSafeRemoveThroughput = -10485760
	DefaultStorageLVMTaggedOnly           = false
)
