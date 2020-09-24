package request

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"regexp"
	"strings"
	"sync"
)

//go:generate mockery --case snake --name Executor

type Executor interface {
	StartAtomicBlock()
	EndAtomicBlock()

	Request(method, url string, form url.Values) ([]byte, error)

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

var errorRegExp = regexp.MustCompile(`^\d+\s*`)

func (exc *PVEExecutor) Request(
	method, path string,
	form url.Values,
) ([]byte, error) {
	absoluteURL, err := exc.getAbsoluteURL(method, path, form)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, absoluteURL.String(), nil)
	if err != nil {
		return nil, err
	}

	if method != http.MethodGet && form != nil {
		body := form.Encode()

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

		status := string(errorRegExp.ReplaceAll([]byte(res.Status), nil))
		return nil, fmt.Errorf("%d - %s", res.StatusCode, status)
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

func (exc *PVEExecutor) SetAuthenticationTicket(
	ticket string,
	method AuthenticationMethod,
) {
	exc.unsetAuthenticationTicket()

	switch method {
	case AuthenticationMethodCookie:
		var authCookie *http.Cookie

		if exc.client.Jar == nil {
			jar, err := cookiejar.New(nil)
			if err != nil {
				panic(fmt.Sprintf("This should never happen: %s", err.Error()))
			}

			exc.client.Jar = jar
		}

		authCookie = &http.Cookie{
			Name:  "PVEAuthCookie",
			Value: ticket,
		}

		exc.client.Jar.SetCookies(exc.base, []*http.Cookie{authCookie})

	case AuthenticationMethodHeader:
		exc.ticket = ticket
	}
}

func (exc *PVEExecutor) getAbsoluteURL(
	method, path string,
	form url.Values,
) (*url.URL, error) {
	resourceURL, err := url.Parse(strings.Trim(path, "/"))
	if err != nil {
		return nil, err
	}

	absoluteURL := exc.base.ResolveReference(resourceURL)

	if method == http.MethodGet && form != nil {
		absoluteURL.RawQuery = form.Encode()
	}

	return absoluteURL, nil
}

func (exc *PVEExecutor) unsetAuthenticationTicket() {
	if exc.client.Jar != nil {
		authCookie := &http.Cookie{Name: "PVEAuthCookie", MaxAge: -1}
		exc.client.Jar.SetCookies(exc.base, []*http.Cookie{authCookie})
	}

	exc.ticket = ""
}
