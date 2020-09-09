package test

import "github.com/xabinapal/gopve/internal/service/task"

func NewTask(node string, upid string) *task.Task {
	return task.NewTask(nil, node, upid)
}
