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
	splits := strings.Split(upid, ":")
	if len(splits) != 9 {
		return nil, task.ErrInvalidUPID
	} else if splits[0] != "UPID" {
		return nil, task.ErrInvalidUPID
	}

	t := NewTask(svc, splits[1], fmt.Sprintf("%s:%s:%s", splits[2], splits[3], splits[4]), splits[5], splits[6], splits[7], splits[8])

	switch t.action {
	case task.ActionQMCreate, task.ActionVZCreate:
		return NewVirtualMachineTask(t)

	default:
		return t, nil
	}
}
