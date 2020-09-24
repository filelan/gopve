package firewall

import (
	"github.com/xabinapal/gopve/internal/types"
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
	delete := types.PVEStringList{Separator: ","}
	values := make(request.Values)

	values.AddBool("enable", obj.Enable)
	values.ConditionalAddString(
		"comment",
		obj.Description,
		obj.Description != "",
	)

	if obj.SecurityGroup == "" {
		values.AddObject("type", obj.Direction)
		values.AddObject("action", obj.Action)

		if obj.SourceAddress == "" {
			delete.Append("source")
		} else {
			values.AddString("source", obj.SourceAddress)
		}
		if obj.DestinationAddress == "" {
			delete.Append("dest")
		} else {
			values.AddString("dest", obj.DestinationAddress)
		}

		if obj.Macro == MacroNone {
			delete.Append("macro")

			if obj.Protocol == ProtocolNone {
				delete.Append("proto")
			} else {
				values.ConditionalAddObject("proto", obj.Protocol, obj.Protocol != ProtocolNone)
			}
			if len(obj.SourcePorts) == 0 {
				delete.Append("sport")
			} else {
				values.ConditionalAddObject("sport", obj.SourcePorts, len(obj.SourcePorts) != 0)
			}
			if len(obj.DestinationPorts) == 0 {
				delete.Append("dport")
			} else {
				values.ConditionalAddObject("dport", obj.DestinationPorts, len(obj.DestinationPorts) != 0)
			}
		} else {
			delete.Append("proto")
			delete.Append("sport")
			delete.Append("dport")

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
		values.ConditionalAddObject(
			"delete",
			delete,
			delete.Len() != 0,
		)
	}

	return values, nil
}
