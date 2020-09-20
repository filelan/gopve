package test

import (
	"github.com/xabinapal/gopve/internal/service/storage"
	"github.com/xabinapal/gopve/pkg/client/test"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func NewService() (*storage.Service, *test.API, *mocks.Executor) {
	cli, exc := test.NewClient()
	api := test.NewAPI()
	return storage.NewService(cli, api), api, exc
}
