package vm

import (
	"fmt"
	"strings"

	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
)

type QEMUVirtualMachine interface {
	VirtualMachine

	CPU() (QEMUCPUProperties, error)
	Memory() (QEMUMemoryProperties, error)

	GetQEMUProperties() (QEMUProperties, error)
	SetQEMUProperties(props QEMUProperties) error
}

type QEMUCreateOptions struct {
	VMID uint
	Node string

	Properties QEMUProperties
}

func (obj QEMUCreateOptions) MapToValues() (request.Values, error) {
	values, err := obj.Properties.MapToValues()
	if err != nil {
		return nil, err
	}

	return values, nil
}

type QEMUProperties struct {
	QEMUGlobalProperties
	CPU     QEMUCPUProperties
	Memory  QEMUMemoryProperties
	Storage QEMUStorageProperties
	Network []QEMUNetworkInterfaceProperties
}

const (
	maxQEMUNetworkInterfacePropertiesArrayCapacity = 32
)

func NewQEMUProperties(props types.Properties) (QEMUProperties, error) {
	obj := QEMUProperties{}

	err := errors.ChainUntilFail(
		func() (err error) {
			obj.QEMUGlobalProperties, err = NewQEMUGlobalProperties(props)
			return err
		},
		func() (err error) {
			obj.CPU, err = NewQEMUCPUProperties(props)
			return err
		},
		func() (err error) {
			obj.Memory, err = NewQEMUMemoryProperties(props)
			return err
		},
		func() (err error) {
			obj.Storage, err = NewQEMUStorageProperties(props)
			return err
		},
		func() (err error) {
			for i := 0; i < maxQEMUNetworkInterfacePropertiesArrayCapacity; i++ {
				propName := fmt.Sprintf("net%d", i)
				prop, ok := props[propName]
				if !ok {
					continue
				}

				x, ok := prop.(string)
				if !ok {
					err := errors.ErrInvalidProperty
					err.AddKey("name", propName)
					err.AddKey("value", prop)
					return err
				}

				if network, err := NewQEMUNetworkInterfaceProperties(x); err == nil {
					obj.Network = append(obj.Network, network)
				} else {
					return err
				}
			}

			return nil
		},
	)

	return obj, err
}

type QEMUGlobalProperties struct {
	OSType qemu.OSType

	Protected bool

	StartOnBoot bool
}

const (
	mkQEMUGlobalPropertyOSType      = "ostype"
	mkQEMUGlobalPropertyProtected   = "protection"
	mkQEMUGlobalPropertyStartOnBoot = "onboot"

	DefaultQEMUGlobalPropertyProtected   bool = false
	DefaultQEMUGlobalPropertyStartOnBoot bool = false
)

func NewQEMUGlobalProperties(
	props types.Properties,
) (QEMUGlobalProperties, error) {
	obj := QEMUGlobalProperties{}

	err := errors.ChainUntilFail(
		func() error {
			return props.SetRequiredFixedValue(
				mkQEMUGlobalPropertyOSType,
				&obj.OSType,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkQEMUGlobalPropertyProtected,
				&obj.Protected,
				DefaultQEMUGlobalPropertyProtected,
				nil,
			)
		},
		func() error {
			return props.SetBool(
				mkQEMUGlobalPropertyStartOnBoot,
				&obj.StartOnBoot,
				DefaultQEMUGlobalPropertyStartOnBoot,
				nil,
			)
		},
	)

	return obj, err
}

func (obj QEMUProperties) MapToValues() (request.Values, error) {
	values := request.Values{}

	cpuValues, err := obj.CPU.MapToValues()
	if err != nil {
		return nil, err
	} else {
		for k, v := range cpuValues {
			values[k] = v
		}
	}

	memoryValues, err := obj.Memory.MapToValues()
	if err != nil {
		return nil, err
	} else {
		for k, v := range memoryValues {
			values[k] = v
		}
	}

	return values, nil
}

