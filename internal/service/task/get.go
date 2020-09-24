package task

import (
	"fmt"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types/task"
)

func (svc *Service) List() ([]task.Task, error) {
	return nil, fmt.Errorf("not implemented")
}

func (svc *Service) Get(upid string) (task.Task, error) {
	splits := types.PVEStringList{Separator: ":"}
	if err := splits.Unmarshal(upid); err != nil {
		return nil, err
	}

	if splits.Len() != 9 {
		return nil, task.ErrInvalidUPID
	} else if splits.Elem(0) != "UPID" {
		return nil, task.ErrInvalidUPID
	}

	tElems := splits.List()
	t := NewTask(
		svc,
		tElems[1],
		fmt.Sprintf("%s:%s:%s", tElems[2], tElems[3], tElems[4]),
		tElems[5],
		tElems[6],
		tElems[7],
		tElems[8],
	)

	switch t.action {
	case task.ActionQMCreate, task.ActionVZCreate:
		return NewVirtualMachineTask(t)

	default:
		return t, nil
	}
}
