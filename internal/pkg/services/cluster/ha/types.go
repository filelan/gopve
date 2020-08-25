package ha

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/internal/pkg/utils"
	"github.com/xabinapal/gopve/pkg/types"
)

type Service struct {
	utils.Service
}

func NewService(api utils.API, client utils.Client) *Service {
	return &Service{
		utils.Service{api, client},
	}
}

type HighAvailabilityGroup struct {
	svc  *Service
	full bool

	name        string
	description string

	restrictedResourceExecution      bool
	migrateResourcesToHigherPriority bool

	nodes types.HighAvailabilityGroupNodes
}

func (obj *HighAvailabilityGroup) Load() error {
	if obj.full {
		return nil
	}

	group, err := obj.svc.GetGroup(obj.name)
	if err != nil {
		return nil
	}

	obj.description, _ = group.Description()
	obj.restrictedResourceExecution, _ = group.RestrictedResourceExecution()
	obj.migrateResourcesToHigherPriority, _ = group.MigrateResourcesToHigherPriority()

	obj.nodes, _ = group.Nodes()

	return nil
}

func (obj *HighAvailabilityGroup) Name() string {
	return obj.name
}

func (obj *HighAvailabilityGroup) Description() (string, error) {
	if err := obj.Load(); err != nil {
		return "", err
	}

	return obj.description, nil
}

func (obj *HighAvailabilityGroup) RestrictedResourceExecution() (bool, error) {
	if err := obj.Load(); err != nil {
		return false, err
	}

	return obj.restrictedResourceExecution, nil
}

func (obj *HighAvailabilityGroup) MigrateResourcesToHigherPriority() (bool, error) {
	if err := obj.Load(); err != nil {
		return false, err
	}

	return obj.migrateResourcesToHigherPriority, nil
}

func (obj *HighAvailabilityGroup) Nodes() (types.HighAvailabilityGroupNodes, error) {
	if err := obj.Load(); err != nil {
		return nil, err
	}

	return obj.nodes, nil
}

func (obj *HighAvailabilityGroup) GetProperties() (types.HighAvailabilityGroupProperties, error) {
	if err := obj.Load(); err != nil {
		return types.HighAvailabilityGroupProperties{}, err
	}

	return types.HighAvailabilityGroupProperties{
		Description:                      obj.description,
		RestrictedResourceExecution:      obj.restrictedResourceExecution,
		MigrateResourcesToHigherPriority: obj.migrateResourcesToHigherPriority,
	}, nil
}

func (obj *HighAvailabilityGroup) SetProperties(props types.HighAvailabilityGroupProperties) error {
	var form utils.RequestValues
	form.AddString("comment", props.Description)
	form.AddBool("restricted", props.RestrictedResourceExecution)
	form.AddBool("nofailback", !props.MigrateResourcesToHigherPriority)

	if err := obj.svc.Client.Request(http.MethodPut, fmt.Sprintf("cluster/ha/groups/%s", obj.name), form, nil); err != nil {
		return err
	}

	obj.description = props.Description
	obj.restrictedResourceExecution = props.RestrictedResourceExecution
	obj.migrateResourcesToHigherPriority = props.MigrateResourcesToHigherPriority

	return nil
}

func (obj *HighAvailabilityGroup) AddNodes(nodes map[string]uint) error {
	if err := obj.Load(); err != nil {
		return err
	}

	var nodeMap types.HighAvailabilityGroupNodes
	for node, priority := range obj.nodes {
		nodeMap[node] = priority
	}

	for node, priority := range nodes {
		nodeMap[HighAvailabilityGroupNode{obj.svc, node}] = priority
	}

	if err := obj.svc.Client.Request(http.MethodPut, fmt.Sprintf("cluster/ha/groups/%s", obj.name), utils.RequestValues{
		"nodes": {nodeMapToString(nodeMap)},
	}, nil); err != nil {
		return err
	}

	obj.nodes = nodeMap

	return nil
}

func (obj *HighAvailabilityGroup) DeleteNodes(nodes []string) error {
	if err := obj.Load(); err != nil {
		return err
	}

	var nodeMap types.HighAvailabilityGroupNodes
	nodeKeys := make(map[string]*types.HighAvailabilityGroupNode)
	for node, priority := range obj.nodes {
		nodeMap[node] = priority
		nodeKeys[node.Name()] = &node
	}

	for _, node := range nodes {
		n, ok := nodeKeys[node]
		if ok {
			delete(nodeMap, *n)
		}
	}

	if err := obj.svc.Client.Request(http.MethodPut, fmt.Sprintf("cluster/ha/groups/%s", obj.name), utils.RequestValues{
		"nodes": {nodeMapToString(nodeMap)},
	}, nil); err != nil {
		return err
	}

	obj.nodes = nodeMap

	return nil
}

func (obj *HighAvailabilityGroup) Delete() error {
	return obj.svc.Client.Request(http.MethodDelete, fmt.Sprintf("cluster/ha/groups/%s", obj.name), nil, nil)
}

type HighAvailabilityGroupNode struct {
	svc  *Service
	name string
}

func (obj HighAvailabilityGroupNode) Name() string {
	return obj.name
}

func (obj HighAvailabilityGroupNode) Get() (types.Node, error) {
	return obj.svc.API.Node().Get(obj.name)
}
