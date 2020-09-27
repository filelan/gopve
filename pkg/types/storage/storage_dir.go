package storage

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

const (
	DefaultStorageLocalPathCreate = true
	DefaultStorageLocalIsManaged  = false
)
