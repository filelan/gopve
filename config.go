package gopve

import (
	"fmt"
	"strings"
)

type Config struct {
	Schema      string
	Host        string
	Port        uint32
	User        string
	Password    string
	InvalidCert string
}

func (cfg *Config) GenerateRootURI() (string, err) {
	schema := strings.ToLower(cfg.Schema)
	if schema != "http" && schema != "https" {
		return nil, "Invalid schema"
	}

	if port == 0 {
		return nil, "Invalid port"
	}

	return fmt.Sprintf("%s://%s:%d/api2/json/", cfg.Schema, cfg.Host, cfg.Port)
}
