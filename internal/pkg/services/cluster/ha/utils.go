package ha

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xabinapal/gopve/pkg/types"
)

func nodeStringToMap(svc *Service, nodes string) (types.HighAvailabilityGroupNodes, error) {
	nodeMap := make(types.HighAvailabilityGroupNodes)
	if nodes == "" {
		return nodeMap, nil
	}

	nodeList := strings.Split(nodes, ",")
	for _, node := range nodeList {
		nodeData := strings.Split(node, ":")

		key := HighAvailabilityGroupNode{svc, nodeData[0]}
		if len(nodeData) == 1 {
			nodeMap[key] = 0
		} else {
			priority, err := strconv.Atoi(nodeData[1])
			if err != nil {
				return nil, err
			}
			nodeMap[key] = uint(priority)
		}
	}

	return nodeMap, nil
}

func nodeMapToString(nodes types.HighAvailabilityGroupNodes) string {
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
