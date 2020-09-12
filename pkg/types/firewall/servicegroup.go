package firewall

type ServiceGroup interface {
	Name() string
	Description() string

	Digest() string

	Rename(name string) error
	GetProperties() (ServiceGroupProperties, error)
	SetProperties(props ServiceGroupProperties) error

	ListFirewallRules() ([]Rule, error)
	GetFirewallRule(pos uint) (Rule, error)
	AddFirewallRule(rule Rule) error
	EditFirewallRule(pos uint, rule Rule) error
	MoveFirewallRule(pos uint, newpos uint) error
	DeleteFirewallRule(pos uint, digest string) error
}

type ServiceGroupProperties struct {
	Description string

	Digest string
}
