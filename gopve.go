package gopve

import (
	"github.com/xabinapal/gopve/internal"
	"github.com/xabinapal/gopve/service"
)

type GoPVE struct {
	Node    service.NodeServiceProvider
	Storage service.StorageServiceProvider
}

func NewGoPVE(cfg *Config) (*GoPVE, error) {
	rootURI, err := cfg.GenerateRootURI()
	if err != nil {
		return nil, err
	}

	c, err := internal.NewClient(rootURI, cfg.User, cfg.Password, cfg.InvalidCert)
	if err != nil {
		return nil, err
	}

	pve := &GoPVE{
		Node:    service.NewNodeService(c),
		Storage: service.NewStorageService(c),
	}

	return pve, nil
}
