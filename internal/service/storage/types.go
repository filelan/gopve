package storage

import (
	"fmt"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

type Storage struct {
	svc *Service

	name    string
	kind    storage.Kind
	content storage.Content

	shared   bool
	disabled bool

	imageFormat     storage.ImageFormat
	maxBackupsPerVM uint

	nodes []string
}

func NewStorage(
	svc *Service,
	name string,
	kind storage.Kind,
	props storage.Properties,
) *Storage {
	return &Storage{
		svc: svc,

		name:    name,
		kind:    kind,
		content: props.Content,

		shared:   props.Shared,
		disabled: props.Disabled,

		imageFormat:     props.ImageFormat,
		maxBackupsPerVM: props.MaxBackupsPerVM,

		nodes: props.Nodes,
	}
}

func NewDynamicStorage(
	svc *Service,
	name string,
	kind storage.Kind,
	props storage.Properties,
) (storage.Storage, error) {
	obj := NewStorage(svc, name, kind, props)

	switch kind {
	case storage.KindDir:
		return NewStorageDir(*obj, props.ExtraProperties)
	case storage.KindLVM:
		return NewStorageLVM(*obj, props.ExtraProperties)
	case storage.KindLVMThin:
		return NewStorageLVMThin(*obj, props.ExtraProperties)
	case storage.KindZFS:
		return NewStorageZFS(*obj, props.ExtraProperties)
	case storage.KindNFS:
		return NewStorageNFS(*obj, props.ExtraProperties)
	case storage.KindCIFS:
		return NewStorageCIFS(*obj, props.ExtraProperties)
	case storage.KindGlusterFS:
		return NewStorageGlusterFS(*obj, props.ExtraProperties)
	case storage.KindISCSIKernelMode:
		return NewStorageISCSIKernelMode(*obj, props.ExtraProperties)
	case storage.KindISCSIUserMode:
		return NewStorageISCSIUserMode(*obj, props.ExtraProperties)
	case storage.KindCephFS:
		return NewStorageCephFS(*obj, props.ExtraProperties)
	case storage.KindRBD:
		return NewStorageRBD(*obj, props.ExtraProperties)
	case storage.KindZFSOverISCSI:
		return NewStorageZFSOverISCSI(*obj, props.ExtraProperties)
	default:
		return nil, storage.ErrInvalidKind
	}
}

func (obj *Storage) Name() string {
	return obj.name
}

func (obj *Storage) Kind() (storage.Kind, error) {
	return obj.kind, nil
}

func (obj *Storage) Content() (storage.Content, error) {
	return obj.content, nil
}

func (obj *Storage) Shared() (bool, error) {
	return obj.shared, nil
}

func (obj *Storage) Disabled() (bool, error) {
	return obj.disabled, nil
}

func (obj *Storage) ImageFormat() (storage.ImageFormat, error) {
	return obj.imageFormat, nil
}

func (obj *Storage) MaxBackupsPerVM() (uint, error) {
	return obj.maxBackupsPerVM, nil
}

func (obj *Storage) Nodes() ([]string, error) {
	return obj.nodes, nil
}

type StorageDir struct {
	Storage

	localPath          string
	localPathCreate    bool
	localPathIsManaged bool
}

func NewStorageDir(
	obj Storage,
	props map[string]interface{},
) (*StorageDir, error) {
	localPath, ok := props["path"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "path")
		return nil, err
	}

	var localPathCreate types.PVEBool
	if v, ok := props["mkdir"].(int); ok {
		localPathCreate = types.NewPVEBoolFromInt(v)
	} else {
		localPathCreate = types.PVEBool(storage.DefaultStorageDirLocalPathCreate)
	}

	var localPathIsManaged types.PVEBool
	if v, ok := props["is_mountpoint"].(int); ok {
		localPathIsManaged = types.NewPVEBoolFromInt(v)
	} else {
		localPathIsManaged = types.PVEBool(storage.DefaultStorageDirLocalIsManaged)
	}

	return &StorageDir{
		Storage: obj,

		localPath:          localPath,
		localPathCreate:    localPathCreate.Bool(),
		localPathIsManaged: localPathIsManaged.Bool(),
	}, nil
}

func (obj *StorageDir) LocalPath() string {
	return obj.localPath
}

func (obj *StorageDir) LocalPathCreate() bool {
	return obj.localPathCreate
}

func (obj *StorageDir) LocalPathIsManaged() bool {
	return obj.localPathIsManaged
}

type StorageLVM struct {
	Storage

	baseStorage string
	volumeGroup string

	safeRemove           bool
	safeRemoveThroughput int

	taggedOnly bool
}

