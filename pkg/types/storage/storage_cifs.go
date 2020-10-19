package storage

import (
	"encoding/json"

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

	return obj, errors.ChainUntilFail(
		func() error {
			return props.SetRequiredString(
				mkCIFSServer,
				&obj.Server,
				nil,
			)
		},
		func() error {
			return props.SetFixedValue(
				mkCIFSSMBVersion,
				&obj.SMBVersion,
				DefaultStorageCIFSSMBVersion,
				nil,
			)
		},
		func() error {
			return props.SetString(
				mkCIFSDomain,
				&obj.Domain,
				DefaultStorageCIFSDomain,
				nil,
			)
		},
		func() error {
			return props.SetString(
				mkCIFSUsername,
				&obj.Username,
				DefaultStorageCIFSUsername,
				nil,
			)
		},
		func() error {
			return props.SetString(
				mkCIFSPassword,
				&obj.Password,
				DefaultStorageCIFSPassword,
				nil,
			)
		},
		func() error {
			return props.SetRequiredString(
				mkCIFSServerShare,
				&obj.ServerShare,
				nil,
			)
		},
		func() error {
			return props.SetRequiredString(
				mkCIFSLocalPath,
				&obj.LocalPath,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkCIFSLocalPathCreate,
				&obj.LocalPathCreate,
				DefaultStorageCIFSLocalPathCreate,
				nil,
			)
		},
	)
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
