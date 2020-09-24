package node

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/firewall"
)

type getFirewallLogResponseJSON struct {
	LineNumber int    `json:"n"`
	Contents   string `json:"t"`
}

func (n *Node) GetFirewallLog(opts firewall.GetLogOptions) (firewall.LogEntries, error) {
	form := make(request.Values)

	form.ConditionalAddUint("start", opts.LineStart, opts.LineStart != 0)
	form.ConditionalAddUint("limit", opts.LineLimit, opts.LineLimit != 0)

	var res []getFirewallLogResponseJSON
	if err := n.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/firewall/log", n.name), form, &res); err != nil {
		return nil, err
	}

	entries := make(firewall.LogEntries)
	for _, entry := range res {
		entries[entry.LineNumber] = entry.Contents
	}

	return entries, nil
}

type getFirewallPropertiesResponseJSON struct {
	Enable           types.PVEBool     `json:"enable"`
	LogLevelIncoming firewall.LogLevel `json:"log_level_in"`
	LogLevelOutgoing firewall.LogLevel `json:"log_level_out"`

	LogTrackedConnections         types.PVEBool `json:"log_nf_conntrack"`
	AllowInvalidConnectionPackets types.PVEBool `json:"nf_conntrack_allow_invalid"`
	MaxTrackedConnections         *uint         `json:"nf_conntrack_max"`
	MaxConnectionEstablishTimeout *uint         `json:"nf_conntrack_tcp_timeout_established"`
	MaxConnectionSYNACKTimeout    *uint         `json:"nf_conntrack_tcp_timeout_syn_recv"`

	EnableNDP types.PVEBool `json:"ndp"`

	EnableSMURFS   types.PVEBool     `json:"nosmurfs"`
	SMURFSLogLevel firewall.LogLevel `json:"smurf_log_level"`

	EnableTCPFlagsFilter   types.PVEBool     `json:"tcpflags"`
	TCPFlagsFilterLogLevel firewall.LogLevel `json:"tcp_flags_log_level"`

	EnableSYNFloodProtection types.PVEBool `json:"protection_synflood"`
	SYNFloodProtectionRate   *uint         `json:"protection_synflood_rate"`
	SYNFloodProtectionBurst  *uint         `json:"protection_synflood_burst"`

	Digest string `json:"digest"`
}

func (obj getFirewallPropertiesResponseJSON) Map() (firewall.NodeProperties, error) {
	var maxTrackedConnections uint
	if obj.MaxTrackedConnections != nil {
		maxTrackedConnections = *obj.MaxTrackedConnections
	} else {
		maxTrackedConnections = firewall.DefaultMaxTrackedConnections
	}

	var maxConnectionEstablishTimeout uint
	if obj.MaxConnectionEstablishTimeout != nil {
		maxConnectionEstablishTimeout = *obj.MaxConnectionEstablishTimeout
	} else {
		maxConnectionEstablishTimeout = firewall.DefaultMaxConnectionEstablishTimeout
	}

	var maxConnectionSYNACKTimeout uint
	if obj.MaxConnectionSYNACKTimeout != nil {
		maxConnectionSYNACKTimeout = *obj.MaxConnectionSYNACKTimeout
	} else {
		maxConnectionSYNACKTimeout = firewall.DefaultMaxConectionSYNACKTimeout
	}

	var synFloodProtectionRate uint
	if obj.SYNFloodProtectionRate != nil {
		synFloodProtectionRate = *obj.SYNFloodProtectionRate
	} else {
		synFloodProtectionRate = firewall.DefaultSYNFloodProtectionRate
	}

	var synFloodProtectionBurst uint
	if obj.SYNFloodProtectionBurst != nil {
		synFloodProtectionBurst = *obj.SYNFloodProtectionBurst
	} else {
		synFloodProtectionBurst = firewall.DefaultSYNFloodProtectionBurst
	}

	return firewall.NodeProperties{
		Enable:           obj.Enable.Bool(),
		LogLevelIncoming: obj.LogLevelIncoming,
		LogLevelOutgoing: obj.LogLevelOutgoing,

		LogTrackedConnections:         obj.LogTrackedConnections.Bool(),
		AllowInvalidConnectionPackets: obj.AllowInvalidConnectionPackets.Bool(),
		MaxTrackedConnections:         maxTrackedConnections,
		MaxConnectionEstablishTimeout: maxConnectionEstablishTimeout,
		MaxConnectionSYNACKTimeout:    maxConnectionSYNACKTimeout,

		EnableNDP: obj.EnableNDP.Bool(),

		EnableSMURFS:   obj.EnableSMURFS.Bool(),
		SMURFSLogLevel: obj.SMURFSLogLevel,

		EnableTCPFlagsFilter:   obj.EnableTCPFlagsFilter.Bool(),
		TCPFlagsFilterLogLevel: obj.TCPFlagsFilterLogLevel,

		EnableSYNFloodProtection: obj.EnableSYNFloodProtection.Bool(),
		SYNFloodProtectionRate:   synFloodProtectionRate,
		SYNFloodProtectionBurst:  synFloodProtectionBurst,

		Digest: obj.Digest,
	}, nil
}

