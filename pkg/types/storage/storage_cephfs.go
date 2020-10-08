package storage

import (
	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

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

type StorageCephFSProperties struct {
	MonitorHosts []string
	Username     string

	UseFUSE bool

	ServerPath string
	LocalPath  string
}

const (
	mkCephFSMonitorHosts = "monhost"
	mkCephFSUsername     = "username"
	mkCephFSUseFUSE      = "fuse"
	mkCephFSServerPath   = "subdir"
	mkCephFSLocalPath    = "path"
)

var DefaultStorageCephFSMonitorHosts = []string{}

const (
	DefaultStorageCephFSUsername   = ""
	DefaultStorageCephFSUseFUSE    = false
	DefaultStorageCephFSServerPath = "/"
)

func NewStorageCephFSProperties(
	props types.Properties,
) (*StorageCephFSProperties, error) {
	obj := new(StorageCephFSProperties)

	if v, ok := props[mkCephFSMonitorHosts].(string); ok {
		monitorHosts := internal_types.PVEList{Separator: " "}
		if err := (&monitorHosts).Unmarshal(v); err != nil {
			err := errors.ErrInvalidProperty
			err.AddKey("name", mkCephFSMonitorHosts)
			err.AddKey("value", v)
			return nil, err
		}

		obj.MonitorHosts = monitorHosts.List()
	} else {
		obj.MonitorHosts = DefaultStorageCephFSMonitorHosts
	}

	if err := props.SetString(mkCephFSUsername, &obj.Username, DefaultStorageCephFSUsername, nil); err != nil {
		return nil, err
	}

	if err := props.SetBool(mkCephFSUseFUSE, &obj.UseFUSE, DefaultStorageCephFSUseFUSE, nil); err != nil {
		return nil, err
	}

	if err := props.SetString(mkCephFSServerPath, &obj.ServerPath, DefaultStorageCephFSServerPath, nil); err != nil {
		return nil, err
	}

	if err := props.SetRequiredString(mkCephFSLocalPath, &obj.LocalPath, nil); err != nil {
		return nil, err
	}

	return obj, nil
}
