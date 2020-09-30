package storage

import "github.com/xabinapal/gopve/internal/types"

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

func (obj *StorageDirProperties) Unmarshal(props ExtraProperties) error {
	if v, ok := props["path"].(string); ok {
		obj.LocalPath = v
	} else {
		err := ErrMissingProperty
		err.AddKey("name", "path")
		return err
	}

	if v, ok := props["mkdir"].(int); ok {
		obj.LocalPathCreate = types.NewPVEBoolFromInt(v).Bool()
	} else {
		obj.LocalPathCreate = DefaultStorageDirLocalPathCreate
	}

	if v, ok := props["is_mountpoint"].(int); ok {
		obj.LocalPathIsManaged = types.NewPVEBoolFromInt(v).Bool()
	} else {
		obj.LocalPathIsManaged = DefaultStorageDirLocalIsManaged
	}

	return nil
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
