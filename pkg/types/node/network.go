package node

type DNSSettings struct {
	FirstDNS     string `json:"dns1"`
	SecondDNS    string `json:"dns2"`
	ThirdDNS     string `json:"dns3"`
	SearchDomain string `json:"search"`
}

type HostsFile struct {
	Contents string `json:"data"`
	Digest   string `json:"digest"`
}
