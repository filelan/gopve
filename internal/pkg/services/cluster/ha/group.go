package ha

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/internal/pkg/utils"

	"github.com/xabinapal/gopve/pkg/types"
)

type getGroupResponseJSON struct {
	Name        string        `json:"group"`
	Description string        `json:"comment"`
	Restricted  utils.PVEBool `json:"restricted"`
	NoFailback  utils.PVEBool `json:"nofailback"`
	Nodes       string        `json:"nodes"`
}

func (res getGroupResponseJSON) ConvertToEntity(svc *Service, full bool) (types.HighAvailabilityGroup, error) {
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

func (svc *Service) ListGroups() ([]types.HighAvailabilityGroup, error) {
	var res []getGroupResponseJSON
	if err := svc.Client.Request(http.MethodGet, "cluster/ha/groups", nil, &res); err != nil {
		return nil, err
	}

	groups := make([]types.HighAvailabilityGroup, len(res))
	for i, group := range res {
		out, err := group.ConvertToEntity(svc, false)
		if err != nil {
			return nil, err
		}

		groups[i] = out
	}

	return groups, nil
}

func (svc *Service) GetGroup(name string) (types.HighAvailabilityGroup, error) {
	var res getGroupResponseJSON
	if err := svc.Client.Request(http.MethodGet, fmt.Sprintf("cluster/ha/groups/%s", name), nil, &res); err != nil {
		return nil, err
	}

	return res.ConvertToEntity(svc, true)
}

func (svc *Service) CreateGroup(name string, props types.HighAvailabilityGroupProperties, nodes types.HighAvailabilityGroupNodes) (types.HighAvailabilityGroup, error) {
	if len(nodes) == 0 {
		return &HighAvailabilityGroup{}, fmt.Errorf("at least one node is required")
	}

	var form utils.RequestValues
	form.AddString("type", "group")
	form.AddString("group", name)

	form.AddString("comment", props.Description)
	form.AddBool("restricted", props.RestrictedResourceExecution)
	form.AddBool("nofailback", !props.MigrateResourcesToHigherPriority)

	form.AddString("nodes", nodeMapToString(nodes))

	if err := svc.Client.Request(http.MethodPost, "cluster/ha/groups", form, nil); err != nil {
		return &HighAvailabilityGroup{}, err
	}

	return svc.GetGroup(name)
}
