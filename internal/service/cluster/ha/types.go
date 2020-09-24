package ha

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/internal/client"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/types/cluster"
	"github.com/xabinapal/gopve/pkg/types/node"
)

type Service struct {
	client client.Client
	api    client.API
}

func NewService(cli client.Client, api client.API) *Service {
	return &Service{
		client: cli,
		api:    api,
	}
}

type HighAvailabilityGroup struct {
	svc  *Service
	full bool

	name        string
	description string

	restrictedResourceExecution      bool
	migrateResourcesToHigherPriority bool

	nodes cluster.HighAvailabilityGroupNodes
}

func (obj *HighAvailabilityGroup) Load() error {
	if obj.full {
		return nil
	}

	group, err := obj.svc.GetGroup(obj.name)
	if err != nil {
		return nil
	}

	switch x := group.(type) {
	case *HighAvailabilityGroup:
		*obj = *x
	default:
		panic("This should never happen")
	}

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

func (obj *HighAvailabilityGroup) Nodes() (cluster.HighAvailabilityGroupNodes, error) {
	if err := obj.Load(); err != nil {
		return nil, err
	}

	return obj.nodes, nil
}

func (obj *HighAvailabilityGroup) GetProperties() (cluster.HighAvailabilityGroupProperties, error) {
	if err := obj.Load(); err != nil {
		return cluster.HighAvailabilityGroupProperties{}, err
	}

	return cluster.HighAvailabilityGroupProperties{
		Description:                      obj.description,
		RestrictedResourceExecution:      obj.restrictedResourceExecution,
		MigrateResourcesToHigherPriority: obj.migrateResourcesToHigherPriority,
	}, nil
}

func (obj *HighAvailabilityGroup) SetProperties(props cluster.HighAvailabilityGroupProperties) error {
	var form request.Values

	form.AddString("comment", props.Description)
	form.AddBool("restricted", props.RestrictedResourceExecution)
	form.AddBool("nofailback", !props.MigrateResourcesToHigherPriority)

	if err := obj.svc.client.Request(http.MethodPut, fmt.Sprintf("cluster/ha/groups/%s", obj.name), form, nil); err != nil {
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

	var nodeMap cluster.HighAvailabilityGroupNodes
	for node, priority := range obj.nodes {
		nodeMap[node] = priority
	}

	for node, priority := range nodes {
		nodeMap[HighAvailabilityGroupNode{obj.svc, node}] = priority
	}

	if err := obj.svc.client.Request(http.MethodPut, fmt.Sprintf("cluster/ha/groups/%s", obj.name), request.Values{
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

	var nodeMap cluster.HighAvailabilityGroupNodes

	nodeKeys := make(map[string]*cluster.HighAvailabilityGroupNode)

	for n, priority := range obj.nodes {
		node := n
		nodeMap[node] = priority
		nodeKeys[node.Name()] = &node
	}

	for _, node := range nodes {
		n, ok := nodeKeys[node]
		if ok {
			delete(nodeMap, *n)
		}
	}

	if err := obj.svc.client.Request(http.MethodPut, fmt.Sprintf("cluster/ha/groups/%s", obj.name), request.Values{
		"nodes": {nodeMapToString(nodeMap)},
	}, nil); err != nil {
		return err
	}

	obj.nodes = nodeMap

	return nil
}

func (obj *HighAvailabilityGroup) Delete() error {
	return obj.svc.client.Request(http.MethodDelete, fmt.Sprintf("cluster/ha/groups/%s", obj.name), nil, nil)
}

type HighAvailabilityGroupNode struct {
	svc  *Service
	name string
}

func (obj HighAvailabilityGroupNode) Name() string {
	return obj.name
}

func (obj HighAvailabilityGroupNode) Get() (node.Node, error) {
	return obj.svc.api.Node().Get(obj.name)
}
