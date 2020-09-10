package ha

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/cluster"
)

type getGroupResponseJSON struct {
	Name        string        `json:"group"`
	Description string        `json:"comment"`
	Restricted  types.PVEBool `json:"restricted"`
	NoFailback  types.PVEBool `json:"nofailback"`
	Nodes       string        `json:"nodes"`
}

func (res getGroupResponseJSON) Map(svc *Service, full bool) (cluster.HighAvailabilityGroup, error) {
	if !full {
		return &HighAvailabilityGroup{
			svc:  svc,
			full: false,

			name: res.Name,
		}, nil
	}

	nodes, err := nodeStringToMap(svc, res.Nodes)
	if err != nil {
		return nil, err
	}

	return &HighAvailabilityGroup{
		svc:  svc,
		full: true,

		name:                             res.Name,
		description:                      res.Description,
		restrictedResourceExecution:      res.Restricted.Bool(),
		migrateResourcesToHigherPriority: !res.NoFailback.Bool(),
		nodes:                            nodes,
	}, nil
}

func (svc *Service) ListGroups() ([]cluster.HighAvailabilityGroup, error) {
	var res []getGroupResponseJSON
	if err := svc.client.Request(http.MethodGet, "cluster/ha/groups", nil, &res); err != nil {
		return nil, err
	}

	groups := make([]cluster.HighAvailabilityGroup, len(res))

	for i, group := range res {
		out, err := group.Map(svc, false)
		if err != nil {
			return nil, err
		}

		groups[i] = out
	}

	return groups, nil
}

func (svc *Service) GetGroup(name string) (cluster.HighAvailabilityGroup, error) {
	var res getGroupResponseJSON
	if err := svc.client.Request(http.MethodGet, fmt.Sprintf("cluster/ha/groups/%s", name), nil, &res); err != nil {
		return nil, err
	}

	return res.Map(svc, true)
}

func (svc *Service) CreateGroup(name string, props cluster.HighAvailabilityGroupProperties, nodes cluster.HighAvailabilityGroupNodes) (cluster.HighAvailabilityGroup, error) {
	if len(nodes) == 0 {
		return nil, fmt.Errorf("at least one node is required")
	}

	var form request.Values

	form.AddString("type", "group")
	form.AddString("group", name)

	form.AddString("comment", props.Description)
	form.AddBool("restricted", props.RestrictedResourceExecution)
	form.AddBool("nofailback", !props.MigrateResourcesToHigherPriority)

	form.AddString("nodes", nodeMapToString(nodes))

	if err := svc.client.Request(http.MethodPost, "cluster/ha/groups", form, nil); err != nil {
		return nil, err
	}

	return svc.GetGroup(name)
}
