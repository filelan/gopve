package node

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/types"
)

func (node *Node) Version() (types.Version, error) {
	var res types.Version
	err := node.svc.client.Request(http.MethodGet, fmt.Sprintf("nodes/%s/version", node.name), nil, &res)
	return res, err
}