type QEMUCPUProperties struct {
	Kind  qemu.CPUType
	Flags []qemu.CPUFlags

	Architecture qemu.CPUArchitecture

	Sockets uint
	Cores   uint
	VCPUs   uint

	Limit uint
	Units uint

	NUMA bool

	FreezeAtStartup bool
}

const (
	mkQEMUCPUPropertyCPU   = "cpu"
	mkQEMUCPUPropertyFlags = "flags"

	mkQEMUCPUPropertyArchitecture = "arch"

	mkQEMUCPUPropertySockets = "sockets"
	mkQEMUCPUPropertyCores   = "cores"
	mkQEMUCPUPropertyVCPUs   = "vcpus"
	mkQEMUCPUPropertyLimit   = "cpulimit"
	mkQEMUCPUPropertyUnits   = "cpuunits"

	mkQEMUCPUPropertyNUMA            = "numa"
	mkQEMUCPUPropertyFreezeAtStartup = "freeze"

	DefaultQEMUCPUPropertyKind         qemu.CPUType         = qemu.CPUTypeKVM64
	DefaultQEMUCPUPropertyArchitecture qemu.CPUArchitecture = qemu.CPUArchitectureHost

	DefaultQEMUCPUPropertyLimit uint = 0
	DefaultQEMUCPUPropertyUnits uint = 1024

	DefaultQEMUCPUPropertyNUMA            bool = false
	DefaultQEMUCPUPropertyFreezeAtStartup bool = false
)

func NewQEMUCPUProperties(props types.Properties) (QEMUCPUProperties, error) {
	obj := QEMUCPUProperties{}

	if err := props.SetFixedValue(mkQEMUCPUPropertyCPU, &obj.Kind, DefaultQEMUCPUPropertyKind, nil); err != nil {
		return obj, err
	}

	obj.Flags = []qemu.CPUFlags{}
	if v, ok := props[mkQEMUCPUPropertyCPU].(string); ok {
		cpuOptions := internal_types.PVEDictionary{
			ListSeparator:     ",",
			KeyValueSeparator: "=",
			AllowNoValue:      true,
		}

		if err := (&cpuOptions).Unmarshal(v); err != nil {
			err := errors.ErrInvalidProperty
			err.AddKey("name", mkQEMUCPUPropertyCPU)
			err.AddKey("value", v)
			return obj, err
		}

		for _, vv := range cpuOptions.List() {
			if !vv.HasValue() {
				if err := (&obj.Kind).Unmarshal(vv.Key()); err != nil {
					err := errors.ErrInvalidProperty
					err.AddKey("name", mkQEMUCPUPropertyCPU)
					err.AddKey("value", v)
					return obj, err
				}
			} else {
				switch vv.Key() {
				case mkQEMUCPUPropertyFlags:
					flags := internal_types.PVEList{
						Separator: ";",
					}

					if err := (&flags).Unmarshal(vv.Value()); err != nil {
						err := errors.ErrInvalidProperty
						err.AddKey("name", mkQEMUCPUPropertyCPU)
						err.AddKey("value", v)
						return obj, err
					}

					for _, v := range flags.List() {
						var flag qemu.CPUFlags
						if err := (&flag).Unmarshal(v); err != nil {
							err := errors.ErrInvalidProperty
							err.AddKey("name", mkQEMUCPUPropertyCPU)
							err.AddKey("value", v)
							return obj, err
						}

						obj.Flags = append(obj.Flags, flag)
					}
				}
			}
		}
	} else {
		obj.Kind = DefaultQEMUCPUPropertyKind
	}

	if err := props.SetFixedValue(mkQEMUCPUPropertyArchitecture, &obj.Architecture, DefaultQEMUCPUPropertyArchitecture, nil); err != nil {
		return obj, err
	}

	if err := props.SetRequiredUint(mkQEMUCPUPropertySockets, &obj.Sockets, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= 4
		},
	}); err != nil {
		return obj, err
	}

	if err := props.SetRequiredUint(mkQEMUCPUPropertyCores, &obj.Cores, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= 128
		},
	}); err != nil {
		return obj, err
	}

	if err := props.SetUint(mkQEMUCPUPropertyVCPUs, &obj.VCPUs, obj.Sockets*obj.Cores, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= obj.Sockets*obj.Cores
		},
	}); err != nil {
		return obj, err
	}

	if err := props.SetUintFromString(mkQEMUCPUPropertyLimit, &obj.Limit, DefaultQEMUCPUPropertyLimit, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= 128
		},
	}); err != nil {
		return obj, err
	}

	if err := props.SetUint(mkQEMUCPUPropertyUnits, &obj.Units, DefaultQEMUCPUPropertyUnits, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val >= 8 && val <= 500000
		},
	}); err != nil {
		return obj, err
	}

	if err := props.SetBool(mkQEMUCPUPropertyNUMA, &obj.NUMA, DefaultQEMUCPUPropertyNUMA, nil); err != nil {
		return obj, err
	}

	if err := props.SetBool(mkQEMUCPUPropertyFreezeAtStartup, &obj.FreezeAtStartup, DefaultQEMUCPUPropertyFreezeAtStartup, nil); err != nil {
		return obj, err
	}

	return obj, nil
}

