package firewall

import "github.com/xabinapal/gopve/pkg/request"

type IPSet interface {
	Name() string
	Description() string

	Digest() string

	Rename(name string) error
	GetProperties() (IPSetProperties, error)
	SetProperties(props IPSetProperties) error

	ListAddresses() ([]IPSetAddress, error)
	GetAddress(cidr string) (IPSetAddress, error)
	AddAddress(address IPSetAddress) error
	EditAddress(address IPSetAddress) error
	DeleteAddress(cidr string, digest string) error
}

type IPSetProperties struct {
	Description string

	Digest string
}

type IPSetAddress struct {
	Address     string
	Description string
	NoMatch     bool

	Digest string
}

func (obj IPSetAddress) MapToValues() (request.Values, error) {
	values := request.Values{}
	values.AddString("comment", obj.Description)
	values.AddBool("nomatch", obj.NoMatch)
	values.ConditionalAddString("digest", obj.Digest, obj.Digest != "")

	return values, nil
}
