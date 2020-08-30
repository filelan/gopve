package service

import "github.com/xabinapal/gopve/pkg/types/task"

type Task interface {
	List() ([]task.Task, error)
	Get(upid string) (task.Task, error)
}
