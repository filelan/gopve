package internal

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"time"
)

type Client struct {
	client    *http.Client
	apiURI    string
	csrfToken string
}

func NewClient(uri string, user string, password string, invalidCert bool) (*Client, error) {
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives:   false,
			IdleConnTimeout:     0,
			MaxIdleConns:        200,
			MaxIdleConnsPerHost: 100,
		},
		Timeout: time.Second * 10,
	}

	if invalidCert {
		client.Transport.(*http.Transport).TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}

	c := &Client{
		client: client,
		apiURI: uri,
	}

	err := c.authenticate(user, password)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (c *Client) authenticate(user string, password string) error {
	form := url.Values{
		"username": {user},
		"password": {password},
	}

	data, err := c.Post("access/ticket", form)
	if err != nil {
		return err
	}

	authCookie := &http.Cookie{
		Name:  "PVEAuthCookie",
		Value: data["ticket"].(string),
	}

	c.client.Jar, err = cookiejar.New(nil)
	if err != nil {
		return err
	}

	cookieURI, err := url.Parse(c.apiURI)
	if err != nil {
		return err
	}
	c.client.Jar.SetCookies(cookieURI, []*http.Cookie{authCookie})
	c.csrfToken = data["CSRFPreventionToken"].(string)

	return nil
}

func (c *Client) request(method string, endpoint string, data url.Values) (*http.Response, error) {
	url := c.apiURI + endpoint

	var buf *bytes.Buffer
	if data != nil {
		buf = bytes.NewBufferString(data.Encode())
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	if data != nil {
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	}

	if c.csrfToken != "" {
		req.Header.Add("CSRFPreventionToken", c.csrfToken)
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *Client) Get(endpoint string) (map[string]interface{}, error) {
	res, err := c.request("GET", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		return nil, err
	}

	return out, nil
}

func (c *Client) Post(endpoint string, data url.Values) (map[string]interface{}, error) {
	res, err := c.request("POST", endpoint, nil)
	if err != nil {
		return nil, err
	}

	var out map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&out)
	if err != nil {
		return nil, err
	}

	return out, nil
}
