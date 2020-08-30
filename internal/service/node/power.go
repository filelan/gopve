package node

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/task"
)

func postStatus(node *Node, command string) error {
	if err := node.svc.client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/status", node.name), request.Values{
		"command": {command},
	}, nil); err != nil {
		return err
	}

	return nil
}

func (node *Node) Shutdown() error {
	return postStatus(node, "shutdown")
}

func (node *Node) Reboot() error {
	return postStatus(node, "reboot")
}

func (node *Node) WakeOnLAN() (task.Task, error) {
	var task string
	if err := node.svc.client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/wakeonlan", node.name), nil, &task); err != nil {
		return nil, err
	}

	return node.svc.api.Task().Get(task)
}
