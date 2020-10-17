package vm

import (
	"fmt"
	"strings"

	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
)

type QEMUStorageProperties struct {
	HardDrives []QEMUHardDriveProperties
	CDROMs     []QEMUCDROMProperties
	EFIDisk    QEMUEFIDiskProperties
}

const (
	maxQEMUIDEPropertiesArrayCapacity    = 4
	maxQEMUSATAPropertiesArrayCapacity   = 6
	maxQEMUSCSIPropertiesArrayCapacity   = 31
	maxQEMUVirtIOPropertiesArrayCapacity = 16
)

func NewQEMUStorageProperties(
	props types.Properties,
) (QEMUStorageProperties, error) {
	obj := QEMUStorageProperties{}

	for _, x := range [](struct {
		Kind  qemu.Bus
		Count int
	}){
		{
			Kind:  qemu.BusIDE,
			Count: maxQEMUIDEPropertiesArrayCapacity,
		},
		{
			Kind:  qemu.BusSATA,
			Count: maxQEMUSATAPropertiesArrayCapacity,
		},
		{
			Kind:  qemu.BusSCSI,
			Count: maxQEMUSCSIPropertiesArrayCapacity,
		},
		{
			Kind:  qemu.BusVirtIO,
			Count: maxQEMUVirtIOPropertiesArrayCapacity,
		},
	} {
		for i := 0; i < x.Count; i++ {
			propName := fmt.Sprintf("%s%d", x.Kind.String(), i)
			prop, ok := props[propName]
			if !ok {
				continue
			}

			media, ok := prop.(string)
			if !ok {
				err := errors.ErrInvalidProperty
				err.AddKey("name", propName)
				err.AddKey("value", prop)
				return obj, err
			}

			if drive, err := NewQEMUStorageDrive(x.Kind, i, media); err == nil {
				switch x := drive.(type) {
				case QEMUHardDriveProperties:
					obj.HardDrives = append(obj.HardDrives, x)
				case QEMUCDROMProperties:
					obj.CDROMs = append(obj.CDROMs, x)
				default:
					panic("this should never happen")
				}
			} else {
				return obj, err
			}
		}

		prop, ok := props["efidisk0"]
		if ok {
			media, ok := prop.(string)
			if !ok {
				err := errors.ErrInvalidProperty
				err.AddKey("name", "efidisk0")
				err.AddKey("value", prop)
				return obj, err
			}

			if efiDisk, err := NewQEMUEFIDiskProperties(media); err == nil {
				obj.EFIDisk = efiDisk
			} else {
				return obj, err
			}
		}
	}

	return obj, nil
}

func NewQEMUStorageDrive(
	busKind qemu.Bus,
	busNumber int,
	media string,
) (interface{}, error) {
	props := internal_types.PVEDictionary{
		ListSeparator:     ",",
		KeyValueSeparator: "=",
		AllowNoValue:      true,
	}

	if err := (&props).Unmarshal(media); err != nil {
		return nil, err
	}

	if x, ok := props.ElemByKey("media"); ok {
		switch x.Value() {
		case "cdrom":
			return NewQEMUCDROMProperties(busKind, busNumber, props)
		default:
			return nil, fmt.Errorf("unknown media type %s", x.Value())
		}
	} else {
		return NewQEMUHardDriveProperties(busKind, busNumber, props)
	}
}

type QEMUDriveBusProperties struct {
	BusKind   qemu.Bus
	BusNumber int
}

type QEMUDriveStorageProperties struct {
	StorageName string
	StorageFile string
}

func (obj *QEMUDriveStorageProperties) setProperties(value string, prefix string) error {
	storage := internal_types.PVEList{
		Separator: ":",
	}

	if err := (&storage).Unmarshal(value); err != nil {
		return err
	} else if storage.Len() != 2 {
		err := errors.ErrInvalidProperty
		return err
	}

	obj.StorageName = storage.Elem(0)

	file := storage.Elem(1)
	if prefix != "" {
		file = strings.TrimPrefix(file, prefix)
	}

	obj.StorageFile = file

	return nil
}

type QEMUHardDriveProperties struct {
	QEMUDriveBusProperties
	QEMUDriveStorageProperties

	Size  string
	Cache qemu.HardDriveCache

	Discard    bool
	EmulateSSD bool
	IOThread   bool
	Backup     bool
	Replicate  bool

	ReadMBLimit    int
	ReadIOPSLimit  int
	ReadMBBurst    int
	ReadIOPSBurst  int
	WriteMBLimit   int
	WriteIOPSLimit int
	WriteMBBurst   int
	WriteIOPSBurst int
}

const (
	DefaultQEMUHardDriveCache qemu.HardDriveCache = qemu.HardDriveCacheNone

	DefaultQEMUHardDriveDiscard    bool = false
	DefaultQEMUHardDriveEmulateSSD bool = false
	DefaultQEMUHardDriveIOThread   bool = false
	DefaultQEMUHardDriveBackup     bool = true
	DefaultQEMUHardDriveReplicate  bool = true
)

