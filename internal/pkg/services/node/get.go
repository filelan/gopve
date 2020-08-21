package node

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/xabinapal/gopve/pkg/types"
)

type getResponseJSON struct {
	Name   string           `json:"node"`
	Status types.NodeStatus `json:"status"`
}

func (res getResponseJSON) ConvertToEntity(svc *Service) (*Node, error) {
	if err := res.Status.IsValid(); err != nil {
		return nil, fmt.Errorf("unsupported node status")
	}

	return &Node{
		svc: svc,

		name:   res.Name,
		status: res.Status,
	}, nil
}

func (svc *Service) GetAll() ([]types.Node, error) {
	var res []getResponseJSON
	if err := svc.Client.Request(http.MethodGet, "cluster/resources", url.Values{
		"type": {"node"},
	}, &res); err != nil {
		return nil, err
	}

	nodes := make([]types.Node, len(res))
	for i, node := range res {
		out, err := node.ConvertToEntity(svc)
		if err != nil {
			return nil, err
		}

		nodes[i] = out
	}

	return nodes, nil
}

func (svc *Service) Get(name string) (types.Node, error) {
	var res []getResponseJSON
	if err := svc.Client.Request(http.MethodGet, "cluster/resources", url.Values{
		"type": {"node"},
	}, &res); err != nil {
		return nil, err
	}

	for _, node := range res {
		if node.Name == name {
			out, err := node.ConvertToEntity(svc)
			if err != nil {
				return nil, err
			}
			return out, nil
		}
	}

	return nil, fmt.Errorf("node not found")
}
