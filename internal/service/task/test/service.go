package test

import (
	"github.com/xabinapal/gopve/internal/service/task"
	"github.com/xabinapal/gopve/pkg/client/test"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func NewService() (*task.Service, *test.API, *mocks.Executor) {
	cli, exc := test.NewClient()
	api := test.NewAPI()
	return task.NewService(cli, api, 0), api, exc
}
