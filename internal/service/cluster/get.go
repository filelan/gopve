package cluster

import (
	"errors"
	"net/http"

	"github.com/xabinapal/gopve/pkg/types/cluster"
)

type getResponseJSON struct {
	Totem struct {
		Name string `json:"cluster_name"`
	} `json:"totem"`
}

func (svc *Service) Get() (cluster.Cluster, error) {
	var res getResponseJSON
	if err := svc.client.Request(http.MethodGet, "cluster/config/join", nil, &res); err != nil {
		if errors.Is(cluster.ErrNotInCluster, err) {
			return NewCluster(svc, cluster.ModeStandalone, ""), nil
		}

		return nil, err
	}

	return NewCluster(svc, cluster.ModeCluster, res.Totem.Name), nil
}
