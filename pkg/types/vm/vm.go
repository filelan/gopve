package vm

import (
	"github.com/xabinapal/gopve/pkg/types"
	"github.com/xabinapal/gopve/pkg/types/errors"
	"github.com/xabinapal/gopve/pkg/types/firewall"
	"github.com/xabinapal/gopve/pkg/types/task"
)

type VirtualMachine interface {
	VMID() uint
	Kind() Kind
	Node() string

	Name() string
	Template() bool

	GetProperties() (Properties, error)

	Description() (string, error)

	Digest() (string, error)

	GetStatus() (Status, error)

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

type Properties struct {
	Description string
	Protected   bool

	StartOnBoot     bool
	StartupOrder    int
	StartDelay      int
	ShutdownTimeout int

	Digest string
}

const (
	mkPropertyDescription = "description"
	mkPropertyProtected   = "protection"

	mkPropertyStartOnBoot        = "onboot"
	mkDictPropertyStartup        = "startup"
	mkKeyPropertyStartupOrder    = "order"
	mkKeyPropertyStartDelay      = "up"
	mkKeyPropertyShutdownTimeout = "down"

	mkPropertyDigest = "digest"
)

const (
	DefaultPropertyDescription string = ""
	DefaultPropertyProtected   bool   = false

	DefaultPropertyStartOnBoot     bool = false
	DefaultPropertyStartupOrder    int  = -1
	DefaultPropertyStartDelay      int  = -1
	DefaultPropertyShutdownTimeout int  = -1
)

func NewProperties(props types.Properties) (obj Properties, err error) {
	return obj, errors.ChainUntilFail(
		func() (err error) {
			return props.SetString(
				mkPropertyDescription,
				&obj.Description,
				DefaultPropertyDescription,
				nil,
			)
		},
		func() (err error) {
			return props.SetBool(
				mkPropertyProtected,
				&obj.Protected,
				DefaultPropertyProtected,
				nil,
			)
		},
		func() (err error) {
			startupOptions, err := props.GetAsDict(
				mkDictPropertyStartup,
				",",
				"=",
				false,
			)
			if err != nil && !errors.ErrMissingProperty.IsBase(err) {
				return err
			}

			return errors.ChainUntilFail(
				func() error {
					return startupOptions.InjectInt(
						mkKeyPropertyStartupOrder,
						&obj.StartupOrder,
						DefaultPropertyStartupOrder,
						nil,
					)
				},
				func() error {
					return startupOptions.InjectInt(
						mkKeyPropertyStartDelay,
						&obj.StartDelay,
						DefaultPropertyStartDelay,
						nil,
					)
				},
				func() error {
					return startupOptions.InjectInt(
						mkKeyPropertyShutdownTimeout,
						&obj.ShutdownTimeout,
						DefaultPropertyShutdownTimeout,
						nil,
					)
				},
			)
		},
		func() (err error) {
			return props.SetRequiredString(mkPropertyDigest, &obj.Digest, nil)
		},
	)
}
