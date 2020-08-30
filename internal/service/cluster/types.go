package cluster

import (
	"github.com/xabinapal/gopve/internal/client"
	"github.com/xabinapal/gopve/internal/service/cluster/ha"
	"github.com/xabinapal/gopve/pkg/service"
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
