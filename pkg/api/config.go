package api

import (
	"net/http"
)

type Config struct {
	Host   string
	Port   int
	Path   string
	Secure bool

	PoolingInterval int
	RequestTimeout  int
	HTTPTransport   *http.Transport
}
