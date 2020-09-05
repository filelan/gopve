package test

import (
	"github.com/xabinapal/gopve/internal/service/node"
	"github.com/xabinapal/gopve/pkg/client/test"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func NewService() (*node.Service, *mocks.Executor) {
	cli, exc := test.NewClient()
	api := test.NewAPI()
	return node.NewService(cli, api), exc
}
