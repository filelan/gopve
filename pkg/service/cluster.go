package service

import (
	"github.com/xabinapal/gopve/pkg/types/cluster"
	"github.com/xabinapal/gopve/pkg/types/firewall"
	"github.com/xabinapal/gopve/pkg/types/task"
)

//go:generate mockery --case snake --name Cluster

type Cluster interface {
	HA() HighAvailability

	Get() (cluster.Cluster, error)
	Create(name string, props cluster.NodeProperties) (task.Task, error)
	Join(
		hostname, password, fingerprint string,
		props cluster.NodeProperties,
	) (task.Task, error)

	GetFirewallProperties() (firewall.ClusterProperties, error)
	SetFirewallProperties(props firewall.ClusterProperties) error

	ListFirewallAliases() ([]firewall.Alias, error)
	GetFirewallAlias(name string) (firewall.Alias, error)

	ListFirewallIPSets() ([]firewall.IPSet, error)
	GetFirewallIPSet(name string) (firewall.IPSet, error)

	ListFirewallServiceGroups() ([]firewall.ServiceGroup, error)
	GetFirewallServiceGroup(name string) (firewall.ServiceGroup, error)

	ListFirewallRules() ([]firewall.Rule, error)
	GetFirewallRule(pos uint) (firewall.Rule, error)
	AddFirewallRule(rule firewall.Rule) error
	EditFirewallRule(pos uint, rule firewall.Rule) error
	MoveFirewallRule(pos uint, newpos uint) error
	DeleteFirewallRule(pos uint, digest string) error
}

type HighAvailability interface {
	ListGroups() ([]cluster.HighAvailabilityGroup, error)
	GetGroup(name string) (cluster.HighAvailabilityGroup, error)
	CreateGroup(
		name string,
		props cluster.HighAvailabilityGroupProperties,
		nodes cluster.HighAvailabilityGroupNodes,
	) (cluster.HighAvailabilityGroup, error)
}
