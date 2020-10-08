package storage

import (
	"encoding/json"

	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageNFS interface {
	Storage

	Server() string
	NFSVersion() NFSVersion

	ServerPath() string
	LocalPath() string
	LocalPathCreate() bool
}

const (
	StorageNFSContents    = ContentQEMUData & ContentContainerData & ContentContainerTemplate & ContentISO & ContentBackup & ContentSnippet
	StorageNFSImageFormat = ImageFormatRaw & ImageFormatQcow2 & ImageFormatVMDK
	StorageNFSShared      = AllowShareForced
	StorageNFSSnapshots   = AllowSnapshotQcow2
	StorageNFSClones      = AllowCloneQcow2
)

type StorageNFSProperties struct {
	Server     string
	NFSVersion NFSVersion

	ServerPath      string
	LocalPath       string
	LocalPathCreate bool
}

const (
	mkNFSServer          = "server"
	mkNFSOptions         = "options"
	mkNFSServerPath      = "export"
	mkNFSLocalPath       = "path"
	mkNFSLocalPathCreate = "mkdir"
)

const (
	DefaultStorageNFSVersion         = NFSVersionNone
	DefaultStorageNFSLocalPathCreate = false
)

func NewStorageNFSProperties(
	props types.Properties,
) (*StorageNFSProperties, error) {
	obj := new(StorageNFSProperties)

	if err := props.SetRequiredString(mkNFSServer, &obj.Server, nil); err != nil {
		return nil, err
	}

	if v, ok := props[mkNFSOptions].(string); ok {
		nfsOptions := internal_types.PVEDictionary{
			ListSeparator:     ",",
			KeyValueSeparator: "=",
			AllowNoValue:      true,
		}

		if err := (&nfsOptions).Unmarshal(v); err != nil {
			err := errors.ErrInvalidProperty
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
					err := errors.ErrInvalidProperty
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

	if err := props.SetRequiredString(mkNFSServerPath, &obj.ServerPath, nil); err != nil {
		return nil, err
	}

	if err := props.SetRequiredString(mkNFSLocalPath, &obj.LocalPath, nil); err != nil {
		return nil, err
	}

	if err := props.SetBool(mkNFSLocalPathCreate, &obj.LocalPathCreate, DefaultStorageNFSLocalPathCreate, nil); err != nil {
		return nil, err
	}

	return obj, nil
}

type NFSVersion string

const (
	NFSVersionNone NFSVersion = ""
	NFSVersion30   NFSVersion = "3"
	NFSVersion40   NFSVersion = "4"
	NFSVersion41   NFSVersion = "4.1"
	NFSVersion42   NFSVersion = "4.2"
)

func (obj NFSVersion) IsValid() bool {
	switch obj {
	case NFSVersionNone, NFSVersion30, NFSVersion40, NFSVersion41, NFSVersion42:
		return true
	default:
		return false
	}
}

func (obj NFSVersion) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj NFSVersion) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *NFSVersion) Unmarshal(s string) error {
	*obj = NFSVersion(s)
	return nil
}

func (obj *NFSVersion) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
