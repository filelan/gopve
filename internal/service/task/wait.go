package task

import (
	"fmt"
	"net/http"
	"time"

	"github.com/xabinapal/gopve/pkg/types/task"
)

func (t *Task) Wait() error {
	ch := make(chan error)
	go func(ch chan<- error) {
		defer close(ch)

		var res struct {
			Status task.Status `json:"status"`
		}

		for i := 1; ; i++ {
			err := t.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/tasks/%s/status", t.node, t.upid), nil, &res)
			if err != nil {
				ch <- err
			}

			if res.Status == task.StatusStopped {
				ch <- nil
				return
			}

			time.Sleep(t.svc.poolingInterval)
		}
	}(ch)

	return <-ch
}
