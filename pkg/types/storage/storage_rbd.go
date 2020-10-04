package storage

import (
	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageRBD interface {
	Storage

	MonitorHosts() []string
	Username() string

	// Always access rbd through krbd kernel module
	UseKRBD() bool

	PoolName() string
}

type StorageRBDProperties struct {
	MonitorHosts []string
	Username     string

	UseKRBD bool

	PoolName string
}

func NewStorageRBDProperties(
	props types.Properties,
) (*StorageRBDProperties, error) {
	obj := new(StorageRBDProperties)

	if v, ok := props["monhost"].(string); ok {
		monitorHosts := internal_types.PVEList{Separator: " "}
		if err := (&monitorHosts).Unmarshal(v); err != nil {
			err := errors.ErrInvalidProperty
			err.AddKey("name", "monhost")
			err.AddKey("value", v)
			return nil, err
		}

		obj.MonitorHosts = monitorHosts.List()
	} else {
		obj.MonitorHosts = DefaultStorageRBDMonitorHosts
	}

	if v, ok := props["username"].(string); ok {
		obj.Username = v
	} else {
		obj.Username = DefaultStorageRBDUsername
	}

	if v, ok := props["krbd"].(float64); ok {
		obj.UseKRBD = internal_types.NewPVEBoolFromFloat64(v).Bool()
	} else {
		obj.UseKRBD = DefaultStorageRBDUseKRBD
	}

	if v, ok := props["pool"].(string); ok {
		obj.PoolName = v
	} else {
		obj.PoolName = DefaultStorageRBDPoolName
	}

	return obj, nil
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
