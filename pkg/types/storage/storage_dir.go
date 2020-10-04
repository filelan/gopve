package storage

import (
	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageDir interface {
	Storage

	// Filesytem path.
	LocalPath() string

	// Create the directory if it doesn't exist.
	LocalPathCreate() bool

	// Assume the given path is an externally managed mountpoint and consider the storage offline if it is not mounted.
	LocalPathIsManaged() bool
}

type StorageDirProperties struct {
	LocalPath          string
	LocalPathCreate    bool
	LocalPathIsManaged bool
}

func NewStorageDirProperties(
	props types.Properties,
) (*StorageDirProperties, error) {
	obj := new(StorageDirProperties)

	if v, ok := props["path"].(string); ok {
		obj.LocalPath = v
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", "path")
		return nil, err
	}

	if v, ok := props["mkdir"].(float64); ok {
		obj.LocalPathCreate = internal_types.NewPVEBoolFromFloat64(v).Bool()
	} else {
		obj.LocalPathCreate = DefaultStorageDirLocalPathCreate
	}

	if v, ok := props["is_mountpoint"].(float64); ok {
		obj.LocalPathIsManaged = internal_types.NewPVEBoolFromFloat64(v).Bool()
	} else {
		obj.LocalPathIsManaged = DefaultStorageDirLocalIsManaged
	}

	return obj, nil
}

const (
	StorageDirContent       = ContentQEMUData & ContentContainerData & ContentISO & ContentContainerTemplate & ContentBackup & ContentSnippet
	StorageDirImageFormat   = ImageFormatRaw & ImageFormatQcow2 & ImageFormatVMDK
	StorageDirAllowShare    = AllowShareNever
	StorageDirAllowSnapshot = AllowSnapshotQcow2
	StorageDirAllowClone    = AllowCloneQcow2
)

const (
	DefaultStorageDirLocalPathCreate = true
	DefaultStorageDirLocalIsManaged  = false
)
