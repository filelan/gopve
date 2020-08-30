package client

import (
	"github.com/xabinapal/gopve/pkg/request"
)

type Client interface {
	Request(method string, resource string, form request.Values, out interface{}) error
	StartAtomicBlock()
	EndAtomicBlock()
}
