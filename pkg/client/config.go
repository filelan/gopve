package client

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/xabinapal/gopve/pkg/request"
)

type Config struct {
	Host   string
	Port   int
	Path   string
	Secure bool

	HTTPTransport   *http.Transport
	RequestTimeout  time.Duration
	PoolingInterval time.Duration

	Executor request.Executor
}

func (cfg Config) Endpoint() (*url.URL, error) {
	scheme, port := cfg.getURLParts()
	return cfg.getURL(scheme, port)
}

func (cfg Config) getURLParts() (scheme string, port int) {
	if cfg.Port != 0 {
		port = cfg.Port
	}

	switch cfg.Secure {
	case true:
		scheme = "https"
		if cfg.Port == 0 {
			port = 80
		}
	case false:
		scheme = "http"
		if cfg.Port == 0 {
			port = 443
		}
	}

	return
}

func (cfg Config) getURL(scheme string, port int) (*url.URL, error) {
	host, err := url.Parse(fmt.Sprintf("%s://%s:%d", scheme, cfg.Host, port))
	if err != nil {
		return nil, err
	}

	base := strings.Trim(cfg.Path, "/")
	if base == "" {
		host.Path = "/api2/json"
		return host, nil
	} else {
		path, err := url.Parse(fmt.Sprintf("/%s/api2/json/", base))
		if err != nil {
			return nil, err
		}

		return host.ResolveReference(path), nil
	}

}
