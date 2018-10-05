package gopve

import (
	"errors"
	"fmt"
	"strings"
)

type Config struct {
	Schema      string
	Host        string
	Port        uint32
	User        string
	Password    string
	InvalidCert bool
}

func (cfg *Config) GenerateRootURI() (string, error) {
	schema := strings.ToLower(cfg.Schema)
	if schema != "http" && schema != "https" {
		return "", errors.New("Invalid schema")
	}

	if cfg.Port == 0 {
		return "", errors.New("Invalid port")
	}

	return fmt.Sprintf("%s://%s:%d/api2/json/", cfg.Schema, cfg.Host, cfg.Port), nil
}
