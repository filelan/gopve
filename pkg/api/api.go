package api

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/xabinapal/gopve/internal/pkg/services/cluster"
	"github.com/xabinapal/gopve/internal/pkg/services/node"
	"github.com/xabinapal/gopve/internal/pkg/services/pool"
	"github.com/xabinapal/gopve/internal/pkg/services/vm"
	"github.com/xabinapal/gopve/pkg/types"
)

type API struct {
	client *client

	cluster *cluster.Service
	node    *node.Service
	pool    *pool.Service
	vm      *vm.Service
}

func New(cfg Config) (*API, error) {
	transport := cfg.HTTPTransport
	if transport == nil {
		transport = &http.Transport{}
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	httpClient := http.Client{
		Transport: transport,
		Timeout:   time.Duration(cfg.RequestTimeout) * time.Second,
		Jar:       jar,
	}

	var scheme string
	if cfg.Secure {
		scheme = "https"
	} else {
		scheme = "http"
	}
	host, err := url.Parse(fmt.Sprintf("%s://%s:%d", scheme, cfg.Host, cfg.Port))
	if err != nil {
		return nil, err
	}

	path, err := url.Parse(fmt.Sprintf("%s/api2/json/", cfg.Path))
	if err != nil {
		return nil, err
	}

	url := host.ResolveReference(path)

	return &API{
		client: &client{
			client: &httpClient,
			url:    url,

			poolingInterval: cfg.PoolingInterval,
		},
	}, nil
}

func (api *API) Cluster() types.ClusterService {
	if api.cluster == nil {
		api.cluster = cluster.NewService(api, api.client)
	}

	return api.cluster
}

func (api *API) Node() types.NodeService {
	if api.node == nil {
		api.node = node.NewService(api, api.client)
	}

	return api.node
}

func (api *API) Pool() types.PoolService {
	if api.pool == nil {
		api.pool = pool.NewService(api, api.client)
	}

	return api.pool
}

func (api *API) VM() types.VMService {
	if api.vm == nil {
		api.vm = vm.NewService(api, api.client)
	}

	return api.vm
}
