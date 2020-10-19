package storage

import (
	"encoding/json"

	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
	"github.com/xabinapal/gopve/pkg/types/schema"
)

type StorageZFSOverISCSI interface {
	Storage

	Portal() string
	Target() string

	PoolName() string

	BlockSize() string
	UseSparse() bool
	WriteCache() bool

	ISCSIProvider() ISCSIProvider

	COMSTARHostGroup() string
	COMSTARTargetGroup() string
	LIOTargetPortalGroup() string
}

const (
	StorageZFSOverISCSIKernelContents    = ContentQEMUData
	StorageZFSOverISCSIKernelImageFormat = ImageFormatRaw
	StorageZFSOverISCSIKernelShared      = AllowShareForced
	StorageZFSOverISCSIKernelSnapshots   = AllowSnapshotAll
	StorageZFSOverISCSIKernelClones      = AllowCloneAll
)

type StorageZFSOverISCSIProperties struct {
	Portal string
	Target string

	PoolName string

	BlockSize  string
	UseSparse  bool
	WriteCache bool

	ISCSIProvider ISCSIProvider

	COMSTARHostGroup     string
	COMSTARTargetGroup   string
	LIOTargetPortalGroup string
}

const (
	mkZFSOverISCSIPortal                      = "portal"
	mkZFSOverISCSITarget                      = "target"
	mkZFSOverISCSIPoolName                    = "pool"
	mkZFSOverISCSIBlockSize                   = "blocksize"
	mkZFSOverISCSIUseSparse                   = "sparse"
	mkZFSOverISCSIWriteCache                  = "nowritecache"
	mkZFSOverISCSIISCSIProvider               = "iscsiprovider"
	mkZFSOverISCSICOMSTARHostGroup            = "comstar_hg"
	mkZFSOverISCSICOMSTARTargetGroup          = "comstar_tg"
	mkZFSOverISCSICOMSTARLIOTargetPortalGroup = "lio_tpg"
)

const (
	DefaultStorageZFSOverISCSIUseSparse            = false
	DefaultStorageZFSOverISCSIWriteCache           = true
	DefaultStorageZFSOverISCSICOMSTARHostGroup     = ""
	DefaultStorageZFSOverISCSICOMSTARTargetGroup   = ""
	DefaultStorageZFSOverISCSILIOTargetPortalGroup = ""
)

func NewStorageZFSOverISCSIProperties(
	props types.Properties,
) (*StorageZFSOverISCSIProperties, error) {
	obj := new(StorageZFSOverISCSIProperties)

	return obj, errors.ChainUntilFail(
		func() error {
			return props.SetRequiredString(
				mkZFSOverISCSIPortal,
				&obj.Portal,
				nil,
			)
		},
		func() error {
			return props.SetRequiredString(
				mkZFSOverISCSITarget,
				&obj.Target,
				nil,
			)
		},
		func() error {
			return props.SetRequiredString(
				mkZFSOverISCSIPoolName,
				&obj.PoolName,
				nil,
			)
		},
		func() error {
			return props.SetRequiredString(
				mkZFSOverISCSIBlockSize,
				&obj.BlockSize,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkZFSOverISCSIUseSparse,
				&obj.UseSparse,
				DefaultStorageZFSOverISCSIUseSparse,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkZFSOverISCSIWriteCache,
				&obj.WriteCache,
				DefaultStorageZFSOverISCSIWriteCache,
				&schema.BoolFunctions{
					TransformFunc: func(val bool) bool {
						return !val
					},
				},
			)
		},
		func() error {
			return props.SetRequiredFixedValue(
				mkZFSOverISCSIISCSIProvider,
				&obj.ISCSIProvider,
				nil,
			)
		},
		func() error {
			return props.SetString(
				mkZFSOverISCSICOMSTARHostGroup,
				&obj.COMSTARHostGroup,
				DefaultStorageZFSOverISCSICOMSTARHostGroup,
				nil,
			)
		},
		func() error {
			return props.SetString(
				mkZFSOverISCSICOMSTARTargetGroup,
				&obj.COMSTARTargetGroup,
				DefaultStorageZFSOverISCSICOMSTARTargetGroup,
				nil,
			)
		},
		func() error {
			return props.SetString(
				mkZFSOverISCSICOMSTARLIOTargetPortalGroup,
				&obj.LIOTargetPortalGroup,
				DefaultStorageZFSOverISCSILIOTargetPortalGroup,
				nil,
			)
		},
	)
}

type ISCSIProvider string

const (
	ISCSIProviderCOMSTAR ISCSIProvider = "comstar"
	ISCSIProviderISTGT   ISCSIProvider = "istgt"
	ISCSIProviderIET     ISCSIProvider = "iet"
	ISCSIProviderLIO     ISCSIProvider = "LIO"
)

func (obj ISCSIProvider) IsValid() bool {
	switch obj {
	case ISCSIProviderCOMSTAR,
		ISCSIProviderISTGT,
		ISCSIProviderIET,
		ISCSIProviderLIO:
		return true
	default:
		return false
	}
}

func (obj ISCSIProvider) IsUnknown() bool {
	return !obj.IsValid()
}

func (obj ISCSIProvider) Marshal() (string, error) {
	return string(obj), nil
}

func (obj *ISCSIProvider) Unmarshal(s string) error {
	*obj = ISCSIProvider(s)
	return nil
}

func (obj *ISCSIProvider) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	return obj.Unmarshal(s)
}
