package cluster

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/firewall"
)

type getFirewallPropertiesResponseJSON struct {
	Enable         types.PVEBool `json:"enable"`
	EnableEbtables types.PVEBool `json:"ebtables"`

	DefaultInputPolicy  firewall.Action `json:"policy_in"`
	DefaultOutputPolicy firewall.Action `json:"policy_out"`

	LogLimit firewall.LogLimit `json:"log_ratelimit"`

	Digest string `json:"digest"`
}

func (obj getFirewallPropertiesResponseJSON) Map() (firewall.ClusterProperties, error) {
	return firewall.ClusterProperties{
		Enable:         obj.Enable.Bool(),
		EnableEbtables: obj.EnableEbtables.Bool(),

		DefaultInputPolicy:  obj.DefaultInputPolicy,
		DefaultOutputPolicy: obj.DefaultOutputPolicy,

		LogLimit: obj.LogLimit,

		Digest: obj.Digest,
	}, nil
}

func (svc *Service) GetFirewallProperties() (firewall.ClusterProperties, error) {
	var res getFirewallPropertiesResponseJSON
	if err := svc.client.Request(http.MethodGet, "cluster/firewall/options", nil, &res); err != nil {
		return firewall.ClusterProperties{}, err
	}

	return res.Map()
}

func (svc *Service) SetFirewallProperties(props firewall.ClusterProperties) error {
	form, err := props.MapToValues()
	if err != nil {
		return err
	}

	return svc.client.Request(http.MethodPut, "cluster/firewall/options", form, nil)
}

type getFirewallAliasResponseJSON struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`
	CIDR    string `json:"cidr"`

	Digest string `json:"digest"`
}

func (obj getFirewallAliasResponseJSON) Map(svc *Service) (firewall.Alias, error) {
	return NewFirewallAlias(svc, obj.Name, obj.Comment, obj.CIDR, obj.Digest), nil
}

func (svc *Service) ListFirewallAliases() ([]firewall.Alias, error) {
	var res []getFirewallAliasResponseJSON
	if err := svc.client.Request(http.MethodGet, "cluster/firewall/aliases", nil, &res); err != nil {
		return nil, err
	}

	aliases := make([]firewall.Alias, len(res))
	for i, alias := range res {
		a, err := alias.Map(svc)
		if err != nil {
			return nil, err
		}

		aliases[i] = a
	}

	return aliases, nil
}

func (svc *Service) GetFirewallAlias(name string) (firewall.Alias, error) {
	var res getFirewallAliasResponseJSON
	if err := svc.client.Request(http.MethodGet, fmt.Sprintf("cluster/firewall/aliases/%s", name), nil, &res); err != nil {
		return nil, err
	}

	return res.Map(svc)
}

type getFirewallIPSetResponseJSON struct {
	Name    string `json:"name"`
	Comment string `json:"comment"`

	Digest string `json:"digest"`
}

func (obj getFirewallIPSetResponseJSON) Map(svc *Service) (firewall.IPSet, error) {
	return NewFirewallIPSet(svc, obj.Name, obj.Comment, obj.Digest), nil
}

func (svc *Service) ListFirewallIPSets() ([]firewall.IPSet, error) {
	var res []getFirewallIPSetResponseJSON
	if err := svc.client.Request(http.MethodGet, "cluster/firewall/ipset", nil, &res); err != nil {
		return nil, err
	}

	ipSets := make([]firewall.IPSet, len(res))
	for i, ipset := range res {
		a, err := ipset.Map(svc)
		if err != nil {
			return nil, err
		}

		ipSets[i] = a
	}

	return ipSets, nil
}

func (svc *Service) GetFirewallIPSet(name string) (firewall.IPSet, error) {
	ipSets, err := svc.ListFirewallIPSets()
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

func (obj getFirewallServiceGroupResponseJSON) Map(svc *Service) (firewall.ServiceGroup, error) {
	alias := NewFirewallServiceGroup(svc, obj.Name, obj.Comment, obj.Digest)
	return alias, nil
}

func (svc *Service) ListFirewallServiceGroups() ([]firewall.ServiceGroup, error) {
	var res []getFirewallServiceGroupResponseJSON
	if err := svc.client.Request(http.MethodGet, "cluster/firewall/groups", nil, &res); err != nil {
		return nil, err
	}

	groups := make([]firewall.ServiceGroup, len(res))
	for i, group := range res {
		a, err := group.Map(svc)
		if err != nil {
			return nil, err
		}

		groups[i] = a
	}

	return groups, nil
}

func (svc *Service) GetFirewallServiceGroup(name string) (firewall.ServiceGroup, error) {
	serviceGroups, err := svc.ListFirewallServiceGroups()
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
	if obj.Type == "in" || obj.Type == "out" {
		return obj.mapToRule()
	} else if obj.Type == "group" {
		return obj.mapToSecurityGroupRule()
	} else {
		return firewall.Rule{}, fmt.Errorf("unknown firewall rule type `%s`", obj.Type)
	}
}

func (svc *Service) ListFirewallRules() ([]firewall.Rule, error) {
	var res []getFirewallRuleResponseJSON
	if err := svc.client.Request(http.MethodGet, "cluster/firewall/rules", nil, &res); err != nil {
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

func (svc *Service) GetFirewallRule(pos uint) (firewall.Rule, error) {
	var res getFirewallRuleResponseJSON
	if err := svc.client.Request(http.MethodGet, fmt.Sprintf("cluster/firewall/rules/%d", pos), nil, &res); err != nil {
		return firewall.Rule{}, err
	}

	return res.Map()
}

func (svc *Service) AddFirewallRule(rule firewall.Rule) error {
	rule.Digest = ""

	form, err := rule.MapToValues(false)
	if err != nil {
		return err
	}

	return svc.client.Request(http.MethodPost, "cluster/firewall/rules", form, nil)
}

func (svc *Service) EditFirewallRule(pos uint, rule firewall.Rule) error {
	form, err := rule.MapToValues(true)
	if err != nil {
		return err
	}

	return svc.client.Request(http.MethodPut, fmt.Sprintf("cluster/firewall/rules/%d", pos), form, nil)
}

func (svc *Service) MoveFirewallRule(pos uint, newpos uint) error {
	form := request.Values{}
	form.AddUint("moveto", newpos)

	return svc.client.Request(http.MethodPut, fmt.Sprintf("cluster/firewall/rules/%d", pos), form, nil)
}

func (svc *Service) DeleteFirewallRule(pos uint, digest string) error {
	var form request.Values

	if digest != "" {
		form = request.Values{
			"digest": {digest},
		}
	}

	return svc.client.Request(http.MethodDelete, fmt.Sprintf("cluster/firewall/rules/%d", pos), form, nil)
}
