package cluster

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/firewall"
)

type FirewallAlias struct {
	svc *Service

	name        string
	description string
	address     string

	digest string
}

func NewFirewallAlias(svc *Service, name string, description string, address string, digest string) *FirewallAlias {
	return &FirewallAlias{
		svc:         svc,
		name:        name,
		description: description,
		address:     address,
		digest:      digest,
	}
}

func (obj *FirewallAlias) Name() string {
	return obj.name
}

func (obj *FirewallAlias) Description() string {
	return obj.description
}

func (obj *FirewallAlias) Address() string {
	return obj.address
}

func (obj *FirewallAlias) Digest() string {
	return obj.digest
}

func (obj *FirewallAlias) Rename(name string) error {
	if err := obj.svc.client.Request(http.MethodPut, fmt.Sprintf("cluster/firewall/aliases/%s", obj.name), request.Values{
		"rename": {name},
	}, nil); err != nil {
		return err
	}

	obj.name = name

	return nil
}

func (obj *FirewallAlias) GetProperties() (firewall.AliasProperties, error) {
	return firewall.AliasProperties{
		Description: obj.description,
		Address:     obj.address,
		Digest:      obj.digest,
	}, nil
}

func (obj *FirewallAlias) SetProperties(props firewall.AliasProperties) error {
	if err := obj.svc.client.Request(http.MethodPut, fmt.Sprintf("cluster/firewall/aliases/%s", obj.name), request.Values{
		"comment": {props.Description},
		"cidr":    {props.Address},
		"digest":  {props.Digest},
	}, nil); err != nil {
		return nil
	}

	obj.description = props.Description
	obj.address = props.Address
	obj.digest = props.Digest

	return nil
}

type FirewallIPSet struct {
	svc *Service

	name        string
	description string

	digest string
}

func NewFirewallIPSet(svc *Service, name string, description, digest string) *FirewallIPSet {
	return &FirewallIPSet{
		svc:         svc,
		name:        name,
		description: description,
		digest:      digest,
	}
}

func (obj *FirewallIPSet) Name() string {
	return obj.name
}

func (obj *FirewallIPSet) Description() string {
	return obj.description
}

func (obj *FirewallIPSet) Digest() string {
	return obj.digest
}

func (obj *FirewallIPSet) Rename(name string) error {
	if err := obj.svc.client.Request(http.MethodPost, "cluster/firewall/ipset", request.Values{
		"group":  {obj.name},
		"rename": {name},
	}, nil); err != nil {
		return err
	}

	obj.name = name
	return nil
}

func (obj *FirewallIPSet) GetProperties() (firewall.IPSetProperties, error) {
	return firewall.IPSetProperties{
		Description: obj.description,
		Digest:      obj.digest,
	}, nil
}

func (obj *FirewallIPSet) SetProperties(props firewall.IPSetProperties) error {
	if err := obj.svc.client.Request(http.MethodPost, "cluster/firewall/ipset", request.Values{
		"group":   {obj.name},
		"comment": {props.Description},
		"digest":  {props.Digest},
	}, nil); err != nil {
		return nil
	}

	obj.description = props.Description
	obj.digest = props.Digest

	return nil
}

type getFirewallIPSetAddressResponseJSON struct {
	CIDR    string        `json:"cidr"`
	Comment string        `json:"comment"`
	NoMatch types.PVEBool `json:"nomatch"`
	Digest  string        `json:"digest"`
}

func (obj getFirewallIPSetAddressResponseJSON) Map() (firewall.IPSetAddress, error) {
	return firewall.IPSetAddress{
		Address:     obj.CIDR,
		Description: obj.Comment,
		NoMatch:     obj.NoMatch.Bool(),
		Digest:      obj.Digest,
	}, nil
}

func (obj *FirewallIPSet) ListAddresses() ([]firewall.IPSetAddress, error) {
	var res []getFirewallIPSetAddressResponseJSON
	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("cluster/firewall/ipset/%s", obj.name), nil, &res); err != nil {
		return nil, err
	}

	addresses := make([]firewall.IPSetAddress, len(res))
	for i, address := range res {
		a, err := address.Map()
		if err != nil {
			return nil, err
		}

		addresses[i] = a
	}

	return addresses, nil
}

func (obj *FirewallIPSet) GetAddress(cidr string) (firewall.IPSetAddress, error) {
	var res getFirewallIPSetAddressResponseJSON
	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("cluster/firewall/ipset/%s/%s", obj.name, cidr), nil, &res); err != nil {
		return firewall.IPSetAddress{}, err
	}

	return res.Map()
}

