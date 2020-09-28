package storage

type StorageZFS interface {
	Storage

	PoolName() string

	BlockSize() uint
	UseSparse() bool

	LocalPath() string
}

const (
	StorageZFSContents    = ContentQEMUData & ContentContainerData
	StorageZFSImageFormat = ImageFormatRaw & ImageFormatSubVolume
	StorageZFSShared      = AllowShareNever
	StorageZFSSnapshots   = AllowSnapshotAll
	StorageZFSClones      = AllowCloneAll
)

const (
	DefaultStorageZFSBlockSize  uint = 8192
	DefaultStorageZFSUseSparse       = false
	DefaultStorageZFSMountPoint      = ""
)
