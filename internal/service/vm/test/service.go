package test

import (
	"github.com/xabinapal/gopve/internal/service/vm"
	"github.com/xabinapal/gopve/pkg/client/test"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func NewService() (*vm.Service, *mocks.Executor) {
	cli, exc := test.NewClient()
	api := test.NewAPI()
	return vm.NewService(cli, api), exc
}
