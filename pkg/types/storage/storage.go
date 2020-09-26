package storage

type AllowShare int

const (
	AllowShareNever AllowShare = iota
	AllowSharePossible
	AllowShareForced
)

type AllowSnapshot int

const (
	AllowSnapshotNever AllowSnapshot = iota
	AllowSnapshotQcow2
	AllowSnapshotAll
)

type AllowClone int

const (
	AllowCloneNever AllowClone = iota
	AllowCloneQcow2
	AllowCloneAll
)

type Storage interface {
	Name() string
	Kind() (Kind, error)
	Shared() (bool, error)
	Content() (Content, error)

	Nodes() ([]string, error)

	ImageFormat() (ImageFormat, error)
	MaxBackupsPerVM() (uint, error)
}

type StorageDir interface {
	Storage

	Path() string
}

const (
	StorageDirContent       = ContentQEMUData & ContentContainerData & ContentISO & ContentContainerTemplate & ContentBackup & ContentSnippet
	StorageDirImageFormat   = ImageFormatRaw & ImageFormatQcow2 & ImageFormatVMDK
	StorageDirAllowShare    = AllowShareNever
	StorageDirAllowSnapshot = AllowSnapshotQcow2
	StorageDirAllowClone    = AllowCloneQcow2
)

type StorageLVM interface {
	Storage

	BaseStorage() string
	VolumeGroup() string

	SafeRemove() bool
	SafeRemoveThroughput() uint
}

const (
	StorageLVMContent       = ContentQEMUData & ContentContainerData
	StorageLVMImageFormat   = ImageFormatRaw
	StorageLVMAllowShare    = AllowSharePossible
	StorageLVMAllowSnapshot = AllowSnapshotNever
	StorageLVMAllowClone    = AllowCloneNever
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

type StorageNFS interface {
	Storage

	Server() string
	ExportPath() string
	LocalPath() string
	NFSVersion() NFSVersion
}

const (
	StorageNFSContents    = ContentQEMUData & ContentContainerData & ContentContainerTemplate & ContentISO & ContentBackup & ContentSnippet
	StorageNFSImageFormat = ImageFormatRaw & ImageFormatQcow2 & ImageFormatVMDK
	StorageNFSShared      = AllowShareForced
	StorageNFSSnapshots   = AllowSnapshotQcow2
	StorageNFSClones      = AllowCloneQcow2
)

type StorageCIFS interface {
	Storage

	Server() string
	Share() string
	Domain() string
	Username() string
	Password() string
	SMBVersion() SMBVersion
}

const (
	StorageCIFSContents    = ContentQEMUData & ContentContainerData & ContentContainerTemplate & ContentISO & ContentBackup & ContentSnippet
	StorageCIFSImageFormat = ImageFormatRaw & ImageFormatQcow2 & ImageFormatVMDK
	StorageCIFSShared      = AllowShareForced
	StorageCIFSSnapshots   = AllowSnapshotQcow2
	StorageCIFSClones      = AllowCloneQcow2
)

type StorageGlusterFS interface {
	Storage

	Server() string
	BackupServer() string
	Volume() string
	Transport() GlusterFSTransport
}

const (
	StorageGlusterFSContents    = ContentQEMUData & ContentContainerTemplate & ContentISO & ContentBackup & ContentSnippet
	StorageGlusterFSImageFormat = ImageFormatRaw & ImageFormatQcow2 & ImageFormatVMDK
	StorageGlusterFSShared      = AllowShareForced
	StorageGlusterFSSnapshots   = AllowSnapshotQcow2
	StorageGlusterFSClones      = AllowSnapshotQcow2
)

type StorageISCSIKernelMode interface {
	Storage

	Portal() string
	Target() string
}

const (
	StorageISCSIKernelModeContents    = ContentQEMUData
	StorageISCSIKernelModeImageFormat = ImageFormatRaw
	StorageISCSIKernelModeShared      = AllowShareForced
	StorageISCSIKernelModeSnapshots   = AllowSnapshotNever
	StorageISCSIKernelModeClones      = AllowCloneNever
)

type StorageISCSIUserMode interface {
	Storage

	Portal() string
	Target() string
}

const (
	StorageISCSIUserModeContents    = ContentQEMUData
	StorageISCSIUserModeImageFormat = ImageFormatRaw
	StorageISCSIUserModeShared      = AllowShareForced
	StorageISCSIUserModeSnapshots   = AllowSnapshotNever
	StorageISCSIUserModeClones      = AllowCloneNever
)

type StorageCephFS interface {
	Storage

	MonitorHosts() []string
	SubDirectory() string
	Username() string
	LocalPath() string
	UseFUSE() bool
}

const (
	StorageCephFSContents    = ContentContainerTemplate & ContentISO & ContentBackup & ContentSnippet
	StorageCephFSImageFormat = ContentNone
	StorageCephFSShared      = AllowShareNever
	StorageCephFSSnapshots   = AllowSnapshotAll
	StorageCephFSClones      = AllowCloneNever
)

type StorageRBD interface {
	Storage

	MonitorHosts() []string
	PoolName() string
	Username() string
	UseKRBD() bool
}

const (
	StorageRBDContents    = ContentQEMUData & ContentContainerData
	StorageRBDImageFormat = ImageFormatRaw
	StorageRBDShared      = AllowShareForced
	StorageRBDSnapshots   = AllowSnapshotAll
	StorageRBDClones      = AllowCloneAll
)

type StorageZFS interface {
	Storage

	PoolName() string
	BlockSize() []string
	LocalPath() string
	UseSparse() bool
}

const (
	StorageZFSContents    = ContentQEMUData & ContentContainerData
	StorageZFSImageFormat = ImageFormatRaw & ImageFormatSubVolume
	StorageZFSShared      = AllowShareNever
	StorageZFSSnapshots   = AllowSnapshotAll
	StorageZFSClones      = AllowCloneAll
)
