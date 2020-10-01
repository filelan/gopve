package storage

import (
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
	props *storage.StorageDirProperties

	localPath          string
	localPathCreate    bool
	localPathIsManaged bool
}

func NewStorageDir(
	obj Storage,
	props storage.ExtraProperties,
) (*StorageDir, error) {
	if storageProps, err := storage.NewStorageDirProperties(props); err == nil {
		return &StorageDir{
			Storage: obj,
			props:   storageProps,
		}, nil
	} else {
		return nil, err
	}
}

func (obj *StorageDir) LocalPath() string {
	return obj.props.LocalPath
}

func (obj *StorageDir) LocalPathCreate() bool {
	return obj.props.LocalPathCreate
}

func (obj *StorageDir) LocalPathIsManaged() bool {
	return obj.props.LocalPathIsManaged
}

type StorageLVM struct {
	Storage
	props *storage.StorageLVMProperties
}

func NewStorageLVM(
	obj Storage,
	props storage.ExtraProperties,
) (*StorageLVM, error) {
	if storageProps, err := storage.NewStorageLVMProperties(props); err == nil {
		return &StorageLVM{
			Storage: obj,
			props:   storageProps,
		}, nil
	} else {
		return nil, err
	}
}

func (obj *StorageLVM) BaseStorage() string {
	return obj.props.BaseStorage
}

func (obj *StorageLVM) VolumeGroup() string {
	return obj.props.VolumeGroup
}

func (obj *StorageLVM) SafeRemove() bool {
	return obj.props.SafeRemove
}

func (obj *StorageLVM) SafeRemoveThroughput() int {
	return obj.props.SafeRemoveThroughput
}

func (obj *StorageLVM) TaggedOnly() bool {
	return obj.props.TaggedOnly
}

type StorageLVMThin struct {
	Storage
	props *storage.StorageLVMThinProperties
}

func NewStorageLVMThin(
	obj Storage,
	props storage.ExtraProperties,
) (*StorageLVMThin, error) {
	if storageProps, err := storage.NewStorageLVMThinProperties(props); err == nil {
		return &StorageLVMThin{
			Storage: obj,
			props:   storageProps,
		}, nil
	} else {
		return nil, err
	}
}

func (obj *StorageLVMThin) VolumeGroup() string {
	return obj.props.VolumeGroup
}

func (obj *StorageLVMThin) ThinPool() string {
	return obj.props.ThinPool
}

type StorageZFS struct {
	Storage
	props *storage.StorageZFSProperties
}

func NewStorageZFS(
	obj Storage,
	props storage.ExtraProperties,
) (*StorageZFS, error) {
	if storageProps, err := storage.NewStorageZFSProperties(props); err == nil {
		return &StorageZFS{
			Storage: obj,
			props:   storageProps,
		}, nil
	} else {
		return nil, err
	}
}

func (obj *StorageZFS) PoolName() string {
	return obj.props.PoolName
}

func (obj *StorageZFS) BlockSize() string {
	return obj.props.BlockSize
}

func (obj *StorageZFS) UseSparse() bool {
	return obj.props.UseSparse
}

func (obj *StorageZFS) LocalPath() string {
	return obj.props.LocalPath
}

type StorageNFS struct {
	Storage
	props *storage.StorageNFSProperties
}

func NewStorageNFS(
	obj Storage,
	props storage.ExtraProperties,
) (*StorageNFS, error) {
	if storageProps, err := storage.NewStorageNFSProperties(props); err == nil {
		return &StorageNFS{
			Storage: obj,
			props:   storageProps,
		}, nil
	} else {
		return nil, err
	}
}

func (obj *StorageNFS) Server() string {
	return obj.props.Server
}

func (obj *StorageNFS) NFSVersion() storage.NFSVersion {
	return obj.props.NFSVersion
}

func (obj *StorageNFS) ServerPath() string {
	return obj.props.ServerPath
}

func (obj *StorageNFS) LocalPath() string {
	return obj.props.LocalPath
}

func (obj *StorageNFS) LocalPathCreate() bool {
	return obj.props.LocalPathCreate
}

type StorageCIFS struct {
	Storage
	props *storage.StorageCIFSProperties
}

func NewStorageCIFS(
	obj Storage,
	props storage.ExtraProperties,
) (*StorageCIFS, error) {
	if storageProps, err := storage.NewStorageCIFSProperties(props); err == nil {
		return &StorageCIFS{
			Storage: obj,
			props:   storageProps,
		}, nil
	} else {
		return nil, err
	}
}

func (obj *StorageCIFS) Server() string {
	return obj.props.Server
}

func (obj *StorageCIFS) SMBVersion() storage.SMBVersion {
	return obj.props.SMBVersion
}

func (obj *StorageCIFS) Domain() string {
	return obj.props.Domain
}

func (obj *StorageCIFS) Username() string {
	return obj.props.Username
}

func (obj *StorageCIFS) Password() string {
	return obj.props.Password
}

func (obj *StorageCIFS) ServerShare() string {
	return obj.props.ServerShare
}

func (obj *StorageCIFS) LocalPath() string {
	return obj.props.LocalPath
}

func (obj *StorageCIFS) LocalPathCreate() bool {
	return obj.props.LocalPathCreate
}

type StorageGlusterFS struct {
	Storage
	props *storage.StorageGlusterFSProperties
}

func NewStorageGlusterFS(
	obj Storage,
	props storage.ExtraProperties,
) (*StorageGlusterFS, error) {
	if storageProps, err := storage.NewStorageGlusterFSProperties(props); err == nil {
		return &StorageGlusterFS{
			Storage: obj,
			props:   storageProps,
		}, nil
	} else {
		return nil, err
	}
}

