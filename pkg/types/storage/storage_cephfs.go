package storage

type StorageCephFS interface {
	Storage

	MonitorHosts() []string
	Username() string

	// Mount CephFS through FUSE.
	UseFUSE() bool

	ServerPath() string
	LocalPath() string
}

const (
	StorageCephFSContents    = ContentContainerTemplate & ContentISO & ContentBackup & ContentSnippet
	StorageCephFSImageFormat = ContentNone
	StorageCephFSShared      = AllowShareNever
	StorageCephFSSnapshots   = AllowSnapshotAll
	StorageCephFSClones      = AllowCloneNever
)

var DefaultStorageCephFSMonitorHosts = []string{}

const (
	DefaultStorageCephFSUsername   = ""
	DefaultStorageCephFSUseFUSE    = false
	DefaultStorageCephFSServerPath = "/"
)
