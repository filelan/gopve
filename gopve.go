package gopve

import (
	"github.com/xabinapal/gopve/service"
)

type GoPVE struct {
	QEMU service.QEMUServiceProvider
	LXC  service.LXCServiceProvider
}

func NewGoPVE(cfg *Config) (*GoPVE, error) {
	rootURI, err = cfg.GenerateRootURI()
	if err != nil {
		return nil, err
	}

	c, err := NewClient(rootURI, cfg.User, cfg.Password, cfg.InvalidCert)
	if err != nil {
		return nil, err
	}

	pve.QEMU = &QEMUService{client: c}
	pve.LXC = &LXCService{client: c}
	return pve, nil
}
