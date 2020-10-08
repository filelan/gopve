package storage

import (
	"encoding/json"

	"github.com/xabinapal/gopve/pkg/types"
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

const (
	StorageCIFSContents    = ContentQEMUData & ContentContainerData & ContentContainerTemplate & ContentISO & ContentBackup & ContentSnippet
	StorageCIFSImageFormat = ImageFormatRaw & ImageFormatQcow2 & ImageFormatVMDK
	StorageCIFSShared      = AllowShareForced
	StorageCIFSSnapshots   = AllowSnapshotQcow2
	StorageCIFSClones      = AllowCloneQcow2
)

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

const (
	mkCIFSServer          = "server"
	mkCIFSSMBVersion      = "smbversion"
	mkCIFSDomain          = "domain"
	mkCIFSUsername        = "username"
	mkCIFSPassword        = "password"
	mkCIFSServerShare     = "share"
	mkCIFSLocalPath       = "path"
	mkCIFSLocalPathCreate = "mkdir"
)

const (
	DefaultStorageCIFSSMBVersion      = SMBVersion30
	DefaultStorageCIFSDomain          = ""
	DefaultStorageCIFSUsername        = ""
	DefaultStorageCIFSPassword        = ""
	DefaultStorageCIFSLocalPathCreate = false
)

func NewStorageCIFSProperties(
	props types.Properties,
) (*StorageCIFSProperties, error) {
	obj := new(StorageCIFSProperties)

	if err := props.SetRequiredString(mkCIFSServer, &obj.Server, nil); err != nil {
		return nil, err
	}

	if err := props.SetFixedValue(mkCIFSSMBVersion, &obj.SMBVersion, DefaultStorageCIFSSMBVersion, nil); err != nil {
		return nil, err
	}

	if err := props.SetString(mkCIFSDomain, &obj.Domain, DefaultStorageCIFSDomain, nil); err != nil {
		return nil, err
	}

	if err := props.SetString(mkCIFSUsername, &obj.Username, DefaultStorageCIFSUsername, nil); err != nil {
		return nil, err
	}

	if err := props.SetString(mkCIFSPassword, &obj.Password, DefaultStorageCIFSPassword, nil); err != nil {
		return nil, err
	}

	if err := props.SetRequiredString(mkCIFSServerShare, &obj.ServerShare, nil); err != nil {
		return nil, err
	}

	if err := props.SetRequiredString(mkCIFSLocalPath, &obj.LocalPath, nil); err != nil {
		return nil, err
	}

	if err := props.SetBool(mkCIFSLocalPathCreate, &obj.LocalPathCreate, DefaultStorageCIFSLocalPathCreate, nil); err != nil {
		return nil, err
	}

	return obj, nil
}

type SMBVersion string

const (
	SMBVersion20 SMBVersion = "2.0"
	SMBVersion21 SMBVersion = "2.1"
	SMBVersion30 SMBVersion = "3.0"
)

func (obj SMBVersion) IsValid() bool {
	switch obj {
	case SMBVersion20, SMBVersion21, SMBVersion30:
		return true
	default:
		return false
	}
}

func (obj SMBVersion) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj SMBVersion) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *SMBVersion) Unmarshal(s string) error {
	*obj = SMBVersion(s)
	return nil
}

func (obj *SMBVersion) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