func NewStorageLVM(
	obj Storage,
	props map[string]interface{},
) (*StorageLVM, error) {
	baseStorage, ok := props["base"].(string)
	if !ok {
		baseStorage = storage.DefaultStorageLVMBaseStorage
	}

	volumeGroup, ok := props["vgname"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "vgname")
		return nil, err
	}

	var safeRemove types.PVEBool
	if v, ok := props["saferemove"].(int); ok {
		safeRemove = types.NewPVEBoolFromInt(v)
	} else {
		safeRemove = types.PVEBool(storage.DefaultStorageLVMSafeRemove)
	}

	safeRemoveThroughput, ok := props["saferemove_throughput"].(int)
	if !ok {
		safeRemoveThroughput = storage.DefaultStorageLVMSafeRemoveThroughput
	}

	var taggedOnly types.PVEBool
	if v, ok := props["tagged_only"].(int); ok {
		taggedOnly = types.NewPVEBoolFromInt(v)
	} else {
		taggedOnly = types.PVEBool(storage.DefaultStorageLVMTaggedOnly)
	}

	return &StorageLVM{
		Storage: obj,

		baseStorage: baseStorage,
		volumeGroup: volumeGroup,

		safeRemove:           safeRemove.Bool(),
		safeRemoveThroughput: safeRemoveThroughput,

		taggedOnly: taggedOnly.Bool(),
	}, nil
}

func (obj *StorageLVM) BaseStorage() string {
	return obj.baseStorage
}

func (obj *StorageLVM) VolumeGroup() string {
	return obj.volumeGroup
}

func (obj *StorageLVM) SafeRemove() bool {
	return obj.safeRemove
}

func (obj *StorageLVM) SafeRemoveThroughput() int {
	return obj.safeRemoveThroughput
}

func (obj *StorageLVM) TaggedOnly() bool {
	return obj.taggedOnly
}

type StorageLVMThin struct {
	Storage

	thinPool    string
	volumeGroup string
}

func NewStorageLVMThin(
	obj Storage,
	props map[string]interface{},
) (*StorageLVMThin, error) {
	thinPool, ok := props["thinpool"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "thinpool")
		return nil, err
	}

	volumeGroup, ok := props["vgname"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "vgname")
		return nil, err
	}

	return &StorageLVMThin{
		Storage: obj,

		thinPool:    thinPool,
		volumeGroup: volumeGroup,
	}, nil
}

func (obj *StorageLVMThin) ThinPool() string {
	return obj.thinPool
}

func (obj *StorageLVMThin) VolumeGroup() string {
	return obj.volumeGroup
}

type StorageZFS struct {
	Storage

	poolName string

	blockSize uint
	useSparse bool

	localPath string
}

func NewStorageZFS(
	obj Storage,
	props map[string]interface{},
) (*StorageZFS, error) {
	poolName, ok := props["pool"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "pool")
		return nil, err
	}

	var blockSize uint
	if v, ok := props["blocksize"].(int); ok {
		blockSize = uint(v)
	} else {
		blockSize = storage.DefaultStorageZFSBlockSize
	}

	var useSparse types.PVEBool
	if v, ok := props["sparse"].(int); ok {
		useSparse = types.NewPVEBoolFromInt(v)
	} else {
		useSparse = types.PVEBool(storage.DefaultStorageZFSUseSparse)
	}

	localPath, ok := props["mountpoint"].(string)
	if !ok {
		localPath = storage.DefaultStorageZFSMountPoint
	}

	return &StorageZFS{
		Storage: obj,

		poolName: poolName,

		blockSize: blockSize,
		useSparse: useSparse.Bool(),

		localPath: localPath,
	}, nil
}

func (obj *StorageZFS) PoolName() string {
	return obj.poolName
}

func (obj *StorageZFS) BlockSize() uint {
	return obj.blockSize
}

func (obj *StorageZFS) UseSparse() bool {
	return obj.useSparse
}

func (obj *StorageZFS) LocalPath() string {
	return obj.localPath
}

type StorageNFS struct {
	Storage

	server     string
	nfsVersion storage.NFSVersion

	serverPath      string
	localPath       string
	createLocalPath bool
}

