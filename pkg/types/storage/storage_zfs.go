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
	DefaultStorageZFSBlockSize = 8192
	DefaultStorageZFSUseSparse = false
	DefaultStorageZFSPoolName  = "rbd"
)