package task

import (
	"time"

	"github.com/xabinapal/gopve/internal/client"
)

type Service struct {
	client client.Client
	api    client.API

	poolingInterval time.Duration
}

func NewService(
	cli client.Client,
	api client.API,
	poolingInterval time.Duration,
) *Service {
	return &Service{
		client: cli,
		api:    api,

		poolingInterval: poolingInterval,
	}
}
