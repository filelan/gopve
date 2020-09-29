package storage

import (
	"fmt"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types/storage"
)

type ExtraProperties map[string]interface{}

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
	case storage.KindISCSIKernel:
		return NewStorageISCSIKernel(*obj, props.ExtraProperties)
	case storage.KindISCSIUser:
		return NewStorageISCSIUser(*obj, props.ExtraProperties)
	case storage.KindCephFS:
		return NewStorageCephFS(*obj, props.ExtraProperties)
	case storage.KindRBD:
		return NewStorageRBD(*obj, props.ExtraProperties)
	case storage.KindDRBD:
		return NewStorageDRBD(*obj, props.ExtraProperties)
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
	props ExtraProperties,
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
	props ExtraProperties,
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
	props ExtraProperties,
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
	props ExtraProperties,
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
	localPathCreate bool
}

func NewStorageNFS(
	obj Storage,
	props ExtraProperties,
) (*StorageNFS, error) {
	server, ok := props["server"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "server")
		return nil, err
	}

	nfsVersion := storage.DefaultStorageNFSVersion
	if v, ok := props["options"].(string); ok {
		nfsOptions := types.PVEStringList{Separator: ","}
		if err := (&nfsOptions).Unmarshal(v); err != nil {
			return nil, fmt.Errorf("invalid options value")
		}

		for _, option := range nfsOptions.List() {
			nfsOption := types.PVEStringKV{Separator: "=", AllowNoValue: true}
			if err := (&nfsOption).Unmarshal(option); err != nil {
				return nil, fmt.Errorf("invalid option value")
			}

			if nfsOption.Key() == "vers" {
				if err := (&nfsVersion).Unmarshal(nfsOption.Value()); err != nil {
					return nil, fmt.Errorf("invalid option value")
				}
				break
			}
		}
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

	var localPathCreate types.PVEBool
	if v, ok := props["mkdir"].(int); ok {
		localPathCreate = types.NewPVEBoolFromInt(v)
	} else {
		localPathCreate = types.PVEBool(storage.DefaultStorageNFSLocalPathCreate)
	}

	return &StorageNFS{
		Storage: obj,

		server:     server,
		nfsVersion: nfsVersion,

		serverPath:      serverPath,
		localPath:       localPath,
		localPathCreate: localPathCreate.Bool(),
	}, nil
}

func (obj *StorageNFS) Server() string {
	return obj.server
}

func (obj *StorageNFS) NFSVersion() storage.NFSVersion {
	return obj.nfsVersion
}

func (obj *StorageNFS) ServerPath() string {
	return obj.serverPath
}

func (obj *StorageNFS) LocalPath() string {
	return obj.localPath
}

func (obj *StorageNFS) LocalPathCreate() bool {
	return obj.localPathCreate
}

type StorageCIFS struct {
	Storage

	server     string
	smbVersion storage.SMBVersion

	domain   string
	username string
	password string

	serverShare     string
	localPath       string
	localPathCreate bool
}

func NewStorageCIFS(
	obj Storage,
	props ExtraProperties,
) (*StorageCIFS, error) {
	server, ok := props["server"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "server")
		return nil, err
	}

	var version storage.SMBVersion
	if v, ok := props["smbversion"].(string); ok {
		(&version).Unmarshal(v)
	} else {
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

	password, ok := props["password"].(string)
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

	var localPathCreate types.PVEBool
	if v, ok := props["mkdir"].(int); ok {
		localPathCreate = types.NewPVEBoolFromInt(v)
	} else {
		localPathCreate = types.PVEBool(storage.DefaultStorageCIFSLocalPathCreate)
	}

	return &StorageCIFS{
		Storage: obj,

		server:     server,
		smbVersion: version,

		domain:   domain,
		username: username,
		password: password,

		serverShare:     serverShare,
		localPath:       localPath,
		localPathCreate: localPathCreate.Bool(),
	}, nil
}

func (obj *StorageCIFS) Server() string {
	return obj.server
}

func (obj *StorageCIFS) SMBVersion() storage.SMBVersion {
	return obj.smbVersion
}

func (obj *StorageCIFS) Domain() string {
	return obj.domain
}

func (obj *StorageCIFS) Username() string {
	return obj.username
}

func (obj *StorageCIFS) Password() string {
	return obj.password
}

func (obj *StorageCIFS) ServerShare() string {
	return obj.serverShare
}

func (obj *StorageCIFS) LocalPath() string {
	return obj.localPath
}

