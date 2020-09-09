package client

import (
	"net/http"

	"github.com/xabinapal/gopve/pkg/types"
)

func (cli *Client) Version() (*types.Version, error) {
	var res types.Version
	return &res, cli.Request(http.MethodGet, "/version", nil, &res)
}
