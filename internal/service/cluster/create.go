package cluster

import (
	"net/http"

	"github.com/xabinapal/gopve/pkg/types/cluster"
	"github.com/xabinapal/gopve/pkg/types/task"
)

func (svc *Service) Create(
	name string,
	props cluster.NodeProperties,
) (task.Task, error) {
	c, err := svc.Get()
	if err != nil {
		return nil, err
	}

	if c.Mode() == cluster.ModeCluster {
		return nil, cluster.ErrAlreadyInCluster
	}

	form, err := props.MapToValues()
	if err != nil {
		return nil, err
	}

	form.AddString("clustername", name)

	var task string
	if err := svc.client.Request(http.MethodPost, "cluster/config", form, &task); err != nil {
		return nil, err
	}

	return svc.api.Task().Get(task)
}

func (svc *Service) Join(
	hostname, password, fingerprint string,
	props cluster.NodeProperties,
) (task.Task, error) {
	c, err := svc.Get()
	if err != nil {
		return nil, err
	}

	if c.Mode() == cluster.ModeCluster {
		return nil, cluster.ErrAlreadyInCluster
	}

	form, err := props.MapToValues()
	if err != nil {
		return nil, err
	}

	form.AddString("hostname", hostname)
	form.AddString("password", password)
	form.AddString("fingerprint", fingerprint)

	var task string
	if err := svc.client.Request(http.MethodPost, "cluster/config/join", form, &task); err != nil {
		return nil, err
	}

	return svc.api.Task().Get(task)
}
