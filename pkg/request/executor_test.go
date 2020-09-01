package request_test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"strings"
	"testing"

	"github.com/xabinapal/gopve/pkg/request"
)

func testCreateHTTPServerHelper(t *testing.T, handler http.HandlerFunc) *httptest.Server {
	t.Helper()

	return httptest.NewServer(handler)
}

func testCreatePVEExecutorHelper(t *testing.T, srv *httptest.Server) *request.PVEExecutor {
	t.Helper()

	url, err := url.Parse(srv.URL)
	if err != nil {
		t.Fatalf("Unexpected url.Parse error: %s", err.Error())
	}

	url.Path = "/api2/json/"

	return request.NewPVEExecutor(url, srv.Client())
}

func testMakeExecutorRequestHelper(t *testing.T, exc *request.PVEExecutor, method, path string, form request.Values) {
	t.Helper()

	if _, err := exc.Request(method, path, form); err != nil {
		t.Fatalf("Unexpected request.PVEExecutor error: %s", err.Error())
	}
}

func TestExecutorRequestURLPath(t *testing.T) {
	srv := testCreateHTTPServerHelper(t, func(res http.ResponseWriter, req *http.Request) {
		url := req.URL.Path
		if url != "/api2/json/test" {
			t.Errorf("Got Resource '%s', expected '/api2/json/test'", url)
		}
	})
	defer srv.Close()

	exc := testCreatePVEExecutorHelper(t, srv)

	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/test", nil)
}

func TestExecutorRequestQueryString(t *testing.T) {
	values := request.Values{
		"test": {"test"},
	}

	srv := testCreateHTTPServerHelper(t, func(res http.ResponseWriter, req *http.Request) {
		form := req.URL.Query()

		data := url.Values(values)
		if !reflect.DeepEqual(form, data) {
			t.Errorf("Got QueryString '%s', expected '%s'", form.Encode(), data.Encode())
		}
	})
	defer srv.Close()

	exc := testCreatePVEExecutorHelper(t, srv)

	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/test", values)
}

func TestExecutorRequestFormData(t *testing.T) {
	values := request.Values{
		"test": {"test"},
	}

	srv := testCreateHTTPServerHelper(t, func(res http.ResponseWriter, req *http.Request) {
		if req.ContentLength != 0 {
			contentType := req.Header.Get("Content-Type")
			if contentType != "application/x-www-form-urlencoded" {
				t.Errorf("Got Content-Type '%s', expected 'application/x-www-form-urlencoded'", contentType)
			}
		}

		transferEncoding := strings.Join(req.TransferEncoding, ", ")
		if transferEncoding != "" {
			t.Errorf("Got Transfer-Encoding '%s', expected '<nil>'", transferEncoding)
		}

		form, err := ioutil.ReadAll(req.Body)
		if err != nil {
			t.Fatalf("Unexpected ioutil.ReadAll error: %s", err.Error())
		}

		formValues, err := url.ParseQuery(string(form))
		if err != nil {
			t.Fatalf("Unexpected url.ParseQuery error: %s", err.Error())
		}

		data := url.Values(values)
		if !reflect.DeepEqual(formValues, data) {
			t.Errorf("Got FormData '%s', expected '%s'", formValues.Encode(), data.Encode())
		}
	})
	defer srv.Close()

	exc := testCreatePVEExecutorHelper(t, srv)
	testMakeExecutorRequestHelper(t, exc, http.MethodPost, "/test", values)
}

func TestExecutorRequestCSRFPrevention(t *testing.T) {
	expectedTokenCount := 1

	srv := testCreateHTTPServerHelper(t, func(res http.ResponseWriter, req *http.Request) {
		csrfToken := req.Header.Get("CSRFPreventionToken")

		expectedToken := fmt.Sprintf("token%d", expectedTokenCount)
		expectedTokenCount++

		if csrfToken != expectedToken {
			t.Errorf("Got CSRF prevention token '%s', expected '%s'", csrfToken, expectedToken)
		}
	})
	defer srv.Close()

	exc := testCreatePVEExecutorHelper(t, srv)

	exc.SetCSRFToken("token1")
	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/", nil)

	exc.SetCSRFToken("token2")
	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/", nil)
}

func TestExecutorRequestCookieAuthentication(t *testing.T) {
	expectedTokenCount := 1

	srv := testCreateHTTPServerHelper(t, func(res http.ResponseWriter, req *http.Request) {
		cookieToken, err := req.Cookie("PVEAuthCookie")
		if err != nil {
			t.Errorf("No Cookie authentication token found: %s", err.Error())
		}

		expectedToken := fmt.Sprintf("token%d", expectedTokenCount)
		expectedTokenCount++

		if cookieToken.Value != expectedToken {
			t.Errorf("Got cookie authentication token '%s', expected '%s'", cookieToken, expectedToken)
		}
	})
	defer srv.Close()

	exc := testCreatePVEExecutorHelper(t, srv)

	exc.SetAuthenticationTicket("token1", request.AuthenticationMethodCookie)
	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/", nil)

	exc.SetAuthenticationTicket("token2", request.AuthenticationMethodCookie)
	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/", nil)
}

func TestExecutorRequestHeaderAuthentication(t *testing.T) {
	expectedTokenCount := 1

	srv := testCreateHTTPServerHelper(t, func(res http.ResponseWriter, req *http.Request) {
		headerToken := req.Header.Get("Authorization")

		expectedToken := fmt.Sprintf("token%d", expectedTokenCount)
		expectedTokenCount++

		if headerToken != expectedToken {
			t.Errorf("Got header authentication token '%s', expected '%s'", headerToken, expectedToken)
		}
	})
	defer srv.Close()

	exc := testCreatePVEExecutorHelper(t, srv)

	exc.SetAuthenticationTicket("token1", request.AuthenticationMethodHeader)
	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/", nil)

	exc.SetAuthenticationTicket("token2", request.AuthenticationMethodHeader)
	testMakeExecutorRequestHelper(t, exc, http.MethodGet, "/", nil)
}

func TestExecutorRequestMixedAuthentication(t *testing.T) {
	runnerHelper := func(t *testing.T, setAuthenticationFunction func(*request.PVEExecutor)) {
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

		exc := testCreatePVEExecutorHelper(t, srv)

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
