package vm

import (
	"fmt"

	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
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
}

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
	)

	return obj, err
}

type QEMUGlobalProperties struct {
	OSType QEMUOSType

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
	Kind  QEMUCPUKind
	Flags []QEMUCPUFlags

	Architecture QEMUCPUArchitecture

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

	DefaultQEMUCPUPropertyKind         QEMUCPUKind         = QEMUCPUKindKVM64
	DefaultQEMUCPUPropertyArchitecture QEMUCPUArchitecture = QEMUCPUArchitectureHost

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

	obj.Flags = []QEMUCPUFlags{}
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
						var flag QEMUCPUFlags
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

	if err := props.SetRequiredUint(mkQEMUCPUPropertyVCPUs, &obj.VCPUs, &types.PropertyUintFunctions{
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
}

const (
	maxQEMUIDEPropertiesArrayCapacity    = 4
	maxQEMUSATAPropertiesArrayCapacity   = 6
	maxQEMUSCSIPropertiesArrayCapacity   = 31
	maxQEMUVirtIOPropertiesArrayCapacity = 16
)

func NewQEMUStorageProperties(props types.Properties) (QEMUStorageProperties, error) {
	obj := QEMUStorageProperties{}

	for i := 0; i < maxQEMUIDEPropertiesArrayCapacity; i++ {
		propName := fmt.Sprintf("ide%d", i)
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

		if drive, err := NewQEMUStorageDrive(media); err == nil {
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

	return obj, nil
}

func NewQEMUStorageDrive(media string) (interface{}, error) {
	dict := internal_types.PVEDictionary{
		ListSeparator:     ",",
		KeyValueSeparator: "=",
		AllowNoValue:      true,
	}

	if err := (&dict).Unmarshal(media); err != nil {
		return nil, err
	}

	if x, ok := dict.ElemByKey("media"); ok {
		switch x.Key() {
		case "cdrom":
			return NewQEMUCDROMProperties(dict)
		default:
			return nil, fmt.Errorf("unknown media type %s", x.Key())
		}
	} else {
		return NewQEMUHardDriveProperties(dict)
	}
}

type QEMUHardDriveProperties struct {
	Size           string
	Cache          string
	Discard        bool
	EmulateSSD     bool
	IOThread       bool
	Backup         bool
	Replicate      bool
	ReadMBLimit    int
	ReadIOPSLimit  int
	ReadMBBurst    int
	ReadIOPSBurst  int
	WriteMBLimit   int
	WriteIOPSLimit int
	WriteMBBurst   int
	WriteIOPSBurst int
}

func NewQEMUHardDriveProperties(dict internal_types.PVEDictionary) (obj QEMUHardDriveProperties, err error) {
	for _, kv := range dict.List() {
		switch kv.Key() {
		case "size":
			obj.Size = kv.Value()
		case "cache":
			obj.Cache = kv.Value()
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
		}
	}

	return obj, nil
}

type QEMUCDROMProperties struct {
}

func NewQEMUCDROMProperties(dict internal_types.PVEDictionary) (QEMUCDROMProperties, error) {
	obj := QEMUCDROMProperties{}

	return obj, nil
}