func (obj QEMUCPUProperties) MapToValues() (request.Values, error) {
	values := request.Values{}

	sockets := obj.Sockets
	if sockets == 0 {
		sockets = 1
	} else if sockets > 4 {
		return nil, fmt.Errorf("Invalid CPU sockets, the maximum allowed is 4")
	}
	values.AddUint("sockets", sockets)

	cores := obj.Cores
	if cores == 0 {
		cores = 1
	} else if cores > 128 {
		return nil, fmt.Errorf("Invalid CPU cores, the maximum allowed is 128")
	}
	values.AddUint("cores", cores)

	if obj.VCPUs != 0 && (obj.VCPUs > sockets*cores) {
		return nil, fmt.Errorf(
			"Invalid CPU hotplugged cores, can't be greater than sockets * cores",
		)
	} else if obj.VCPUs != 0 {
		values.AddUint("vcpus", obj.VCPUs)
	}

	if obj.Limit > 128 {
		return nil, fmt.Errorf("Invalid CPU limit, must be between 0 and 128")
	} else if obj.Limit != 0 {
		values.AddUint("cpulimit", obj.Limit)
	}

	if obj.Units != 0 && (obj.Units < 2 || obj.Units > 262144) {
		return nil, fmt.Errorf(
			"Invalid CPU units, must be between 2 and 262144",
		)
	} else if obj.Units != 0 && obj.Units != 1024 {
		values.AddUint("cpuunits", obj.Units)
	}

	values.AddBool("numa", obj.NUMA)

	values.AddBool("freeze", obj.FreezeAtStartup)

	return values, nil
}

type QEMUMemoryProperties struct {
	Memory uint

	Ballooning    bool
	MinimumMemory uint
	Shares        uint
}

const (
	mkQEMUMemoryPropertyMemory  = "memory"
	mkQEMUMemoryPropertyBalloon = "balloon"
	mkQEMUMemoryPropertyShares  = "shares"

	DefaultQEMUMemoryShares uint = 1000
)

func NewQEMUMemoryProperties(
	props types.Properties,
) (QEMUMemoryProperties, error) {
	obj := QEMUMemoryProperties{}

	if err := props.SetRequiredUint(mkQEMUMemoryPropertyMemory, &obj.Memory, &types.PropertyUintFunctions{
		ValidateFunc: func(val uint) bool {
			return val <= 4178944
		},
	}); err != nil {
		return obj, err
	}

	if v, ok := props[mkQEMUMemoryPropertyBalloon].(float64); ok {
		if v != float64(int(v)) || v < 0 || uint(v) > obj.Memory {
			err := errors.ErrInvalidProperty
			err.AddKey("name", mkQEMUMemoryPropertyBalloon)
			err.AddKey("value", v)
			return obj, err
		} else if v == 0 {
			obj.Ballooning = false
			obj.MinimumMemory = obj.Memory
			obj.Shares = 0
		} else if uint(v) == obj.Memory {
			obj.Ballooning = true
			obj.MinimumMemory = obj.Memory
			obj.Shares = 0
		} else {
			obj.Ballooning = true
			obj.MinimumMemory = uint(v)

			if v, ok := props[mkQEMUMemoryPropertyShares].(float64); ok {
				if v != float64(int(v)) || v < 0 || v > 50000 {
					err := errors.ErrInvalidProperty
					err.AddKey("name", mkQEMUMemoryPropertyShares)
					err.AddKey("value", v)
					return obj, err
				}

				obj.Shares = uint(v)
			} else {
				obj.Shares = DefaultQEMUMemoryShares
			}
		}
	} else {
		obj.Ballooning = true
		obj.MinimumMemory = obj.Memory
		obj.Shares = 0
	}

	return obj, nil
}

