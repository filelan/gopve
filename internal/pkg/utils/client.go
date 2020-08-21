package utils

import (
	"net/url"
)

type Client interface {
	Request(method string, resource string, form url.Values, out interface{}) error

	Lock()
	Unlock()
}
