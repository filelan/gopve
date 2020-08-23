package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"sync"

	"github.com/xabinapal/gopve/internal/pkg/debug"
	"github.com/xabinapal/gopve/internal/pkg/utils"
)

type client struct {
	client *http.Client
	url    *url.URL
	csrf   string
	ticket string

	poolingInterval int

	mutex sync.Mutex
}

func (c *client) Request(method string, resource string, form utils.RequestValues, out interface{}) error {
	resourceURL, err := url.Parse(strings.TrimLeft(resource, "/"))
	if err != nil {
		return err
	}

	absoluteURL := c.url.ResolveReference(resourceURL)

	var enc string
	if form != nil && len(form) > 0 {
		if method == http.MethodGet {
			absoluteURL.RawQuery = url.Values(form).Encode()
		} else {
			enc = url.Values(form).Encode()
		}
	}

	debug.Printf("request=%s", absoluteURL.String())

	buf := bytes.NewBufferString(enc)
	req, err := http.NewRequest(method, absoluteURL.String(), buf)
	if err != nil {
		return err
	}

	if enc != "" {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(enc)))
	}

	if c.csrf != "" {
		req.Header.Add("CSRFPreventionToken", c.csrf)
	}

	if c.ticket != "" {
		req.Header.Add("Authorization", c.ticket)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf(res.Status)
	}

	if out != nil {
		decoder := json.NewDecoder(res.Body)

		var raw struct {
			Data json.RawMessage
		}
		err = decoder.Decode(&raw)
		if err != nil {
			return err
		}

		debug.Printf("%s", raw.Data)

		err = json.Unmarshal(raw.Data, &out)
		if err != nil {
			return err
		}
	}

	return nil
}

func (c *client) Lock() {
	c.mutex.Lock()
}

func (c *client) Unlock() {
	c.mutex.Unlock()
}
