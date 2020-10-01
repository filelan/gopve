package storage

import (
	"github.com/xabinapal/gopve/internal/types"
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

type StorageCephFSProperties struct {
	MonitorHosts []string
	Username     string

	UseFUSE bool

	ServerPath string
	LocalPath  string
}

func NewStorageCephFSProperties(props ExtraProperties) (*StorageCephFSProperties, error) {
	obj := new(StorageCephFSProperties)

	if v, ok := props["monhost"].(string); ok {
		monitorHosts := types.PVEList{Separator: " "}
		if err := (&monitorHosts).Unmarshal(v); err != nil {
			err := ErrInvalidProperty
			err.AddKey("name", "monhost")
			err.AddKey("value", v)
			return nil, err
		}

		obj.MonitorHosts = monitorHosts.List()
	} else {
		obj.MonitorHosts = DefaultStorageCephFSMonitorHosts
	}

	if v, ok := props["username"].(string); ok {
		obj.Username = v
	} else {
		obj.Username = DefaultStorageCephFSUsername
	}

	if v, ok := props["fuse"].(int); ok {
		obj.UseFUSE = types.NewPVEBoolFromInt(v).Bool()
	} else {
		obj.UseFUSE = DefaultStorageCephFSUseFUSE
	}

	if v, ok := props["subdir"].(string); ok {
		obj.ServerPath = v
	} else {
		obj.ServerPath = DefaultStorageCephFSServerPath
	}

	if v, ok := props["path"].(string); ok {
		obj.LocalPath = v
	} else {
		err := ErrMissingProperty
		err.AddKey("name", "path")
		return nil, err
	}

	return obj, nil
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
