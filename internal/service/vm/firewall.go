package vm

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

func (obj *VirtualMachine) GetFirewallLog(opts firewall.GetLogOptions) (firewall.LogEntries, error) {
	form := make(request.Values)

	form.ConditionalAddUint("start", opts.LineStart, opts.LineStart != 0)
	form.ConditionalAddUint("limit", opts.LineLimit, opts.LineLimit != 0)

	var res []getFirewallLogResponseJSON
	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/%s/%d/firewall/log", obj.node, obj.kind.String(), obj.vmid), form, &res); err != nil {
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

	DefaultInputPolicy  firewall.Action `json:"policy_in"`
	DefaultOutputPolicy firewall.Action `json:"policy_out"`

	EnableNDP       types.PVEBool `json:"ndp"`
	EnableRADV      types.PVEBool `json:"radv"`
	EnableDHCP      types.PVEBool `json:"dhcp"`
	EnableMACFilter types.PVEBool `json:"macfilter"`
	EnableIPFilter  types.PVEBool `json:"ipfilter"`

	Digest string `json:"digest"`
}

func (obj getFirewallPropertiesResponseJSON) Map() (firewall.VMProperties, error) {
	return firewall.VMProperties{
		Enable:           obj.Enable.Bool(),
		LogLevelIncoming: obj.LogLevelIncoming,
		LogLevelOutgoing: obj.LogLevelOutgoing,

		DefaultInputPolicy:  obj.DefaultInputPolicy,
		DefaultOutputPolicy: obj.DefaultOutputPolicy,

		EnableNDP:       obj.EnableNDP.Bool(),
		EnableRADV:      obj.EnableRADV.Bool(),
		EnableDHCP:      obj.EnableDHCP.Bool(),
		EnableMACFilter: obj.EnableMACFilter.Bool(),
		EnableIPFilter:  obj.EnableIPFilter.Bool(),

		Digest: obj.Digest,
	}, nil
}

func (obj *VirtualMachine) GetFirewallProperties() (firewall.VMProperties, error) {
	var res getFirewallPropertiesResponseJSON
	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/%s/%d/firewall/options", obj.node, obj.kind.String(), obj.vmid), nil, &res); err != nil {
		return firewall.VMProperties{}, err
	}

	return res.Map()
}

func (obj *VirtualMachine) SetFirewallProperties(props firewall.VMProperties) error {
	form, err := props.MapToValues()
	if err != nil {
		return err
	}

	return obj.svc.client.Request(http.MethodPut, fmt.Sprintf("nodes/%s/%s/%d/firewall/options", obj.node, obj.kind.String(), obj.vmid), form, nil)
}

type getFirewallAliasResponseJSON struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	CIDR    string `json:"cidr"`

	Digest string `json:"digest"`
}

func (obj getFirewallAliasResponseJSON) Map(vm *VirtualMachine) (firewall.Alias, error) {
	return NewFirewallAlias(vm, obj.Name, obj.Comment, obj.CIDR, obj.Digest), nil
}

func (obj *VirtualMachine) ListFirewallAliases() ([]firewall.Alias, error) {
	var res []getFirewallAliasResponseJSON
	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/%s/%d/firewall/aliases", obj.node, obj.kind.String(), obj.vmid), nil, &res); err != nil {
		return nil, err
	}

	aliases := make([]firewall.Alias, len(res))
	for i, alias := range res {
		a, err := alias.Map(obj)
		if err != nil {
			return nil, err
		}

		aliases[i] = a
	}

	return aliases, nil
}

func (obj *VirtualMachine) GetFirewallAlias(name string) (firewall.Alias, error) {
	var res getFirewallAliasResponseJSON
	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/%s/%d/firewall/aliases/%s", obj.node, obj.kind.String(), obj.vmid, name), nil, &res); err != nil {
		return nil, err
	}

	return res.Map(obj)
}

