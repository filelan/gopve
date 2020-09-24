package test

import (
	"github.com/xabinapal/gopve/internal/service/cluster"
	"github.com/xabinapal/gopve/pkg/client/test"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func NewService() (*cluster.Service, *test.API, *mocks.Executor) {
	cli, exc := test.NewClient()
	api := test.NewAPI()

	return cluster.NewService(cli, api), api, exc
}
