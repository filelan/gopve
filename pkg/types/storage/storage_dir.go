package storage

import (
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

const (
	StorageDirContent       = ContentQEMUData & ContentContainerData & ContentISO & ContentContainerTemplate & ContentBackup & ContentSnippet
	StorageDirImageFormat   = ImageFormatRaw & ImageFormatQcow2 & ImageFormatVMDK
	StorageDirAllowShare    = AllowShareNever
	StorageDirAllowSnapshot = AllowSnapshotQcow2
	StorageDirAllowClone    = AllowCloneQcow2
)

type StorageDirProperties struct {
	LocalPath          string
	LocalPathCreate    bool
	LocalPathIsManaged bool
}

const (
	mkDirLocalPath          = "path"
	mkDirLocalPathCreate    = "mkdir"
	mkDirLocalPathIsManaged = "is_mountpoint"
)

const (
	DefaultStorageDirLocalPathCreate = true
	DefaultStorageDirLocalIsManaged  = false
)

func NewStorageDirProperties(
	props types.Properties,
) (*StorageDirProperties, error) {
	obj := new(StorageDirProperties)

	return obj, errors.ChainUntilFail(
		func() error {
			return props.SetRequiredString(
				mkDirLocalPath,
				&obj.LocalPath,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkDirLocalPathCreate,
				&obj.LocalPathCreate,
				DefaultStorageDirLocalPathCreate,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkDirLocalPathIsManaged,
				&obj.LocalPathIsManaged,
				DefaultStorageDirLocalIsManaged,
				nil,
			)
		},
	)

	return obj, nil
}
