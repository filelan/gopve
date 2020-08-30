package client

import (
	"net/http"

	"github.com/xabinapal/gopve/pkg/types"
)

func (cli *Client) Version() (*types.Version, error) {
	var res types.Version
	err := cli.Request(http.MethodGet, "/version", nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
