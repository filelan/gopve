package firewall

type Alias interface {
	Name() string
	Description() string
	Address() string

	Digest() string

	Rename(name string) error
	GetProperties() (AliasProperties, error)
	SetProperties(props AliasProperties) error
}

type AliasProperties struct {
	Description string
	Address     string

	Digest string
}
