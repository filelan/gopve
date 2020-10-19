package qemu

import (
	"fmt"
	"strings"

	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
)

type StorageProperties struct {
	HardDrives []HardDriveProperties
	CDROMs     []CDROMProperties
	EFIDisk    EFIDiskProperties
}

const (
	maxIDEPropertiesArrayCapacity    = 4
	maxSATAPropertiesArrayCapacity   = 6
	maxSCSIPropertiesArrayCapacity   = 31
	maxVirtIOPropertiesArrayCapacity = 16
)

func NewStorageProperties(
	props types.Properties,
) (StorageProperties, error) {
	obj := StorageProperties{}

	for _, x := range [](struct {
		Kind  Bus
		Count int
	}){
		{
			Kind:  BusIDE,
			Count: maxIDEPropertiesArrayCapacity,
		},
		{
			Kind:  BusSATA,
			Count: maxSATAPropertiesArrayCapacity,
		},
		{
			Kind:  BusSCSI,
			Count: maxSCSIPropertiesArrayCapacity,
		},
		{
			Kind:  BusVirtIO,
			Count: maxVirtIOPropertiesArrayCapacity,
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

			if drive, err := NewStorageDrive(x.Kind, i, media); err == nil {
				switch x := drive.(type) {
				case HardDriveProperties:
					obj.HardDrives = append(obj.HardDrives, x)
				case CDROMProperties:
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

			if efiDisk, err := NewEFIDiskProperties(media); err == nil {
				obj.EFIDisk = efiDisk
			} else {
				return obj, err
			}
		}
	}

	return obj, nil
}

func NewStorageDrive(
	deviceBus Bus,
	deviceNumber int,
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

	if x, ok := props.Elem("media"); ok {
		switch x {
		case "cdrom":
			return NewCDROMProperties(deviceBus, deviceNumber, props)
		default:
			return nil, fmt.Errorf("unknown media type %s", x)
		}
	} else {
		return NewHardDriveProperties(deviceBus, deviceNumber, props)
	}
}

type DriveStorageProperties struct {
	StorageName string
	StorageFile string
}

func (obj *DriveStorageProperties) setProperties(
	value string,
	prefix string,
) error {
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

type HardDriveProperties struct {
	DeviceBus    Bus
	DeviceNumber int

	DriveStorageProperties

	Size  string
	Cache HardDriveCache

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
	DefaultHardDriveCache HardDriveCache = HardDriveCacheNone

	DefaultHardDriveDiscard    bool = false
	DefaultHardDriveEmulateSSD bool = false
	DefaultHardDriveIOThread   bool = false
	DefaultHardDriveBackup     bool = true
	DefaultHardDriveReplicate  bool = true
)

func NewHardDriveProperties(
	deviceBus Bus,
	deviceNumber int,
	props internal_types.PVEDictionary,
) (obj HardDriveProperties, err error) {
	obj.DeviceBus = deviceBus
	obj.DeviceNumber = deviceNumber

	obj.Cache = DefaultHardDriveCache

	obj.Discard = DefaultHardDriveDiscard
	obj.EmulateSSD = DefaultHardDriveEmulateSSD
	obj.IOThread = DefaultHardDriveIOThread
	obj.Backup = DefaultHardDriveBackup
	obj.Replicate = DefaultHardDriveReplicate

	for _, kv := range props.List() {
		if !kv.HasValue() {
			if err := (&obj.DriveStorageProperties).setProperties(kv.Key(), ""); err != nil {
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
			return obj, err
		}
	}

	return obj, nil
}

type CDROMProperties struct {
	DeviceBus    Bus
	DeviceNumber int

	DriveStorageProperties

	Source CDROMSource
	Size   string
}

type CDROMSource int

const (
	CDROMSourceNone CDROMSource = iota
	CDROMSourcePhysical
	CDROMSourceISOFile
)

func NewCDROMProperties(
	deviceBus Bus,
	deviceNumber int,
	props internal_types.PVEDictionary,
) (obj CDROMProperties, err error) {
	obj.DeviceBus = deviceBus
	obj.DeviceNumber = deviceNumber

	for _, kv := range props.List() {
		if !kv.HasValue() {
			switch kv.Key() {
			case "none":
				obj.Source = CDROMSourceNone
			case "cdrom":
				obj.Source = CDROMSourcePhysical
			default:
				obj.Source = CDROMSourceISOFile
				if err := (&obj.DriveStorageProperties).setProperties(kv.Key(), "iso/"); err != nil {
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
			return obj, err
		}
	}

	return obj, nil
}

type EFIDiskProperties struct {
	DriveStorageProperties

	Size string
}

func NewEFIDiskProperties(
	media string,
) (obj EFIDiskProperties, err error) {
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
			if err := (&obj.DriveStorageProperties).setProperties(kv.Key(), ""); err != nil {
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
