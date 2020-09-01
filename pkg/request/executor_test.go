package request_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/xabinapal/gopve/pkg/request"
)

func testCreateHTTPServerHelper(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()

	return httptest.NewServer(handler)
}

func testCreatePVEExecutor(t *testing.T, srv *httptest.Server) *request.PVEExecutor {
	url, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatalf("Unexpected url.Parse error: %s", err.Error())
	}

	return request.NewPVEExecutor(url, srv.Client())
}

func TestExecutorRequestHeaders(t *testing.T) {
	srv := testCreateHTTPServerHelper(t, func(res http.ResponseWriter, req *http.Request) {
		transferEncoding := strings.Join(req.TransferEncoding, ", ")
		if transferEncoding != "" {
			t.Errorf("Got Transfer-Encoding '%s', expected '<nil>'", transferEncoding)
		}

		if req.ContentLength != 0 {
			contentType := req.Header.Get("Content-Type")
			if contentType != "application/x-www-form-urlencoded" {
				t.Errorf("Got Content-Type '%s', expected 'application/x-www-form-urlencoded'", contentType)
			}
		}
	})
	defer srv.Close()

	exc := testCreatePVEExecutor(t, srv)

	if _, err := exc.Request(http.MethodGet, "/", nil); err != nil {
		t.Fatalf("Unexpected request.PVEExecutor error: %s", err.Error())
	}

	if _, err := exc.Request(http.MethodPost, "/", nil); err != nil {
		t.Fatalf("Unexpected request.PVEExecutor error: %s", err.Error())
	}

	if _, err := exc.Request(http.MethodPost, "/", request.Values{
		"test": {"test"},
	}); err != nil {
		t.Fatalf("Unexpected request.PVEExecutor error: %s", err.Error())
	}
}

func TestExecutorRequestCSRFPrevention(t *testing.T) {
	srv := testCreateHTTPServerHelper(t, func(res http.ResponseWriter, req *http.Request) {
		csrfToken := req.Header.Get("CSRFPreventionToken")
		if csrfToken != "token" {
			t.Errorf("Got CSRF prevention token '%s', expected 'token'", csrfToken)
		}
	})
	defer srv.Close()

	exc := testCreatePVEExecutor(t, srv)
	exc.SetCSRFToken("token")

	if _, err := exc.Request(http.MethodGet, "/", nil); err != nil {
		t.Fatalf("Unexpected request.PVEExecutor error: %s", err.Error())
	}
}

func TestExecutorRequestCookieAuthentication(t *testing.T) {
	srv := testCreateHTTPServerHelper(t, func(res http.ResponseWriter, req *http.Request) {
		cookieToken, err := req.Cookie("PVEAuthCookie")
		if err != nil {
			t.Errorf("No Cookie authentication token found: %s", err.Error())
		}
		if cookieToken.Value != "token" {
			t.Errorf("Got Cookie authentication token '%s', expected 'token'", cookieToken)
		}
	})
	defer srv.Close()

	exc := testCreatePVEExecutor(t, srv)
	exc.SetAuthenticationTicket("token", request.AuthenticationMethodCookie)

	if _, err := exc.Request(http.MethodGet, "/", nil); err != nil {
		t.Fatalf("Unexpected request.PVEExecutor error: %s", err.Error())
	}
}

func TestExecutorRequestHeaderAuthentication(t *testing.T) {
	srv := testCreateHTTPServerHelper(t, func(res http.ResponseWriter, req *http.Request) {
		headerToken := req.Header.Get("Authorization")
		if headerToken != "token" {
			t.Errorf("Got Cookie authentication token '%s', expected token", headerToken)
		}
	})
	defer srv.Close()

	exc := testCreatePVEExecutor(t, srv)
	exc.SetAuthenticationTicket("token", request.AuthenticationMethodHeader)

	if _, err := exc.Request(http.MethodGet, "/", nil); err != nil {
		t.Fatalf("Unexpected request.PVEExecutor error: %s", err.Error())
	}
}

func TestExecutorRequestMixedAuthentication(t *testing.T) {
	runnerHelper := func(t *testing.T, fn func(*request.PVEExecutor)) {
		t.Helper()

		srv := testCreateHTTPServerHelper(t, func(res http.ResponseWriter, req *http.Request) {
			var cookieToken string
			cookie, err := req.Cookie("PVEAuthCookie")
			if err == nil {
				cookieToken = cookie.Value
			}

			headerToken := req.Header.Get("Authorization")

			if cookieToken != "" && headerToken != "" {
				t.Errorf("Got both Cookie and Header authentication tokens, expected only one of them")
			}
		})
		defer srv.Close()

		exc := testCreatePVEExecutor(t, srv)
		fn(exc)

		if _, err := exc.Request(http.MethodGet, "/", nil); err != nil {
			t.Fatalf("Unexpected request.PVEExecutor error: %s", err.Error())
		}
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
