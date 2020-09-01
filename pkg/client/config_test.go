package client_test

import (
	"testing"

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

		if endpoint, err := cfg.Endpoint(); err != nil {
			t.Fatalf("Unexpected client.Config.Endpoint error: %s", err.Error())
		} else if endpoint.String() != parts.URL {
			t.Fatalf("Got endpoint '%s', expected '%s'", endpoint.String(), parts.URL)
		}
	}
}

func TestConfigEndpointGuards(t *testing.T) {
	cfg := client.Config{}
	if _, err := cfg.Endpoint(); err == nil {
		t.Errorf("Got no error, expected 'no host specified'")
	} else if err.Error() != "no host specified" {
		t.Errorf("Got error '%s', expected 'no host specified'", err.Error())
	}
}

func TestConfigEndpointPorts(t *testing.T) {
	options := map[string]urlParts{
		"InsecureDefault": {"http://localhost:80/api2/json/", "localhost", 0, "", false},
		"SecureDefault":   {"https://localhost:8006/api2/json/", "localhost", 0, "", true},
		"InsecureCustom":  {"http://localhost:8080/api2/json/", "localhost", 8080, "", false},
		"SecureCustom":    {"https://localhost:443/api2/json/", "localhost", 443, "", true},
	}

	for n, tt := range options {
		tt := tt
		t.Run(n, helpConfigCheckEndpoint(t, tt))
	}
}

func TestConfigEndpointPaths(t *testing.T) {
	options := map[string]urlParts{
		"InsecureNoPath":   {"http://localhost:80/api2/json/", "localhost", 0, "", false},
		"SecureNoPath":     {"https://localhost:8006/api2/json/", "localhost", 0, "", true},
		"InsecureRootPath": {"http://localhost:80/", "localhost", 0, "/", false},
		"SecureRootPath":   {"https://localhost:8006/", "localhost", 0, "/", true},
		"InsecureSubPath":  {"http://localhost:80/api/", "localhost", 0, "/api/", false},
		"SecureSubPath":    {"https://localhost:8006/api/", "localhost", 0, "/api/", true},
	}

	for n, tt := range options {
		tt := tt
		t.Run(n, helpConfigCheckEndpoint(t, tt))
	}
}