func NewQEMUHardDriveProperties(
	busKind qemu.Bus,
	busNumber int,
	props internal_types.PVEDictionary,
) (obj QEMUHardDriveProperties, err error) {
	obj.BusKind = busKind
	obj.BusNumber = busNumber

	obj.Cache = DefaultQEMUHardDriveCache

	obj.Discard = DefaultQEMUHardDriveDiscard
	obj.EmulateSSD = DefaultQEMUHardDriveEmulateSSD
	obj.IOThread = DefaultQEMUHardDriveIOThread
	obj.Backup = DefaultQEMUHardDriveBackup
	obj.Replicate = DefaultQEMUHardDriveReplicate

	for _, kv := range props.List() {
		if !kv.HasValue() {
			if err := (&obj.QEMUDriveStorageProperties).setProperties(kv.Key(), ""); err != nil {
				return obj, err
			}
			continue
		}

		switch kv.Key() {
		case "size":
			obj.Size = kv.Value()
		case "cache":
			if err := (&obj.Cache).Unmarshal(kv.Value()); err != nil {
				return obj, err
			}
		case "discard":
			if obj.Discard, err = kv.ValueAsBool(); err != nil {
				return obj, err
			}
		case "ssd":
			if obj.EmulateSSD, err = kv.ValueAsBool(); err != nil {
				return obj, err
			}
		case "iothread":
			if obj.IOThread, err = kv.ValueAsBool(); err != nil {
				return obj, err
			}
		case "backup":
			if obj.Backup, err = kv.ValueAsBool(); err != nil {
				return obj, err
			}
		case "replicate":
			if obj.Replicate, err = kv.ValueAsBool(); err != nil {
				return obj, err
			}
		case "mbps_rd":
			if obj.ReadMBLimit, err = kv.ValueAsInt(); err != nil {
				return obj, err
			}
		case "iops_rd":
			if obj.ReadIOPSLimit, err = kv.ValueAsInt(); err != nil {
				return obj, err
			}
		case "mbps_rd_max":
			if obj.ReadMBBurst, err = kv.ValueAsInt(); err != nil {
				return obj, err
			}
		case "ios_rd_max":
			if obj.ReadIOPSBurst, err = kv.ValueAsInt(); err != nil {
				return obj, err
			}
		case "mbps_wr":
			if obj.WriteMBLimit, err = kv.ValueAsInt(); err != nil {
				return obj, err
			}
		case "iops_wr":
			if obj.WriteIOPSLimit, err = kv.ValueAsInt(); err != nil {
				return obj, err
			}
		case "mbps_wr_max":
			if obj.WriteMBBurst, err = kv.ValueAsInt(); err != nil {
				return obj, err
			}
		case "ios_wr_max":
			if obj.WriteIOPSBurst, err = kv.ValueAsInt(); err != nil {
				return obj, err
			}
		default:
			err := errors.ErrInvalidProperty
			err.AddKey("name", fmt.Sprintf("%s%d", busKind.String(), busNumber))
			return obj, err
		}
	}

	return obj, nil
}

type QEMUCDROMProperties struct {
	QEMUDriveBusProperties
	QEMUDriveStorageProperties

	Source QEMUCDROMSource
	Size   string
}

type QEMUCDROMSource int

const (
	QEMUCDROMSourceNone QEMUCDROMSource = iota
	QEMUCDROMSourcePhysical
	QEMUCDROMSourceISOFile
)

func NewQEMUCDROMProperties(
	busKind qemu.Bus,
	busNumber int,
	props internal_types.PVEDictionary,
) (obj QEMUCDROMProperties, err error) {
	obj.BusKind = busKind
	obj.BusNumber = busNumber

	for _, kv := range props.List() {
		if !kv.HasValue() {
			switch kv.Key() {
			case "none":
				obj.Source = QEMUCDROMSourceNone
			case "cdrom":
				obj.Source = QEMUCDROMSourcePhysical
			default:
				obj.Source = QEMUCDROMSourceISOFile
				if err := (&obj.QEMUDriveStorageProperties).setProperties(kv.Key(), "iso/"); err != nil {
					return obj, err
				}
			}

			continue
		}

		switch kv.Key() {
		case "media":
			continue
		case "size":
			obj.Size = kv.Value()
		default:
			err := errors.ErrInvalidProperty
			err.AddKey("name", fmt.Sprintf("%s%d", busKind.String(), busNumber))
			return obj, err
		}
	}

	return obj, nil
}

type QEMUEFIDiskProperties struct {
	QEMUDriveStorageProperties

	Size string
}

func NewQEMUEFIDiskProperties(
	media string,
) (obj QEMUEFIDiskProperties, err error) {
	props := internal_types.PVEDictionary{
		ListSeparator:     ",",
		KeyValueSeparator: "=",
		AllowNoValue:      true,
	}

	if err := (&props).Unmarshal(media); err != nil {
		return obj, err
	}

	for _, kv := range props.List() {
		if !kv.HasValue() {
			if err := (&obj.QEMUDriveStorageProperties).setProperties(kv.Key(), ""); err != nil {
				return obj, err
			}

			continue
		}

		switch kv.Key() {
		case "size":
			obj.Size = kv.Value()
		default:
			err := errors.ErrInvalidProperty
			err.AddKey("name", "efidisk0")
			return obj, err
		}
	}

	return obj, nil
}
