package storage

import (
	"encoding/json"
	"fmt"

	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageCIFS interface {
	Storage

	Server() string
	SMBVersion() SMBVersion

	Domain() string
	Username() string
	Password() string

	ServerShare() string
	LocalPath() string
	LocalPathCreate() bool
}

type StorageCIFSProperties struct {
	Server     string
	SMBVersion SMBVersion

	Domain   string
	Username string
	Password string

	ServerShare     string
	LocalPath       string
	LocalPathCreate bool
}

func NewStorageCIFSProperties(
	props types.Properties,
) (*StorageCIFSProperties, error) {
	obj := new(StorageCIFSProperties)

	if v, ok := props["server"].(string); ok {
		obj.Server = v
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", "server")
		return nil, err
	}

	if v, ok := props["smbversion"].(string); ok {
		if err := (&obj.SMBVersion).Unmarshal(v); err != nil {
			err := errors.ErrInvalidProperty
			err.AddKey("name", "smbversion")
			err.AddKey("value", v)
			return nil, err
		}
	} else {
		obj.SMBVersion = DefaultStorageCIFSSMBVersion
	}

	if v, ok := props["domain"].(string); ok {
		obj.Domain = v
	} else {
		obj.Domain = DefaultStorageCIFSDomain
	}

	if v, ok := props["username"].(string); ok {
		obj.Username = v
	} else {
		obj.Username = DefaultStorageCIFSUsername
	}

	if v, ok := props["password"].(string); ok {
		obj.Password = v
	} else {
		obj.Password = DefaultStorageCIFSPassword
	}

	if v, ok := props["share"].(string); ok {
		obj.ServerShare = v
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", "share")
		return nil, err
	}

	if v, ok := props["path"].(string); ok {
		obj.LocalPath = v
	} else {
		err := errors.ErrMissingProperty
		err.AddKey("name", "path")
		return nil, err
	}

	if v, ok := props["mkdir"].(float64); ok {
		obj.LocalPathCreate = internal_types.NewPVEBoolFromFloat64(v).Bool()
	} else {
		obj.LocalPathCreate = DefaultStorageCIFSLocalPathCreate
	}

	return obj, nil
}

const (
	StorageCIFSContents    = ContentQEMUData & ContentContainerData & ContentContainerTemplate & ContentISO & ContentBackup & ContentSnippet
	StorageCIFSImageFormat = ImageFormatRaw & ImageFormatQcow2 & ImageFormatVMDK
	StorageCIFSShared      = AllowShareForced
	StorageCIFSSnapshots   = AllowSnapshotQcow2
	StorageCIFSClones      = AllowCloneQcow2
)

const (
	DefaultStorageCIFSSMBVersion      = SMBVersion30
	DefaultStorageCIFSDomain          = ""
	DefaultStorageCIFSUsername        = ""
	DefaultStorageCIFSPassword        = ""
	DefaultStorageCIFSLocalPathCreate = false
)

type SMBVersion uint

const (
	SMBVersion20 SMBVersion = iota
	SMBVersion21
	SMBVersion30
)

func (obj SMBVersion) Marshal() (string, error) {
	switch obj {
	case SMBVersion20:
		return "2.0", nil
	case SMBVersion21:
		return "2.1", nil
	case SMBVersion30:
		return "3.0", nil
	default:
		return "", fmt.Errorf("unknown smb version")
	}
}

func (obj *SMBVersion) Unmarshal(s string) error {
	switch s {
	case "2.0":
		*obj = SMBVersion20
	case "2.1":
		*obj = SMBVersion21
	case "3.0":
		*obj = SMBVersion30
	default:
		return fmt.Errorf("can't unmarshal smb version %s", s)
	}

	return nil
}

func (obj *SMBVersion) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
