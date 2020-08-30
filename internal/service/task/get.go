package task

import (
	"fmt"
	"strings"

	"github.com/xabinapal/gopve/pkg/types/task"
)

func (svc *Service) List() ([]task.Task, error) {
	return nil, fmt.Errorf("not implemented")
}


func (svc *Service) Get(upid string) (task.Task, error) {
	splits := strings.SplitN(upid, ":", 3)
	if len(splits) < 3 {
		return nil, fmt.Errorf("invalid UPID")
	}

	node := splits[1]

	return &Task{
		node: node,
		upid: upid,
	}, nil
}
