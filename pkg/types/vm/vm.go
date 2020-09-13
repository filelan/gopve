package vm

import (
	"github.com/xabinapal/gopve/pkg/types/firewall"
	"github.com/xabinapal/gopve/pkg/types/node"
	"github.com/xabinapal/gopve/pkg/types/task"
)

type VirtualMachine interface {
	Node() string

	GetNode() (node.Node, error)
	Kind() Kind

	VMID() uint
	Name() string
	IsTemplate() bool

	Status() (Status, error)

	Clone(options CloneOptions) (task.Task, error)

	Start() (task.Task, error)
	Stop() (task.Task, error)
	Reset() (task.Task, error)
	Shutdown() (task.Task, error)
	Reboot() (task.Task, error)
	Suspend() (task.Task, error)
	Resume() (task.Task, error)

	ListSnapshots() ([]Snapshot, error)
	GetSnapshot(name string) (Snapshot, error)
	CreateSnapshot(name string, props SnapshotProperties) (task.Task, error)
	RollbackToSnapshot(name string) (task.Task, error)

	GetFirewallLog(opts firewall.GetLogOptions) (firewall.LogEntries, error)
	GetFirewallProperties() (firewall.VMProperties, error)
	SetFirewallProperties(props firewall.VMProperties) error

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
