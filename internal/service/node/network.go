package node

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/node"
)

func (n *Node) GetDNSSettings() (node.DNSSettings, error) {
	var res node.DNSSettings
	if err := n.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/dns", n.name), nil, &res); err != nil {
		return node.DNSSettings{}, err
	}

	return res, nil
}

func (n *Node) SetDNSSettings(settings node.DNSSettings) error {
	return n.svc.client.Request(
		http.MethodPut,
		fmt.Sprintf("nodes/%s/dns", n.name),
		request.Values{
			"dns1":   {settings.FirstDNS},
			"dns2":   {settings.SecondDNS},
			"dns3":   {settings.ThirdDNS},
			"search": {settings.SearchDomain},
		},
		nil,
	)
}

func (n *Node) GetHostsFile() (node.HostsFile, error) {
	var res node.HostsFile
	if err := n.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/hosts", n.name), nil, &res); err != nil {
		return node.HostsFile{}, err
	}

	return res, nil
}

func (n *Node) SetHostsFile(file node.HostsFile) error {
	form := request.Values{
		"data": {file.Contents},
	}
	form.ConditionalAddString("digest", file.Digest, file.Digest != "")

	return n.svc.client.Request(
		http.MethodPut,
		fmt.Sprintf("nodes/%s/hosts", n.name),
		form,
		nil,
	)
}
