package node

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/node"
)

type getResponseJSON struct {
	Name   string      `json:"node"`
	Status node.Status `json:"status"`
}

func (res getResponseJSON) Map(svc *Service) (node.Node, error) {
	if err := res.Status.IsValid(); err != nil {
		return nil, fmt.Errorf("unsupported node status")
	}

	return NewNode(svc, res.Name, res.Status), nil
}

func (svc *Service) List() ([]node.Node, error) {
	var res []getResponseJSON
	if err := svc.client.Request(http.MethodGet, "cluster/resources", request.Values{
		"type": {"node"},
	}, &res); err != nil {
		return nil, err
	}

	nodes := make([]node.Node, len(res))

	for i, node := range res {
		out, err := node.Map(svc)
		if err != nil {
			return nil, err
		}

		nodes[i] = out
	}

	return nodes, nil
}

func (svc *Service) Get(name string) (node.Node, error) {
	nodes, err := svc.List()
	if err != nil {
		return nil, err
	}

	for _, node := range nodes {
		if node.Name() == name {
			return node, nil
		}
	}

	return nil, fmt.Errorf("node not found")
}
