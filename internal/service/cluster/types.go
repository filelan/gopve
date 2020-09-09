package cluster

import (
	"github.com/xabinapal/gopve/internal/client"
	"github.com/xabinapal/gopve/internal/service/cluster/ha"
	"github.com/xabinapal/gopve/pkg/service"
	"github.com/xabinapal/gopve/pkg/types/cluster"
)

type Service struct {
	client client.Client
	api    client.API

	haService *ha.Service
}

func NewService(cli client.Client, api client.API) *Service {
	return &Service{
		client: cli,
		api:    api,

		haService: ha.NewService(cli, api),
	}
}

func (svc *Service) HA() service.HighAvailability {
	return svc.haService
}

type Cluster struct {
	svc *Service

	mode cluster.Mode
	name string
}

func NewCluster(svc *Service, mode cluster.Mode, name string) *Cluster {
	return &Cluster{
		svc:  svc,
		mode: mode,
		name: name,
	}
}

func (obj *Cluster) Mode() cluster.Mode {
	return obj.mode
}

func (obj *Cluster) Name() string {
	return obj.name
}
