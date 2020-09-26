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

type StorageProperties struct {
	Content  storage.Content
	Shared   bool
	Disabled bool

	ImageFormat     storage.ImageFormat
	MaxBackupsPerVM uint

	Nodes []string

	ExtraProperties map[string]interface{}
}

func NewStorage(
	svc *Service,
	name string,
	kind storage.Kind,
	props StorageProperties,
) (storage.Storage, error) {
	obj := &Storage{
		svc:  svc,
		name: name,
		kind: kind,
	}

	switch kind {
	case storage.KindDir:
		return NewStorageDir(*obj, props.ExtraProperties)
	case storage.KindLVM:
		return NewStorageLVM(*obj, props.ExtraProperties)
	case storage.KindLVMThin:
		return NewStorageLVMThin(*obj, props.ExtraProperties)
	case storage.KindNFS:
		return NewStorageNFS(*obj, props.ExtraProperties)
	case storage.KindCIFS:
		return NewStorageCIFS(*obj, props.ExtraProperties)
	default:
		return nil, fmt.Errorf("unsupported storage type")
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

	path string
}

func NewStorageDir(obj Storage, props map[string]interface{}) (*StorageDir, error) {
	path, ok := props["path"].(string)
	if !ok {
		return nil, fmt.Errorf("can't find property path")
	}

	return &StorageDir{
		Storage: obj,
		path:    path,
	}, nil
}

type StorageLVM struct {
	Storage

	baseStorage string
	volumeGroup string

	safeRemove           bool
	safeRemoveThroughput int

	taggedOnly bool
}

func NewStorageLVM(obj Storage, props map[string]interface{}) (*StorageLVM, error) {
	baseStorage, ok := props["base"].(string)
	if !ok {
		return nil, fmt.Errorf("can't find property base")
	}

	volumeGroup, ok := props["vgname"].(string)
	if !ok {
		return nil, fmt.Errorf("can't find property vgname")
	}

	safeRemove, ok := props["saferemove"].(types.PVEBool)
	if !ok {
		safeRemove = types.PVEBool(false)
	}

	safeRemoveThroughput, ok := props["saferemove_throughput"].(int)
	if !ok {
		safeRemoveThroughput = -10485760
	}

	taggedOnly, ok := props["tagged_only"].(types.PVEBool)
	if !ok {
		taggedOnly = types.PVEBool(false)
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

type StorageLVMThin struct {
	Storage

	thinPool    string
	volumeGroup string
}

func NewStorageLVMThin(obj Storage, props map[string]interface{}) (*StorageLVMThin, error) {
	thinPool, ok := props["thinpool"].(string)
	if !ok {
		return nil, fmt.Errorf("can't find property thinpool")
	}

	volumeGroup, ok := props["vgname"].(string)
	if !ok {
		return nil, fmt.Errorf("can't find property vgname")
	}

	return &StorageLVMThin{
		Storage: obj,

		thinPool:    thinPool,
		volumeGroup: volumeGroup,
	}, nil
}

type StorageNFS struct {
	Storage

	server     string
	nfsVersion storage.NFSVersion

	serverPath      string
	localPath       string
	createLocalPath bool
}

func NewStorageNFS(obj Storage, props map[string]interface{}) (*StorageNFS, error) {
	server, ok := props["server"].(string)
	if !ok {
		return nil, fmt.Errorf("can't find property server")
	}

	serverPath, ok := props["export"].(string)
	if !ok {
		return nil, fmt.Errorf("can't find property export")
	}

	localPath, ok := props["path"].(string)
	if !ok {
		return nil, fmt.Errorf("can't find property path")
	}

	createLocalPath, ok := props["mkdir"].(types.PVEBool)
	if !ok {
		createLocalPath = types.PVEBool(false)
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

func NewStorageCIFS(obj Storage, props map[string]interface{}) (*StorageCIFS, error) {
	server, ok := props["server"].(string)
	if !ok {
		return nil, fmt.Errorf("can't find property server")
	}

	version, ok := props["smbversion"].(storage.SMBVersion)
	if !ok {
		version = storage.SMBVersion30
	}

	domain, ok := props["domain"].(string)
	if !ok {
		domain = ""
	}

	username, ok := props["username"].(string)
	if !ok {
		username = ""
	}

	password, ok := props["username"].(string)
	if !ok {
		password = ""
	}

	serverShare, ok := props["share"].(string)
	if !ok {
		return nil, fmt.Errorf("can't find property share")
	}

	localPath, ok := props["path"].(string)
	if !ok {
		return nil, fmt.Errorf("can't find property path")
	}

	createLocalPath, ok := props["mkdir"].(types.PVEBool)
	if !ok {
		createLocalPath = types.PVEBool(false)
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
