package utils

import (
	"net/url"
	"strconv"

	"github.com/xabinapal/gopve/pkg/types"
)

type Client interface {
	Request(method string, resource string, form RequestValues, out interface{}) error

	WaitForTask(upid string) types.Task

	Lock()
	Unlock()
}

type RequestValues url.Values

func (v RequestValues) AddString(k string, s string) {
	v[k] = []string{s}
}

func (v RequestValues) AddInt(k string, i int) {
	v[k] = []string{strconv.Itoa(i)}
}

func (v RequestValues) AddUint(k string, u uint) {
	v[k] = []string{strconv.Itoa(int(u))}
}

func (v RequestValues) AddBool(k string, b bool) {
	v[k] = []string{PVEBool(b).String()}
}
