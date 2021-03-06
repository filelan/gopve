package test

import (
	"github.com/xabinapal/gopve/internal/service/pool"
	"github.com/xabinapal/gopve/pkg/client/test"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func NewService() (*pool.Service, *test.API, *mocks.Executor) {
	cli, exc := test.NewClient()
	api := test.NewAPI()
	return pool.NewService(cli, api), api, exc
}