func (obj *FirewallIPSet) AddAddress(address firewall.IPSetAddress) error {
	address.Digest = ""

	form, err := address.MapToValues()
	if err != nil {
		return err
	}

	form.AddString("cidr", address.Address)

	return obj.svc.client.Request(http.MethodPost, fmt.Sprintf("cluster/firewall/ipset/%s", obj.name), form, nil)
}

func (obj *FirewallIPSet) EditAddress(address firewall.IPSetAddress) error {
	form, err := address.MapToValues()
	if err != nil {
		return err
	}

	return obj.svc.client.Request(http.MethodPut, fmt.Sprintf("cluster/firewall/ipset/%s/%s", obj.name, address.Address), form, nil)
}

func (obj *FirewallIPSet) DeleteAddress(cidr string, digest string) error {
	var form request.Values

	if digest != "" {
		form = request.Values{
			"digest": {digest},
		}
	}

	return obj.svc.client.Request(http.MethodDelete, fmt.Sprintf("cluster/firewall/ipset/%s/%s", obj.name, cidr), form, nil)
}

type FirewallServiceGroup struct {
	svc *Service

	name        string
	description string

	digest string
}

func NewFirewallServiceGroup(svc *Service, name string, description string, digest string) *FirewallServiceGroup {
	return &FirewallServiceGroup{
		svc:         svc,
		name:        name,
		description: description,
		digest:      digest,
	}
}

func (obj *FirewallServiceGroup) Name() string {
	return obj.name
}

func (obj *FirewallServiceGroup) Description() string {
	return obj.description
}

func (obj *FirewallServiceGroup) Digest() string {
	return obj.digest
}

func (obj *FirewallServiceGroup) Rename(name string) error {
	if err := obj.svc.client.Request(http.MethodPost, "cluster/firewall/groups", request.Values{
		"group":  {obj.name},
		"rename": {name},
	}, nil); err != nil {
		return err
	}

	obj.name = name
	return nil
}

func (obj *FirewallServiceGroup) GetProperties() (firewall.ServiceGroupProperties, error) {
	return firewall.ServiceGroupProperties{
		Description: obj.description,
		Digest:      obj.digest,
	}, nil
}

func (obj *FirewallServiceGroup) SetProperties(props firewall.ServiceGroupProperties) error {
	if err := obj.svc.client.Request(http.MethodPost, "cluster/firewall/groups", request.Values{
		"group":   {obj.name},
		"comment": {props.Description},
		"digest":  {props.Digest},
	}, nil); err != nil {
		return nil
	}

	obj.description = props.Description
	obj.digest = props.Digest

	return nil
}

func (obj *FirewallServiceGroup) ListFirewallRules() ([]firewall.Rule, error) {
	var res []getFirewallRuleResponseJSON
	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("cluster/firewall/groups/%s", obj.name), nil, &res); err != nil {
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

func (obj *FirewallServiceGroup) GetFirewallRule(pos uint) (firewall.Rule, error) {
	var res getFirewallRuleResponseJSON
	if err := obj.svc.client.Request(http.MethodGet, fmt.Sprintf("cluster/firewall/groups/%s/%d", obj.name, pos), nil, &res); err != nil {
		return firewall.Rule{}, err
	}

	return res.Map()
}

func (obj *FirewallServiceGroup) AddFirewallRule(rule firewall.Rule) error {
	rule.Digest = ""

	form, err := rule.MapToValues(false)
	if err != nil {
		return err
	}

	return obj.svc.client.Request(http.MethodPost, fmt.Sprintf("cluster/firewall/groups/%s", obj.name), form, nil)
}

func (obj *FirewallServiceGroup) EditFirewallRule(pos uint, rule firewall.Rule) error {
	form, err := rule.MapToValues(true)
	if err != nil {
		return err
	}

	return obj.svc.client.Request(http.MethodPut, fmt.Sprintf("cluster/firewall/groups/%s/%d", obj.name, pos), form, nil)
}

func (obj *FirewallServiceGroup) MoveFirewallRule(pos uint, newpos uint) error {
	form := request.Values{}
	form.AddUint("moveto", newpos)

	return obj.svc.client.Request(http.MethodPut, fmt.Sprintf("cluster/firewall/groups/%s/%d", obj.name, pos), form, nil)
}

func (obj *FirewallServiceGroup) DeleteFirewallRule(pos uint, digest string) error {
	var form request.Values

	if digest != "" {
		form = request.Values{
			"digest": {digest},
		}
	}

	return obj.svc.client.Request(http.MethodDelete, fmt.Sprintf("cluster/firewall/groups/%s/%d", obj.name, pos), form, nil)
}
