package ha

import (
	"fmt"
	"strconv"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types/cluster"
)

func nodeStringToMap(
	svc *Service,
	nodes string,
) (cluster.HighAvailabilityGroupNodes, error) {
	nodeMap := make(cluster.HighAvailabilityGroupNodes)

	if nodes == "" {
		return nodeMap, nil
	}

	nodeList := types.PVEList{Separator: ","}
	if err := nodeList.Unmarshal(nodes); err != nil {
		return nil, err
	}

	for _, node := range nodeList.List() {
		nodeData := types.PVEKeyValue{Separator: ":", AllowNoValue: true}
		if err := nodeData.Unmarshal(node); err != nil {
			return nil, err
		}

		key := HighAvailabilityGroupNode{svc, nodeData.Key()}

		if nodeData.HasValue() {
			priority, err := strconv.Atoi(nodeData.Value())
			if err != nil {
				return nil, err
			}

			nodeMap[key] = uint(priority)
		} else {
			nodeMap[key] = 0
		}
	}

	return nodeMap, nil
}

func nodeMapToList(
	nodes cluster.HighAvailabilityGroupNodes,
) types.PVEList {
	nodeList := types.PVEList{Separator: ","}

	if nodes == nil {
		return nodeList
	}

	for node, priority := range nodes {
		if priority == 0 {
			nodeList.Append(node.Name())
		} else {
			nodeList.Append(fmt.Sprintf("%s:%d", node.Name(), priority))
		}
	}

	return nodeList
}