type getFirewallIPSetResponseJSON struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`

	Digest string `json:"digest"`
}

func (obj getFirewallIPSetResponseJSON) Map(virtualMachine *VirtualMachine) (firewall.IPSet, error) {
	return NewFirewallIPSet(virtualMachine, obj.Name, obj.Comment, obj.Digest), nil
}

func (obj *VirtualMachine) ListFirewallIPSets() ([]firewall.IPSet, error) {
	var res []getFirewallIPSetResponseJSON

	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/%s/%d/firewall/ipset", obj.node, obj.kind.String(), obj.vmid), nil, &res); err != nil {
		return nil, err
	}

	ipSets := make([]firewall.IPSet, len(res))

	for i, ipset := range res {
		a, err := ipset.Map(obj)
		if err != nil {
			return nil, err
		}

		ipSets[i] = a
	}

	return ipSets, nil
}

func (obj *VirtualMachine) GetFirewallIPSet(name string) (firewall.IPSet, error) {
	ipSets, err := obj.ListFirewallIPSets()
	if err != nil {
		return nil, err
	}

	for _, ipSet := range ipSets {
		if ipSet.Name() == name {
			return ipSet, nil
		}
	}

	return nil, fmt.Errorf("ipset not found")
}

type getFirewallServiceGroupResponseJSON struct {
	Name    string `json:"group"`
	Comment string `json:"comment"`

	Digest string `json:"digest"`
}

func (obj getFirewallServiceGroupResponseJSON) Map(vm *VirtualMachine) (firewall.ServiceGroup, error) {
	alias := NewFirewallServiceGroup(vm, obj.Name, obj.Comment, obj.Digest)
	return alias, nil
}

func (obj *VirtualMachine) ListFirewallServiceGroups() ([]firewall.ServiceGroup, error) {
	var res []getFirewallServiceGroupResponseJSON
	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/%s/%d/firewall/groups", obj.node, obj.kind.String(), obj.vmid), nil, &res); err != nil {
		return nil, err
	}

	groups := make([]firewall.ServiceGroup, len(res))

	for i, group := range res {
		a, err := group.Map(obj)
		if err != nil {
			return nil, err
		}

		groups[i] = a
	}

	return groups, nil
}

func (obj *VirtualMachine) GetFirewallServiceGroup(name string) (firewall.ServiceGroup, error) {
	serviceGroups, err := obj.ListFirewallServiceGroups()
	if err != nil {
		return nil, err
	}

	for _, serviceGroup := range serviceGroups {
		if serviceGroup.Name() == name {
			return serviceGroup, nil
		}
	}

	return nil, fmt.Errorf("service group not found")
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

func (obj *VirtualMachine) ListFirewallRules() ([]firewall.Rule, error) {
	var res []getFirewallRuleResponseJSON
	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/%s/%d/firewall/rules", obj.node, obj.kind.String(), obj.vmid), nil, &res); err != nil {
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

func (obj *VirtualMachine) GetFirewallRule(pos uint) (firewall.Rule, error) {
	var res getFirewallRuleResponseJSON
	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/%s/%d/firewall/rules/%d", obj.node, obj.kind.String(), obj.vmid, pos), nil, &res); err != nil {
		return firewall.Rule{}, err
	}

	return res.Map()
}

func (obj *VirtualMachine) AddFirewallRule(rule firewall.Rule) error {
	rule.Digest = ""

	form, err := rule.MapToValues(false)
	if err != nil {
		return err
	}

	return obj.svc.client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/%s/%d/firewall/rules", obj.node, obj.kind.String(), obj.vmid), form, nil)
}

func (obj *VirtualMachine) EditFirewallRule(pos uint, rule firewall.Rule) error {
	form, err := rule.MapToValues(true)
	if err != nil {
		return err
	}

	return obj.svc.client.Request(http.MethodPut, fmt.Sprintf("nodes/%s/%s/%d/firewall/rules/%d", obj.node, obj.kind.String(), obj.vmid, pos), form, nil)
}

func (obj *VirtualMachine) MoveFirewallRule(pos uint, newpos uint) error {
	form := request.Values{}
	form.AddUint("moveto", newpos)

	return obj.svc.client.Request(http.MethodPut, fmt.Sprintf("nodes/%s/%s/%d/firewall/rules/%d", obj.node, obj.kind.String(), obj.vmid, pos), form, nil)
}

func (obj *VirtualMachine) DeleteFirewallRule(pos uint, digest string) error {
	var form request.Values

	if digest != "" {
		form = request.Values{
			"digest": {digest},
		}
	}

	return obj.svc.client.Request(http.MethodDelete, fmt.Sprintf("nodes/%s/%s/%d/firewall/rules/%d", obj.node, obj.kind.String(), obj.vmid, pos), form, nil)
}
