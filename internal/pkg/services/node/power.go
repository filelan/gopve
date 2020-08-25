package node

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/internal/pkg/utils"
	"github.com/xabinapal/gopve/pkg/types"
)

func postStatus(node *Node, command string) error {
	if err := node.svc.Client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/status", node.name), utils.RequestValues{
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

func (node *Node) WakeOnLAN() (types.Task, error) {
	var task string
	if err := node.svc.Client.Request(http.MethodPost, fmt.Sprintf("nodes/%s/wakeonlan", node.Name), nil, &task); err != nil {
		return nil, err
	}

	return node.svc.Client.WaitForTask(task), nil
}