func NewStorageNFS(
	obj Storage,
	props map[string]interface{},
) (*StorageNFS, error) {
	server, ok := props["server"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "server")
		return nil, err
	}

	serverPath, ok := props["export"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "export")
		return nil, err
	}

	localPath, ok := props["path"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "path")
		return nil, err
	}

	var createLocalPath types.PVEBool
	if v, ok := props["mkdir"].(int); ok {
		createLocalPath = types.NewPVEBoolFromInt(v)
	} else {
		createLocalPath = types.PVEBool(storage.DefaultStorageNFSCreateLocalPath)
	}

	return &StorageNFS{
		Storage: obj,

		server: server,

		serverPath:      serverPath,
		localPath:       localPath,
		createLocalPath: createLocalPath.Bool(),
	}, nil
}

type StorageCIFS struct {
	Storage

	server  string
	version storage.SMBVersion

	domain   string
	username string
	password string

	serverShare     string
	localPath       string
	createLocalPath bool
}

func NewStorageCIFS(
	obj Storage,
	props map[string]interface{},
) (*StorageCIFS, error) {
	server, ok := props["server"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "server")
		return nil, err
	}

	version, ok := props["smbversion"].(storage.SMBVersion)
	if !ok {
		version = storage.DefaultStorageCIFSSMBVersion
	}

	domain, ok := props["domain"].(string)
	if !ok {
		domain = storage.DefaultStorageCIFSDomain
	}

	username, ok := props["username"].(string)
	if !ok {
		username = storage.DefaultStorageCIFSUsername
	}

	password, ok := props["username"].(string)
	if !ok {
		password = storage.DefaultStorageCIFSPassword
	}

	serverShare, ok := props["share"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "share")
		return nil, err
	}

	localPath, ok := props["path"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "path")
		return nil, err
	}

	var createLocalPath types.PVEBool
	if v, ok := props["mkdir"].(int); ok {
		createLocalPath = types.NewPVEBoolFromInt(v)
	} else {
		createLocalPath = types.PVEBool(storage.DefaultStorageCIFSCreateLocalPath)
	}

	return &StorageCIFS{
		Storage: obj,

		server:  server,
		version: version,

		domain:   domain,
		username: username,
		password: password,

		serverShare:     serverShare,
		localPath:       localPath,
		createLocalPath: createLocalPath.Bool(),
	}, nil
}

type StorageGlusterFS struct {
	Storage

	mainServer   string
	backupServer string
	transport    storage.GlusterFSTransport

	volume string
}

func NewStorageGlusterFS(
	obj Storage,
	props map[string]interface{},
) (*StorageGlusterFS, error) {
	mainServer, ok := props["server"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "server")
		return nil, err
	}

	backupServer, ok := props["server2"].(string)
	if !ok {
		backupServer = storage.DefaultStorageGlusterFSBackupService
	}

	transport, ok := props["transport"].(storage.GlusterFSTransport)
	if !ok {
		transport = storage.DefaultStorageGlusterFSTransport
	}

	volume, ok := props["volume"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "volume")
		return nil, err
	}

	return &StorageGlusterFS{
		Storage: obj,

		mainServer:   mainServer,
		backupServer: backupServer,
		transport:    transport,

		volume: volume,
	}, nil
}

type StorageISCSIKernel struct {
	Storage

	portal string
	target string
}

func NewStorageISCSIKernelMode(
	obj Storage,
	props map[string]interface{},
) (*StorageISCSIKernel, error) {
	portal, ok := props["portal"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "portal")
		return nil, err
	}

	target, ok := props["target"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "target")
		return nil, err
	}

	return &StorageISCSIKernel{
		Storage: obj,

		portal: portal,
		target: target,
	}, nil
}

type StorageISCSIUser struct {
	Storage

	portal string
	target string
}

func NewStorageISCSIUserMode(
	obj Storage,
	props map[string]interface{},
) (*StorageISCSIUser, error) {
	portal, ok := props["portal"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "portal")
		return nil, err
	}

	target, ok := props["target"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "target")
		return nil, err
	}

	return &StorageISCSIUser{
		Storage: obj,

		portal: portal,
		target: target,
	}, nil
}

type StorageCephFS struct {
	Storage

	monitorHosts []string
	username     string

	useFUSE bool

	serverPath string
	localPath  string
}