func (obj QEMUMemoryProperties) MapToValues() (request.Values, error) {
	values := request.Values{}

	memory := obj.Memory
	if memory == 0 {
		memory = 512
	} else if memory < 16 || memory > 4178944 {
		return nil, fmt.Errorf("Invalid memory, must be between 16 and 4178944")
	}
	values.AddUint("memory", memory)

	if obj.Ballooning {
		minimumMemory := obj.MinimumMemory
		if minimumMemory == 0 {
			minimumMemory = memory
		} else if minimumMemory > memory {
			return nil, fmt.Errorf("Invalid memory ballooning minimum, can't be greater than total memory")
		}

		values.AddUint("balloon", minimumMemory)

		if minimumMemory == memory {
			values.AddUint("shares", 0)
		} else {
			if obj.Shares == 0 {
				values.AddUint("shares", 1000)
			} else {
				values.AddUint("shares", obj.Shares)
			}
		}
	} else {
		values.AddUint("balloon", 0)
		values.AddUint("shares", 0)
	}

	return values, nil
}

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
			err.AddKey("name", fmt.Sprintf("%s%d", "efidisk0"))
			return obj, err
		}
	}

	return obj, nil
}

type QEMUNetworkInterfaceProperties struct {
	Model      qemu.NetworkModel
	MACAddress string

	Bridge string
	VLAN   int

	Enabled        bool
	EnableFirewall bool

	RateLimitMBps int
	Multiqueue    int
}

func NewQEMUNetworkInterfaceProperties(
	media string,
) (obj QEMUNetworkInterfaceProperties, err error) {
	props := internal_types.PVEDictionary{
		ListSeparator:     ",",
		KeyValueSeparator: "=",
		AllowNoValue:      true,
	}

	if err := (&props).Unmarshal(media); err != nil {
		return obj, err
	}

	for _, kv := range props.List() {
		switch kv.Key() {
		case "e1000":
			obj.Model = qemu.NetworkModelIntelE1000
			obj.MACAddress = kv.Value()
		case "virtio":
			obj.Model = qemu.NetworkModelVirtIO
			obj.MACAddress = kv.Value()
		case "rtl8139":
			obj.Model = qemu.NetworkModelRealtekRTL8139
			obj.MACAddress = kv.Value()
		case "vmxnet3":
			obj.Model = qemu.NetworkModelVMwareVMXNET3
			obj.MACAddress = kv.Value()
		case "bridge":
			obj.Bridge = kv.Value()
		case "tag":
			if obj.VLAN, err = kv.ValueAsInt(); err != nil {
				return obj, err
			}
		case "link_down":
			if obj.Enabled, err = kv.ValueAsBool(); err != nil {
				return obj, err
			}
		case "firewall":
			if obj.EnableFirewall, err = kv.ValueAsBool(); err != nil {
				return obj, err
			}
		case "rate":
			if obj.RateLimitMBps, err = kv.ValueAsInt(); err != nil {
				return obj, err
			}
		case "queues":
			if obj.Multiqueue, err = kv.ValueAsInt(); err != nil {
				return obj, err
			}
		default:
			return obj, fmt.Errorf("unknown property %s", kv.Key())
		}
	}

	return obj, nil
}
