package storage

import (
	"encoding/json"
	"fmt"

	"github.com/xabinapal/gopve/internal/types"
)

type StorageNFS interface {
	Storage

	Server() string
	NFSVersion() NFSVersion

	ServerPath() string
	LocalPath() string
	LocalPathCreate() bool
}

type StorageNFSProperties struct {
	Server     string
	NFSVersion NFSVersion

	ServerPath      string
	LocalPath       string
	LocalPathCreate bool
}

func NewStorageNFSProperties(props ExtraProperties) (*StorageNFSProperties, error) {
	obj := new(StorageNFSProperties)

	if v, ok := props["server"].(string); ok {
		obj.Server = v
	} else {
		err := ErrMissingProperty
		err.AddKey("name", "server")
		return nil, err
	}

	if v, ok := props["options"].(string); ok {
		nfsOptions := types.PVEDictionary{
			ListSeparator:     ",",
			KeyValueSeparator: "=",
			AllowNoValue:      true,
		}

		if err := (&nfsOptions).Unmarshal(v); err != nil {
			err := ErrInvalidProperty
			err.AddKey("name", "options")
			err.AddKey("value", v)
			return nil, err
		}

		for _, option := range nfsOptions.List() {
			if option.Key() == "vers" {
				var nfsVersion NFSVersion
				if err := (&nfsVersion).Unmarshal(option.Value()); err == nil {
					obj.NFSVersion = nfsVersion
				} else {
					err := ErrInvalidProperty
					err.AddKey("name", "options")
					err.AddKey("value", v)
					return nil, err
				}

				break
			}
		}
	} else {
		obj.NFSVersion = DefaultStorageNFSVersion
	}

	if v, ok := props["export"].(string); ok {
		obj.ServerPath = v
	} else {
		err := ErrMissingProperty
		err.AddKey("name", "export")
		return nil, err
	}

	if v, ok := props["path"].(string); ok {
		obj.LocalPath = v
	} else {
		err := ErrMissingProperty
		err.AddKey("name", "path")
		return nil, err
	}

	if v, ok := props["mkdir"].(int); ok {
		obj.LocalPathCreate = types.NewPVEBoolFromInt(v).Bool()
	} else {
		obj.LocalPathCreate = DefaultStorageNFSLocalPathCreate
	}

	return obj, nil
}

const (
	StorageNFSContents    = ContentQEMUData & ContentContainerData & ContentContainerTemplate & ContentISO & ContentBackup & ContentSnippet
	StorageNFSImageFormat = ImageFormatRaw & ImageFormatQcow2 & ImageFormatVMDK
	StorageNFSShared      = AllowShareForced
	StorageNFSSnapshots   = AllowSnapshotQcow2
	StorageNFSClones      = AllowCloneQcow2
)

const (
	DefaultStorageNFSVersion         = NFSVersionNone
	DefaultStorageNFSLocalPathCreate = false
)

type NFSVersion uint

const (
	NFSVersionNone NFSVersion = iota
	NFSVersion30
	NFSVersion40
	NFSVersion41
	NFSVersion42
)

func (obj NFSVersion) Marshal() (string, error) {
	switch obj {
	case NFSVersion30:
		return "3", nil
	case NFSVersion40:
		return "4", nil
	case NFSVersion41:
		return "4.1", nil
	case NFSVersion42:
		return "4.2", nil
	default:
		return "", fmt.Errorf("unknown nfs version")
	}
}

func (obj *NFSVersion) Unmarshal(s string) error {
	switch s {
	case "3":
		*obj = NFSVersion30
	case "4":
		*obj = NFSVersion40
	case "4.1":
		*obj = NFSVersion41
	case "4.2":
		*obj = NFSVersion42
	default:
		return fmt.Errorf("can't unmarshal nfs version %s", s)
	}

	return nil
}

func (obj *NFSVersion) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
