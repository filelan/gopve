package vm

import (
	"fmt"

	internal_types "github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types/vm/qemu"
)

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
