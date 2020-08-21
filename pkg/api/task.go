package api

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/xabinapal/gopve/internal/pkg/utils"
	"github.com/xabinapal/gopve/pkg/types"
)

const defaultPoolingInterval = 60

func (api *API) GetTask(upid string) (types.Task, error) {
	splits := strings.SplitN(upid, ":", 3)
	if len(splits) < 3 {
		return nil, fmt.Errorf("invalid UPID")
	}

	node := splits[1]

	interval := api.client.poolingInterval
	if interval <= 0 {
		interval = defaultPoolingInterval
	}

	return api.client.waitForTask(node, upid, interval), nil
}

func (c *client) WaitForTask(upid string) types.Task {
	node := strings.SplitN(upid, ":", 3)[1]
	if c.poolingInterval > 0 {
		return c.waitForTask(node, upid, c.poolingInterval)
	}

	return nil
}

func (c *client) waitForTask(node string, upid string, interval int) types.Task {
	ch := make(chan error)
	go func(ch chan<- error) {
		defer close(ch)

		var res struct {
			Status types.TaskStatus `json:"status"`
		}

		for i := 1; ; i++ {
			err := c.Request(http.MethodGet, fmt.Sprintf("nodes/%s/tasks/%s/status", node, upid), nil, &res)
			if err != nil {
				ch <- err
			}

			if res.Status == types.TaskStopped {
				ch <- nil
				return
			}

			time.Sleep(time.Duration(interval) * time.Second)
		}
	}(ch)

	return utils.NewPVETask(node, upid, ch)
}
