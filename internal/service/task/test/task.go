package test

import (
	"github.com/xabinapal/gopve/internal/service/task"
	"github.com/xabinapal/gopve/pkg/client/test"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func NewTask(node, uuid, action, id, user, extra string) (*task.Task, *test.API, *mocks.Executor) {
	svc, api, exc := NewService()
	return task.NewTask(svc, node, uuid, action, id, user, extra), api, exc
}
