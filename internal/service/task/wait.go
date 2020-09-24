package task

import (
	"time"

	"github.com/xabinapal/gopve/pkg/types/task"
)

func (t *Task) Wait() error {
	ch := make(chan error)
	go func(ch chan<- error) {
		defer close(ch)

		for i := 1; ; i++ {
			status, err := t.GetStatus()
			if err != nil {
				ch <- err
			}

			if status == task.StatusStopped {
				ch <- nil
				return
			}

			time.Sleep(t.svc.poolingInterval)
		}
	}(ch)

	return <-ch
}
