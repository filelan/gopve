package client_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/client"
)

type urlParts struct {
	URL    string
	Host   string
	Port   uint16
	Path   string
	Secure bool
}

func helpConfigCheckEndpoint(t *testing.T, parts urlParts) func(t *testing.T) {
	t.Helper()

	return func(t *testing.T) {
		cfg := client.Config{
			Host:   parts.Host,
			Port:   parts.Port,
			Path:   parts.Path,
			Secure: parts.Secure,
		}

		endpoint, err := cfg.Endpoint()
		assert.NoError(t, err)
		assert.Equal(t, parts.URL, endpoint.String())
	}
}

func TestConfigEndpointGuards(t *testing.T) {
	cfg := client.Config{}
	_, err := cfg.Endpoint()
	require.Error(t, err)
	require.Equal(t, client.ErrConfigNoHost, err)
}

func TestConfigEndpointPorts(t *testing.T) {
	options := map[string]urlParts{
		"InsecureDefault": {
			"http://localhost:80/api2/json/",
			"localhost",
			0,
			"",
			false,
		},
		"SecureDefault": {
			"https://localhost:8006/api2/json/",
			"localhost",
			0,
			"",
			true,
		},
		"InsecureCustom": {
			"http://localhost:8080/api2/json/",
			"localhost",
			8080,
			"",
			false,
		},
		"SecureCustom": {
			"https://localhost:443/api2/json/",
			"localhost",
			443,
			"",
			true,
		},
	}

	for n, tt := range options {
		tt := tt
		t.Run(n, helpConfigCheckEndpoint(t, tt))
	}
}

func TestConfigEndpointPaths(t *testing.T) {
	options := map[string]urlParts{
		"InsecureNoPath": {
			"http://localhost:80/api2/json/",
			"localhost",
			0,
			"",
			false,
		},
		"SecureNoPath": {
			"https://localhost:8006/api2/json/",
			"localhost",
			0,
			"",
			true,
		},
		"InsecureRootPath": {
			"http://localhost:80/",
			"localhost",
			0,
			"/",
			false,
		},
		"SecureRootPath": {
			"https://localhost:8006/",
			"localhost",
			0,
			"/",
			true,
		},
		"InsecureSubPath": {
			"http://localhost:80/api/",
			"localhost",
			0,
			"/api/",
			false,
		},
		"SecureSubPath": {
			"https://localhost:8006/api/",
			"localhost",
			0,
			"/api/",
			true,
		},
	}

	for n, tt := range options {
		tt := tt
		t.Run(n, helpConfigCheckEndpoint(t, tt))
	}
}
