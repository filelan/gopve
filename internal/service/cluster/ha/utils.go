package ha

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xabinapal/gopve/internal/types"
	"github.com/xabinapal/gopve/pkg/types/cluster"
)

func nodeStringToMap(svc *Service, nodes string) (cluster.HighAvailabilityGroupNodes, error) {
	nodeMap := make(cluster.HighAvailabilityGroupNodes)

	if nodes == "" {
		return nodeMap, nil
	}

	var nodeList types.PVEStringList
	if err := nodeList.Unmarshal(nodes); err != nil {
		return nil, err
	}

	for _, node := range nodeList {
		nodeData := types.PVEStringKV{Separator: ":", AllowNoValue: true}
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

func nodeMapToString(nodes cluster.HighAvailabilityGroupNodes) string {
	if nodes == nil {
		return ""
	}

	var nodeList []string

	for node, priority := range nodes {
		if priority == 0 {
			nodeList = append(nodeList, node.Name())
		} else {
			nodeList = append(nodeList, fmt.Sprintf("%s:%d", node.Name(), priority))
		}
	}

	return strings.Join(nodeList, ",")
}
