package request

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync"
)

type Executor interface {
	StartAtomicBlock()
	EndAtomicBlock()

	Request(method string, url string, form Values) ([]byte, error)

	SetCSRFToken(token string)
	SetAuthenticationTicket(ticket string, method AuthenticationMethod)
}

type PVEExecutor struct {
	mux    *sync.Mutex
	client *http.Client
	base   *url.URL

	csrf   string
	ticket string
}

func NewPVEExecutor(base *url.URL, client *http.Client) *PVEExecutor {
	if client == nil {
		client = new(http.Client)
	}

	return &PVEExecutor{
		mux:    new(sync.Mutex),
		client: client,
		base:   base,
	}
}

func (exc *PVEExecutor) StartAtomicBlock() {
	exc.mux.Lock()
}

func (exc *PVEExecutor) EndAtomicBlock() {
	exc.mux.Unlock()
}

func (exc *PVEExecutor) Request(method string, path string, form Values) ([]byte, error) {
	absoluteURL, err := exc.getAbsoluteURL(path)
	if err != nil {
		return nil, err
	}

	if method == http.MethodGet && form != nil {
		absoluteURL.RawQuery = url.Values(form).Encode()
	}

	req, err := http.NewRequest(method, absoluteURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if method != http.MethodGet && form != nil {
		body := url.Values(form).Encode()

		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		req.ContentLength = int64(len(body))

		buf := bytes.NewBufferString(body)
		req.Body = ioutil.NopCloser(buf)
	}

	if exc.csrf != "" {
		req.Header.Add("CSRFPreventionToken", exc.csrf)
	}

	if exc.ticket != "" {
		req.Header.Add("Authorization", exc.ticket)
	}

	res, err := exc.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		res.Body.Close()
		return nil, fmt.Errorf(res.Status)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (exc *PVEExecutor) SetCSRFToken(token string) {
	exc.csrf = token
}

func (exc *PVEExecutor) SetAuthenticationTicket(ticket string, method AuthenticationMethod) {
	exc.unsetAuthenticationTicket()

	switch method {
	case AuthenticationMethodCookie:
		var authCookie *http.Cookie

		if exc.client.Jar == nil {
			jar, _ := cookiejar.New(nil)
			exc.client.Jar = jar
		}

		cookies := exc.client.Jar.Cookies(exc.base)
		for _, cookie := range cookies {
			if cookie.Name == "PVEAuthCookie" {
				authCookie = cookie
				break
			}
		}

		if authCookie == nil {
			authCookie = &http.Cookie{Name: "PVEAuthCookie"}
			cookies = append(cookies, authCookie)
		}

		authCookie.Value = ticket

		exc.client.Jar.SetCookies(exc.base, cookies)

	case AuthenticationMethodHeader:
		exc.ticket = ticket
	}
}

func (exc *PVEExecutor) getAbsoluteURL(path string) (*url.URL, error) {
	resourceURL, err := url.Parse(strings.TrimLeft(path, "/"))
	if err != nil {
		return nil, err
	}

	return exc.base.ResolveReference(resourceURL), nil
}

func (exc *PVEExecutor) unsetAuthenticationTicket() {
	if exc.client.Jar != nil {
		authCookie := &http.Cookie{Name: "PVEAuthCookie", MaxAge: -1}
		exc.client.Jar.SetCookies(exc.base, []*http.Cookie{authCookie})
	}

	exc.ticket = ""
}
