package task

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/types/task"
)

type getStatusResponseJSON struct {
	Status task.Status `json:"status"`
}

func (t *Task) GetStatus() (task.Status, error) {
	var res getStatusResponseJSON
	err := t.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/tasks/%s/status", t.node, t.upid), nil, &res)
	if err != nil {
		return task.StatusStopped, err
	}

	return res.Status, nil
}
