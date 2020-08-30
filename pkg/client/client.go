package client

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/xabinapal/gopve/pkg/request"
)

const DefaultRequestTimeout = time.Duration(30) * time.Second
const DefaultPoolingInterval = time.Duration(5) * time.Second

type Client struct {
	cfg      Config
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

	executor := cfg.Executor
	if executor == nil {
		timeout := cfg.RequestTimeout
		if timeout == 0 {
			timeout = DefaultRequestTimeout
		}

		executor = request.NewPVEExecutor(url, &http.Client{
			Transport: transport,
			Timeout:   timeout,
		})
	}

	poolingInterval := cfg.PoolingInterval
	if poolingInterval == 0 {
		poolingInterval = DefaultPoolingInterval
	}

	cli := &Client{
		cfg:      cfg,
		executor: executor,
	}

	cli.api = NewAPI(cli, poolingInterval)

	return cli, nil
}

func (cli *Client) StartAtomicBlock() {
	cli.executor.StartAtomicBlock()
}

func (cli *Client) EndAtomicBlock() {
	cli.executor.EndAtomicBlock()
}

func (cli *Client) Request(method string, resource string, form request.Values, out interface{}) error {
	data, err := cli.executor.Request(method, resource, form)
	if err != nil {
		return err
	}

	if out != nil {
		var raw struct {
			Data json.RawMessage
		}

		if err := json.Unmarshal(data, &raw); err != nil {
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
