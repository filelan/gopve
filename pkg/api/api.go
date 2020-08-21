package api

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/xabinapal/gopve/pkg/types"
)

type API struct {
	client *client
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

type getVersionResponseJSON struct {
	Release string `json:"release"`
	Version string `json:"version"`
	RepoID  string `json:"repoid"`
}

func (res getVersionResponseJSON) ConvertToEntity() *types.Version {
	return &types.Version{
		Release: res.Release,
		Version: res.Version,
		RepoID:  res.RepoID,
	}
}

func (api *API) Version() (*types.Version, error) {
	var out getVersionResponseJSON
	err := api.client.Request(http.MethodGet, "/version", nil, &out)
	if err != nil {
		return nil, err
	}
	return out.ConvertToEntity(), nil
}