func (obj *StorageGlusterFS) MainServer() string {
	return obj.props.MainServer
}

func (obj *StorageGlusterFS) BackupServer() string {
	return obj.props.BackupServer
}

func (obj *StorageGlusterFS) Transport() storage.GlusterFSTransport {
	return obj.props.Transport
}

func (obj *StorageGlusterFS) Volume() string {
	return obj.props.Volume
}

type StorageISCSIKernel struct {
	Storage
	props *storage.StorageISCSIKernelProperties

	portal string
	target string
}

func NewStorageISCSIKernel(
	obj Storage,
	props storage.ExtraProperties,
) (*StorageISCSIKernel, error) {
	if storageProps, err := storage.NewStorageISCSIKernelProperties(props); err == nil {
		return &StorageISCSIKernel{
			Storage: obj,
			props:   storageProps,
		}, nil
	} else {
		return nil, err
	}
}

func (obj *StorageISCSIKernel) Portal() string {
	return obj.props.Portal
}

func (obj *StorageISCSIKernel) Target() string {
	return obj.props.Target
}

type StorageISCSIUser struct {
	Storage
	props *storage.StorageISCSIUserProperties
}

func NewStorageISCSIUser(
	obj Storage,
	props storage.ExtraProperties,
) (*StorageISCSIUser, error) {
	if storageProps, err := storage.NewStorageISCSIUserProperties(props); err == nil {
		return &StorageISCSIUser{
			Storage: obj,
			props:   storageProps,
		}, nil
	} else {
		return nil, err
	}
}

func (obj *StorageISCSIUser) Portal() string {
	return obj.props.Portal
}

func (obj *StorageISCSIUser) Target() string {
	return obj.props.Target
}

type StorageCephFS struct {
	Storage
	props *storage.StorageCephFSProperties
}

func NewStorageCephFS(
	obj Storage,
	props storage.ExtraProperties,
) (*StorageCephFS, error) {
	if storageProps, err := storage.NewStorageCephFSProperties(props); err == nil {
		return &StorageCephFS{
			Storage: obj,
			props:   storageProps,
		}, nil
	} else {
		return nil, err
	}
}

func (obj *StorageCephFS) MonitorHosts() []string {
	return obj.props.MonitorHosts
}

func (obj *StorageCephFS) Username() string {
	return obj.props.Username
}

func (obj *StorageCephFS) UseFUSE() bool {
	return obj.props.UseFUSE
}

func (obj *StorageCephFS) ServerPath() string {
	return obj.props.ServerPath
}

func (obj *StorageCephFS) LocalPath() string {
	return obj.props.LocalPath
}

type StorageRBD struct {
	Storage
	props *storage.StorageRBDProperties
}

func NewStorageRBD(
	obj Storage,
	props storage.ExtraProperties,
) (*StorageRBD, error) {
	if storageProps, err := storage.NewStorageRBDProperties(props); err == nil {
		return &StorageRBD{
			Storage: obj,
			props:   storageProps,
		}, nil
	} else {
		return nil, err
	}
}

func (obj *StorageRBD) MonitorHosts() []string {
	return obj.props.MonitorHosts
}

func (obj *StorageRBD) Username() string {
	return obj.props.Username
}

func (obj *StorageRBD) UseKRBD() bool {
	return obj.props.UseKRBD
}

func (obj *StorageRBD) PoolName() string {
	return obj.props.PoolName
}

type StorageDRBD struct {
	Storage
	props *storage.StorageDRBDProperties

	redundancy uint
}

func NewStorageDRBD(
	obj Storage,
	props storage.ExtraProperties,
) (*StorageDRBD, error) {
	if storageProps, err := storage.NewStorageDRBDProperties(props); err == nil {
		return &StorageDRBD{
			Storage: obj,
			props:   storageProps,
		}, nil
	} else {
		return nil, err
	}
}

func (obj *StorageDRBD) Redundancy() uint {
	return obj.props.Redundancy
}

type StorageZFSOverISCSI struct {
	Storage
	props *storage.StorageZFSOverISCSIProperties
}

func NewStorageZFSOverISCSI(
	obj Storage,
	props storage.ExtraProperties,
) (*StorageZFSOverISCSI, error) {
	if storageProps, err := storage.NewStorageZFSOverISCSIProperties(props); err == nil {
		return &StorageZFSOverISCSI{
			Storage: obj,
			props:   storageProps,
		}, nil
	} else {
		return nil, err
	}
}

func (obj *StorageZFSOverISCSI) Portal() string {
	return obj.props.Portal
}

func (obj *StorageZFSOverISCSI) Target() string {
	return obj.props.Target
}

func (obj *StorageZFSOverISCSI) PoolName() string {
	return obj.props.PoolName
}

func (obj *StorageZFSOverISCSI) BlockSize() string {
	return obj.props.BlockSize
}

func (obj *StorageZFSOverISCSI) UseSparse() bool {
	return obj.props.UseSparse
}

func (obj *StorageZFSOverISCSI) WriteCache() bool {
	return obj.props.WriteCache
}

func (obj *StorageZFSOverISCSI) ISCSIProvider() storage.ISCSIProvider {
	return obj.props.ISCSIProvider
}

func (obj *StorageZFSOverISCSI) ComstarHostGroup() string {
	return obj.props.ComstarHostGroup
}

func (obj *StorageZFSOverISCSI) ComstarTargetGroup() string {
	return obj.props.ComstarTargetGroup
}

func (obj *StorageZFSOverISCSI) LIOTargetPortalGroup() string {
	return obj.props.LIOTargetPortalGroup
}
