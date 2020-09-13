package test

import (
	"github.com/xabinapal/gopve/internal/service/vm"
	"github.com/xabinapal/gopve/pkg/client/test"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func NewService() (*vm.Service, *test.API, *mocks.Executor) {
	cli, exc := test.NewClient()
	api := test.NewAPI()
	return vm.NewService(cli, api), api, exc
}
