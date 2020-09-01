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
	Port   uint16
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

func (cfg Config) getURLParts() (scheme string, port uint16) {
	if cfg.Port != 0 {
		port = cfg.Port
	}

	switch cfg.Secure {
	case true:
		scheme = "https"

		if cfg.Port == 0 {
			port = 8006
		}

	case false:
		scheme = "http"

		if cfg.Port == 0 {
			port = 80
		}
	}

	return
}

func (cfg Config) getURL(scheme string, port uint16) (*url.URL, error) {
	host := strings.Trim(cfg.Host, " ")
	if host == "" {
		return nil, fmt.Errorf("no host specified")
	}

	absoluteURL, err := url.Parse(fmt.Sprintf("%s://%s:%d/", scheme, cfg.Host, port))
	if err != nil {
		return nil, err
	}

	basePath := cfg.getNormalizedPath()
	if basePath == "" {
		absoluteURL.Path = "api2/json/"

		return absoluteURL, nil
	}

	path, err := url.Parse(basePath)
	if err != nil {
		return nil, err
	}

	return absoluteURL.ResolveReference(path), nil
}

func (cfg Config) getNormalizedPath() string {
	if cfg.Path == "" {
		return cfg.Path
	}

	basePath := strings.Trim(cfg.Path, "/")

	return fmt.Sprintf("%s/", basePath)
}
