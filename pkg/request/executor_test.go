package request_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/xabinapal/gopve/pkg/request"
)

func helpExecutorCreateHTTPServer(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()

	srv := httptest.NewServer(handler)

	t.Cleanup(func() {
		srv.Close()
	})

	return srv
}

func helpExecutorCreatePVEExecutor(t *testing.T, srv *httptest.Server) *request.PVEExecutor {
	t.Helper()

	url, err := url.Parse(srv.URL)
	require.NoError(t, err)

	url.Path = "/api2/json/"

	return request.NewPVEExecutor(url, srv.Client())
}

func testMakeExecutorRequestHelper(t *testing.T, exc *request.PVEExecutor, method, path string, form url.Values) {
	t.Helper()

	_, err := exc.Request(method, path, form)
	require.NoError(t, err)
}

func TestExecutorRequestURLPath(t *testing.T) {
	srv := helpExecutorCreateHTTPServer(t, func(res http.ResponseWriter, req *http.Request) {
		url := req.URL.Path
		assert.Equal(t, "/api2/json/test", url)
	})

	exc := helpExecutorCreatePVEExecutor(t, srv)

	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/test", nil)
}

func TestExecutorRequestQueryString(t *testing.T) {
	values := url.Values{
		"test": {"test"},
	}

	srv := helpExecutorCreateHTTPServer(t, func(res http.ResponseWriter, req *http.Request) {
		form := req.URL.Query()
		assert.Equal(t, values, form)
	})

	exc := helpExecutorCreatePVEExecutor(t, srv)

	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/test", values)
}

func TestExecutorRequestFormData(t *testing.T) {
	values := url.Values{
		"test": {"test"},
	}

	srv := helpExecutorCreateHTTPServer(t, func(res http.ResponseWriter, req *http.Request) {
		if req.ContentLength != 0 {
			contentType := req.Header.Get("Content-Type")
			assert.Equal(t, "application/x-www-form-urlencoded", contentType)
		}

		transferEncoding := strings.Join(req.TransferEncoding, ", ")
		if transferEncoding != "" {
			assert.Equal(t, req.TransferEncoding, []string{})
		}

		body, err := ioutil.ReadAll(req.Body)
		require.NoError(t, err)

		form, err := url.ParseQuery(string(body))
		require.NoError(t, err)

		require.Equal(t, values, form)
	})

	exc := helpExecutorCreatePVEExecutor(t, srv)
	testMakeExecutorRequestHelper(t, exc, http.MethodPost, "/test", values)
}

func TestExecutorRequestCSRFPrevention(t *testing.T) {
	expectedTokenCount := 1

	srv := helpExecutorCreateHTTPServer(t, func(res http.ResponseWriter, req *http.Request) {
		expectedToken := fmt.Sprintf("token%d", expectedTokenCount)
		csrfToken := req.Header.Get("CSRFPreventionToken")
		assert.Equal(t, expectedToken, csrfToken)
		expectedTokenCount++
	})

	exc := helpExecutorCreatePVEExecutor(t, srv)

	exc.SetCSRFToken("token1")
	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/", nil)

	exc.SetCSRFToken("token2")
	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/", nil)
}

func TestExecutorRequestCookieAuthentication(t *testing.T) {
	expectedTokenCount := 1

	srv := helpExecutorCreateHTTPServer(t, func(res http.ResponseWriter, req *http.Request) {
		expectedToken := fmt.Sprintf("token%d", expectedTokenCount)
		cookieToken, err := req.Cookie("PVEAuthCookie")
		require.NoError(t, err)
		assert.Equal(t, expectedToken, cookieToken.Value)
		expectedTokenCount++
	})

	exc := helpExecutorCreatePVEExecutor(t, srv)

	exc.SetAuthenticationTicket("token1", request.AuthenticationMethodCookie)
	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/", nil)

	exc.SetAuthenticationTicket("token2", request.AuthenticationMethodCookie)
	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/", nil)
}

func TestExecutorRequestHeaderAuthentication(t *testing.T) {
	expectedTokenCount := 1

	srv := helpExecutorCreateHTTPServer(t, func(res http.ResponseWriter, req *http.Request) {
		expectedToken := fmt.Sprintf("token%d", expectedTokenCount)
		headerToken := req.Header.Get("Authorization")
		assert.Equal(t, expectedToken, headerToken)
		expectedTokenCount++
	})

	exc := helpExecutorCreatePVEExecutor(t, srv)

	exc.SetAuthenticationTicket("token1", request.AuthenticationMethodHeader)
	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/", nil)

	exc.SetAuthenticationTicket("token2", request.AuthenticationMethodHeader)
	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/", nil)
}

func TestExecutorRequestMixedAuthentication(t *testing.T) {
	runnerHelper := func(t *testing.T, setAuthenticationFunction func(*request.PVEExecutor)) {
		t.Helper()

		srv := helpExecutorCreateHTTPServer(t, func(res http.ResponseWriter, req *http.Request) {
			var cookieToken string
			cookie, err := req.Cookie("PVEAuthCookie")
			if err == nil {
				cookieToken = cookie.Value
			}

			headerToken := req.Header.Get("Authorization")

			assert.Condition(t, func() bool {
				return (cookieToken == "") != (headerToken == "")
			})
		})

		exc := helpExecutorCreatePVEExecutor(t, srv)

		setAuthenticationFunction(exc)
		testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/", nil)
	}

	t.Run("CookieFirst", func(t *testing.T) {
		runnerHelper(t, func(exc *request.PVEExecutor) {
			exc.SetAuthenticationTicket("token", request.AuthenticationMethodCookie)
			exc.SetAuthenticationTicket("token", request.AuthenticationMethodHeader)
		})
	})

	t.Run("HeaderFirst", func(t *testing.T) {
		runnerHelper(t, func(exc *request.PVEExecutor) {
			exc.SetAuthenticationTicket("token", request.AuthenticationMethodHeader)
			exc.SetAuthenticationTicket("token", request.AuthenticationMethodCookie)
		})
	})
}
