package firewall

import (
	"strings"

	"github.com/xabinapal/gopve/pkg/request"
)

type Rule struct {
	Enable      bool
	Description string

	SecurityGroup string
	Interface     string

	Direction          Direction
	Action             Action
	SourceAddress      string
	DestinationAddress string

	Macro    Macro
	Protocol Protocol

	SourcePorts      PortRanges
	DestinationPorts PortRanges

	LogLevel LogLevel

	Digest string
}

func (obj Rule) MapToValues(update bool) (request.Values, error) {
	var delete []string
	values := make(request.Values)

	values.AddBool("enable", obj.Enable)
	values.ConditionalAddString("comment", obj.Description, obj.Description != "")

	if obj.SecurityGroup == "" {
		values.AddObject("type", obj.Direction)
		values.AddObject("action", obj.Action)

		if obj.SourceAddress == "" {
			delete = append(delete, "source")
		} else {
			values.AddString("source", obj.SourceAddress)
		}
		if obj.DestinationAddress == "" {
			delete = append(delete, "dest")
		} else {
			values.AddString("dest", obj.DestinationAddress)
		}

		if obj.Macro == MacroNone {
			delete = append(delete, "macro")

			if obj.Protocol == ProtocolNone {
				delete = append(delete, "proto")
			} else {
				values.ConditionalAddObject("proto", obj.Protocol, obj.Protocol != ProtocolNone)
			}
			if len(obj.SourcePorts) == 0 {
				delete = append(delete, "sport")
			} else {
				values.ConditionalAddObject("sport", obj.SourcePorts, len(obj.SourcePorts) != 0)
			}
			if len(obj.DestinationPorts) == 0 {
				delete = append(delete, "dport")
			} else {
				values.ConditionalAddObject("dport", obj.DestinationPorts, len(obj.DestinationPorts) != 0)
			}
		} else {
			delete = append(delete, "proto")
			delete = append(delete, "sport")
			delete = append(delete, "dport")

			values.AddObject("macro", obj.Macro)
		}

		values.AddObject("log", obj.LogLevel)
	} else {
		values.AddString("type", "group")
		values.AddString("action", obj.SecurityGroup)
	}

	values.ConditionalAddString("iface", obj.Interface, obj.Interface != "")

	values.ConditionalAddString("digest", obj.Digest, obj.Digest != "")

	if update {
		values.ConditionalAddString("delete", strings.Join(delete, ","), len(delete) != 0)
	}

	return values, nil
}
