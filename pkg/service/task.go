package service

import "github.com/xabinapal/gopve/pkg/types/task"

//go:generate mockery --case snake --name Task

type Task interface {
	List() ([]task.Task, error)
	Get(upid string) (task.Task, error)
}
