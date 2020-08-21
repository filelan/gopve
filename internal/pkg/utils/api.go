package utils

import "github.com/xabinapal/gopve/pkg/types"

type API interface {
	Node() types.NodeService
}
