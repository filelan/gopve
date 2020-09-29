package storage

type StorageRBD interface {
	Storage

	MonitorHosts() []string
	Username() string

	// Always access rbd through krbd kernel module
	UseKRBD() bool

	PoolName() string
}

const (
	StorageRBDContents    = ContentQEMUData & ContentContainerData
	StorageRBDImageFormat = ImageFormatRaw
	StorageRBDShared      = AllowShareForced
	StorageRBDSnapshots   = AllowSnapshotAll
	StorageRBDClones      = AllowCloneAll
)

var DefaultStorageRBDMonitorHosts = []string{}

const (
	DefaultStorageRBDUsername = ""
	DefaultStorageRBDUseKRBD  = false
	DefaultStorageRBDPoolName = "rbd"
)
