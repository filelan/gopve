package client

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/xabinapal/gopve/pkg/request"
)

const (
	DefaultRequestTimeout  = time.Duration(30) * time.Second
	DefaultPoolingInterval = time.Duration(5) * time.Second
)

type Client struct {
	executor request.Executor

	api API
}

func NewClient(cfg Config) (*Client, error) {
	url, err := cfg.Endpoint()
	if err != nil {
		return nil, err
	}

	transport := cfg.HTTPTransport
	if transport == nil {
		transport = new(http.Transport)
	}

	timeout := cfg.RequestTimeout
	if timeout == 0 {
		timeout = DefaultRequestTimeout
	}

	exc := request.NewPVEExecutor(url, &http.Client{
		Transport: transport,
		Timeout:   timeout,
	})

	return NewClientWithExecutor(exc, cfg.PoolingInterval), nil
}

func NewClientWithExecutor(
	exc request.Executor,
	poolingInterval time.Duration,
) *Client {
	if poolingInterval < time.Duration(1)*time.Second {
		poolingInterval = DefaultPoolingInterval
	}

	cli := &Client{
		executor: exc,
	}

	cli.api = NewAPI(cli, poolingInterval)

	return cli
}

func (cli *Client) StartAtomicBlock() {
	cli.executor.StartAtomicBlock()
}

func (cli *Client) EndAtomicBlock() {
	cli.executor.EndAtomicBlock()
}

func (cli *Client) Request(
	method, resource string,
	form request.Values,
	out interface{},
) error {
	data, err := cli.executor.Request(method, resource, url.Values(form))
	if err != nil {
		return err
	}

	if out != nil {
		var raw struct {
			Data json.RawMessage
		}

		if err = json.Unmarshal(data, &raw); err != nil {
			return err
		}

		if err = json.Unmarshal(raw.Data, &out); err != nil {
			return err
		}
	}

	return nil
}

func (cli *Client) API() API {
	return cli.api
}
