package test

import (
	"github.com/xabinapal/gopve/pkg/client"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func NewClient() (*client.Client, *mocks.Executor) {
	exc := new(mocks.Executor)
	cli := client.NewClientWithExecutor(exc, 0)
	return cli, exc
}