func (obj *StorageCIFS) LocalPathCreate() bool {
	return obj.localPathCreate
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
	props ExtraProperties,
) (*StorageGlusterFS, error) {
	mainServer, ok := props["server"].(string)
	if !ok {
		err := storage.ErrMissingProperty
		err.AddKey("name", "server")
		return nil, err
	}

	backupServer, ok := props["server2"].(string)
	if !ok {
		backupServer = storage.DefaultStorageGlusterFSBackupServer
	}

	var transport storage.GlusterFSTransport
	if v, ok := props["transport"].(string); ok {
		(&transport).Unmarshal(v)
	} else {
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

func (obj *StorageGlusterFS) MainServer() string {
	return obj.mainServer
}

func (obj *StorageGlusterFS) BackupServer() string {
	return obj.backupServer
}

func (obj *StorageGlusterFS) Transport() storage.GlusterFSTransport {
	return obj.transport
}

func (obj *StorageGlusterFS) Volume() string {
	return obj.volume
}

type StorageISCSIKernel struct {
	Storage

	portal string
	target string
}

func NewStorageISCSIKernel(
	obj Storage,
	props ExtraProperties,
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

func (obj *StorageISCSIKernel) Portal() string {
	return obj.portal
}

func (obj *StorageISCSIKernel) Target() string {
	return obj.target
}

type StorageISCSIUser struct {
	Storage

	portal string
	target string
}

func NewStorageISCSIUser(
	obj Storage,
	props ExtraProperties,
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

func (obj *StorageISCSIUser) Portal() string {
	return obj.portal
}

func (obj *StorageISCSIUser) Target() string {
	return obj.target
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
	props ExtraProperties,
) (*StorageCephFS, error) {
	var monitorHosts types.PVEStringList
	hosts, ok := props["monhost"].(string)
	if ok {
		monitorHosts = types.PVEStringList{Separator: " "}
		if err := (&monitorHosts).Unmarshal(hosts); err != nil {
			return nil, fmt.Errorf("invalid monhost value")
		}
	} else {
		monitorHosts = types.NewPVEStringList(" ", storage.DefaultStorageCephFSMonitorHosts)
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
		err := storage.ErrMissingProperty
		err.AddKey("name", "path")
		return nil, err
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

func (obj *StorageCephFS) MonitorHosts() []string {
	return obj.monitorHosts
}

func (obj *StorageCephFS) Username() string {
	return obj.username
}

func (obj *StorageCephFS) UseFUSE() bool {
	return obj.useFUSE
}

func (obj *StorageCephFS) ServerPath() string {
	return obj.serverPath
}

func (obj *StorageCephFS) LocalPath() string {
	return obj.localPath
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
	props ExtraProperties,
) (*StorageRBD, error) {
	var monitorHosts types.PVEStringList
	hosts, ok := props["monhost"].(string)
	if ok {
		monitorHosts = types.PVEStringList{Separator: " "}
		if err := (&monitorHosts).Unmarshal(hosts); err != nil {
			return nil, fmt.Errorf("invalid monhost value")
		}
	} else {
		monitorHosts = types.NewPVEStringList(" ", storage.DefaultStorageRBDMonitorHosts)
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

func (obj *StorageRBD) MonitorHosts() []string {
	return obj.monitorHosts
}

func (obj *StorageRBD) Username() string {
	return obj.username
}

func (obj *StorageRBD) UseKRBD() bool {
	return obj.useKRBD
}

func (obj *StorageRBD) PoolName() string {
	return obj.poolName
}

type StorageDRBD struct {
	Storage

	redundancy uint
}

func NewStorageDRBD(
	obj Storage,
	props ExtraProperties,
) (*StorageDRBD, error) {
	var redundancy uint
	if v, ok := props["redundancy"].(int); ok {
		redundancy = uint(v)
	} else {
		redundancy = storage.DefaultStorageDRBDRedundancy
	}

	return &StorageDRBD{
		Storage: obj,

		redundancy: redundancy,
	}, nil
}

func (obj *StorageDRBD) Redundancy() uint {
	return obj.redundancy
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
	props ExtraProperties,
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

	var blockSize uint
	if v, ok := props["blocksize"].(int); ok {
		blockSize = uint(v)
	} else {
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

	var iSCSIProvider storage.ISCSIProvider

	if v, ok := props["iscsiprovider"].(string); ok {
		(&iSCSIProvider).Unmarshal(v)
	} else {
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

func (obj *StorageZFSOverISCSI) Portal() string {
	return obj.portal
}

func (obj *StorageZFSOverISCSI) Target() string {
	return obj.target
}

func (obj *StorageZFSOverISCSI) PoolName() string {
	return obj.poolName
}

func (obj *StorageZFSOverISCSI) BlockSize() uint {
	return obj.blockSize
}

func (obj *StorageZFSOverISCSI) UseSparse() bool {
	return obj.useSparse
}

func (obj *StorageZFSOverISCSI) WriteCache() bool {
	return obj.writeCache
}

func (obj *StorageZFSOverISCSI) ISCSIProvider() storage.ISCSIProvider {
	return obj.iSCSIProvider
}

func (obj *StorageZFSOverISCSI) ComstarHostGroup() string {
	return obj.comstarHostGroup
}

func (obj *StorageZFSOverISCSI) ComstarTargetGroup() string {
	return obj.comstarTargetGroup
}

func (obj *StorageZFSOverISCSI) LIOTargetPortalGroup() string {
	return obj.lioTargetPortalGroup
}