func (n *Node) GetFirewallProperties() (firewall.NodeProperties, error) {
	var res getFirewallPropertiesResponseJSON
	if err := n.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/firewall/options", n.name), nil, &res); err != nil {
		return firewall.NodeProperties{}, err
	}

	return res.Map()
}

func (n *Node) SetFirewallProperties(props firewall.NodeProperties) error {
	form, err := props.MapToValues()
	if err != nil {
		return err
	}

	return n.svc.client.Request(http.MethodPut, fmt.Sprintf("nodes/%s/firewall/options", n.name), form, nil)
}

type getFirewallRuleResponseJSON struct {
	Enable   types.PVEBool     `json:"enable"`
	Comment  string            `json:"comment"`
	LogLevel firewall.LogLevel `json:"log"`

	Type   string `json:"type"`
	Action string `json:"action"`

	Interface string `json:"iface"`

	SourceAddress      string `json:"source"`
	DestinationAddress string `json:"dest"`

	Macro    firewall.Macro    `json:"macro"`
	Protocol firewall.Protocol `json:"proto"`

	SourcePort      firewall.PortRanges `json:"sport"`
	DestinationPort firewall.PortRanges `json:"dport"`

	Digest string `json:"digest"`
}

func (obj getFirewallRuleResponseJSON) mapToRule() (firewall.Rule, error) {
	var direction firewall.Direction
	if err := (&direction).Unmarshal(obj.Type); err != nil {
		return firewall.Rule{}, err
	}

	var action firewall.Action
	if err := (&action).Unmarshal(obj.Action); err != nil {
		return firewall.Rule{}, err
	}

	return firewall.Rule{
		Enable:      obj.Enable.Bool(),
		Description: obj.Comment,
		LogLevel:    obj.LogLevel,

		Direction: direction,
		Action:    action,

		Interface: obj.Interface,

		SourceAddress:      obj.SourceAddress,
		DestinationAddress: obj.DestinationAddress,

		Macro:    obj.Macro,
		Protocol: obj.Protocol,

		SourcePorts:      obj.SourcePort,
		DestinationPorts: obj.DestinationPort,

		Digest: obj.Digest,
	}, nil
}

func (obj getFirewallRuleResponseJSON) mapToSecurityGroupRule() (firewall.Rule, error) {
	return firewall.Rule{
		Enable:      obj.Enable.Bool(),
		Description: obj.Comment,

		SecurityGroup: obj.Action,
		Interface:     obj.Interface,

		Digest: obj.Digest,
	}, nil
}

func (obj getFirewallRuleResponseJSON) Map() (firewall.Rule, error) {
	switch obj.Type {
	case "in", "out":
		return obj.mapToRule()
	case "group":
		return obj.mapToSecurityGroupRule()
	default:
		return firewall.Rule{}, fmt.Errorf("unknown firewall rule type `%s`", obj.Type)
	}
}

func (obj *Node) ListFirewallRules() ([]firewall.Rule, error) {
	var res []getFirewallRuleResponseJSON
	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/firewall/rules", obj.name), nil, &res); err != nil {
		return nil, err
	}

	rules := make([]firewall.Rule, len(res))

	for i, rule := range res {
		r, err := rule.Map()
		if err != nil {
			return nil, err
		}

		rules[i] = r
	}

	return rules, nil
}

func (obj *Node) GetFirewallRule(pos uint) (firewall.Rule, error) {
	var res getFirewallRuleResponseJSON
	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/firewall/rules/%d", obj.name, pos), nil, &res); err != nil {
		return firewall.Rule{}, err
	}

	return res.Map()
}

func (obj *Node) AddFirewallRule(rule firewall.Rule) error {
	form, err := rule.MapToValues(false)
	if err != nil {
		return err
	}

	return obj.svc.client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/firewall/rules", obj.name), form, nil)
}

func (obj *Node) EditFirewallRule(pos uint, rule firewall.Rule) error {
	form, err := rule.MapToValues(true)
	if err != nil {
		return err
	}

	return obj.svc.client.Request(http.MethodPut, fmt.Sprintf("nodes/%s/firewall/rules/%d", obj.name, pos), form, nil)
}

func (obj *Node) MoveFirewallRule(pos uint, newpos uint) error {
	form := request.Values{}
	form.AddUint("moveto", newpos)

	return obj.svc.client.Request(http.MethodPut, fmt.Sprintf("nodes/%s/firewall/rules/%d", obj.name, pos), form, nil)
}

func (obj *Node) DeleteFirewallRule(pos uint, digest string) error {
	var form request.Values

	if digest != "" {
		form = request.Values{
			"digest": {digest},
		}
	}

	return obj.svc.client.Request(http.MethodDelete, fmt.Sprintf("nodes/%s/firewall/rules/%d", obj.name, pos), form, nil)
}
