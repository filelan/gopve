package utils

import (
	"net/url"

	"github.com/xabinapal/gopve/pkg/types"
)

type Client interface {
	Request(method string, resource string, form url.Values, out interface{}) error

	WaitForTask(upid string) types.Task

	Lock()
	Unlock()
}