func NewStorageCephFS(
	obj Storage,
	props map[string]interface{},
) (*StorageCephFS, error) {
	monitorHosts := types.PVEStringList{Separator: " "}
	hosts, ok := props["monhost"].(string)
	if ok {
		if err := (&monitorHosts).Unmarshal(hosts); err != nil {
			return nil, fmt.Errorf("invalid monhost value")
		}
	}

	username, ok := props["username"].(string)
	if !ok {
		username = storage.DefaultStorageCephFSUsername
	}

	var useFUSE types.PVEBool
	if v, ok := props["fuse"].(int); ok {
		useFUSE = types.NewPVEBoolFromInt(v)
	} else {
		useFUSE = types.PVEBool(storage.DefaultStorageCephFSUseFUSE)
	}

	serverPath, ok := props["subdir"].(string)
	if !ok {
		serverPath = storage.DefaultStorageCephFSServerPath
	}

	localPath, ok := props["path"].(string)
	if !ok {
		localPath = fmt.Sprintf(storage.DefaultStorageCephFSLocalPath, obj.name)
	}

	return &StorageCephFS{
		Storage: obj,

		monitorHosts: monitorHosts.List(),
		username:     username,

		useFUSE: useFUSE.Bool(),

		serverPath: serverPath,
		localPath:  localPath,
	}, nil
}

type StorageRBD struct {
	Storage

	monitorHosts []string
	username     string

	useKRBD bool

	poolName string
}

func NewStorageRBD(
	obj Storage,
	props map[string]interface{},
) (*StorageRBD, error) {
	monitorHosts := types.PVEStringList{Separator: " "}
	hosts, ok := props["monhost"].(string)
	if ok {
		if err := (&monitorHosts).Unmarshal(hosts); err != nil {
			return nil, fmt.Errorf("invalid monhost value")
		}
	}

	username, ok := props["username"].(string)
	if !ok {
		username = storage.DefaultStorageRBDUsername
	}

	var useKRBD types.PVEBool
	if v, ok := props["krbd"].(int); ok {
		useKRBD = types.NewPVEBoolFromInt(v)
	} else {
		useKRBD = types.PVEBool(storage.DefaultStorageRBDUseKRBD)
	}

	poolName, ok := props["pool"].(string)
	if !ok {
		poolName = storage.DefaultStorageRBDPoolName
	}

	return &StorageRBD{
		Storage: obj,

		monitorHosts: monitorHosts.List(),
		username:     username,

		useKRBD: useKRBD.Bool(),

		poolName: poolName,
	}, nil
}

type StorageZFSOverISCSI struct {
	Storage

	portal string
	target string

	poolName string

	blockSize  uint
	useSparse  bool
	writeCache bool

	iSCSIProvider storage.ISCSIProvider

	comstarHostGroup     string
	comstarTargetGroup   string
	lioTargetPortalGroup string
}

func NewStorageZFSOverISCSI(
	obj Storage,
	props map[string]interface{},
) (*StorageZFSOverISCSI, error) {
	portal, ok := props["portal"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "portal")
		return nil, err
	}

	target, ok := props["target"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "target")
		return nil, err
	}

	poolName, ok := props["pool"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "pool")
		return nil, err
	}

	blockSize, ok := props["blocksize"].(uint)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "blocksize")
		return nil, err
	}

	var useSparse types.PVEBool
	if v, ok := props["sparse"].(int); ok {
		useSparse = types.NewPVEBoolFromInt(v)
	} else {
		useSparse = types.PVEBool(storage.DefaultStorageZFSOverISCSIUseSparse)
	}

	var writeCache types.PVEBool
	if v, ok := props["nowritecache"].(int); ok {
		writeCache = !types.NewPVEBoolFromInt(v)
	} else {
		writeCache = types.PVEBool(storage.DefaultStorageZFSOverISCSIWriteCache)
	}

	iSCSIProvider, ok := props["iscsiprovider"].(storage.ISCSIProvider)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "iscsiprovider")
		return nil, err
	}

	comstarHostGroup, ok := props["comstar_hg"].(string)
	if !ok {
		comstarHostGroup = storage.DefaultStorageZFSOverISCSIComstarHostGroup
	}

	comstarTargetGroup, ok := props["comstar_tg"].(string)
	if !ok {
		comstarTargetGroup = storage.DefaultStorageZFSOverISCSIComstarTargetGroup
	}

	lioTargetPortalGroup, ok := props["lio_tpg"].(string)
	if !ok {
		lioTargetPortalGroup = storage.DefaultStorageZFSOverISCSILIOTargetPortalGroup
	}

	return &StorageZFSOverISCSI{
		Storage: obj,

		portal: portal,
		target: target,

		poolName: poolName,

		blockSize:  blockSize,
		useSparse:  useSparse.Bool(),
		writeCache: writeCache.Bool(),

		iSCSIProvider: iSCSIProvider,

		comstarHostGroup:     comstarHostGroup,
		comstarTargetGroup:   comstarTargetGroup,
		lioTargetPortalGroup: lioTargetPortalGroup,
	}, nil
}
