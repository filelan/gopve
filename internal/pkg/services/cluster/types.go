package cluster

import (
	"github.com/xabinapal/gopve/internal/pkg/services/cluster/ha"
	"github.com/xabinapal/gopve/internal/pkg/utils"
	"github.com/xabinapal/gopve/pkg/types"
)

type Service struct {
	utils.Service

	haService *ha.Service
}

func NewService(api utils.API, client utils.Client) *Service {
	return &Service{
		Service: utils.Service{API: api, Client: client},

		haService: nil,
	}
}

func (svc *Service) HA() types.HighAvailabilityService {
	if svc.haService == nil {
		svc.haService = ha.NewService(svc.API, svc.Client)
	}

	return svc.haService
}
