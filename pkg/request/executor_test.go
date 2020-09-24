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

func helpExecutorCreateServer(
	t *testing.T,
	handler http.HandlerFunc,
) *httptest.Server {
	t.Helper()

	srv := httptest.NewServer(handler)

	t.Cleanup(func() {
		srv.Close()
	})

	return srv
}

func helpExecutorCreateExecutor(
	t *testing.T,
	srv *httptest.Server,
) *request.PVEExecutor {
	t.Helper()

	url, err := url.Parse(srv.URL)
	require.NoError(t, err)

	url.Path = "/api2/json/"

	return request.NewPVEExecutor(url, srv.Client())
}

func helpExecutorMakeRequest(
	t *testing.T,
	exc *request.PVEExecutor,
	method, path string,
	form url.Values,
) {
	t.Helper()

	_, err := exc.Request(method, path, form)
	require.NoError(t, err)
}

func TestExecutorRequestURLPath(t *testing.T) {
	srv := helpExecutorCreateServer(
		t,
		func(res http.ResponseWriter, req *http.Request) {
			url := req.URL.Path
			assert.Equal(t, "/api2/json/test", url)
		},
	)

	exc := helpExecutorCreateExecutor(t, srv)

	helpExecutorMakeRequest(t, exc, http.MethodGet, "test", nil)
	helpExecutorMakeRequest(t, exc, http.MethodGet, "/test", nil)
	helpExecutorMakeRequest(t, exc, http.MethodGet, "test/", nil)
	helpExecutorMakeRequest(t, exc, http.MethodGet, "/test/", nil)
}

func TestExecutorRequestFormData(t *testing.T) {
	values := url.Values{
		"test": {"test"},
	}

	queryMethods := []string{http.MethodGet}
	for _, method := range queryMethods {
		t.Run(method, func(t *testing.T) {
			srv := helpExecutorCreateServer(
				t,
				func(res http.ResponseWriter, req *http.Request) {
					assert.Equal(t, method, req.Method)

					form := req.URL.Query()
					assert.Equal(t, values, form)
				},
			)

			exc := helpExecutorCreateExecutor(t, srv)

			helpExecutorMakeRequest(t, exc, method, "/test", values)
		})
	}

	bodyMethods := []string{http.MethodPost, http.MethodPut, http.MethodDelete}
	for _, method := range bodyMethods {
		t.Run(method, func(t *testing.T) {
			srv := helpExecutorCreateServer(
				t,
				func(res http.ResponseWriter, req *http.Request) {
					assert.Equal(t, method, req.Method)

					if req.ContentLength != 0 {
						contentType := req.Header.Get("Content-Type")
						assert.Equal(
							t,
							"application/x-www-form-urlencoded",
							contentType,
						)
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
				},
			)

			exc := helpExecutorCreateExecutor(t, srv)
			helpExecutorMakeRequest(t, exc, method, "/test", values)
		})
	}
}

func TestExecutorRequestCSRFPrevention(t *testing.T) {
	expectedTokenCount := 1

	srv := helpExecutorCreateServer(
		t,
		func(res http.ResponseWriter, req *http.Request) {
			expectedToken := fmt.Sprintf("token%d", expectedTokenCount)
			csrfToken := req.Header.Get("CSRFPreventionToken")
			assert.Equal(t, expectedToken, csrfToken)
			expectedTokenCount++
		},
	)

	exc := helpExecutorCreateExecutor(t, srv)

	exc.SetCSRFToken("token1")
	helpExecutorMakeRequest(t, exc, http.MethodGet, "/", nil)

	exc.SetCSRFToken("token2")
	helpExecutorMakeRequest(t, exc, http.MethodGet, "/", nil)
}

func TestExecutorRequestAuthenticationMethod(t *testing.T) {
	t.Run("Cookie", func(t *testing.T) {
		expectedTokenCount := 1

		srv := helpExecutorCreateServer(
			t,
			func(res http.ResponseWriter, req *http.Request) {
				expectedToken := fmt.Sprintf("token%d", expectedTokenCount)
				cookieToken, err := req.Cookie("PVEAuthCookie")
				require.NoError(t, err)
				assert.Equal(t, expectedToken, cookieToken.Value)
				expectedTokenCount++
			},
		)

		exc := helpExecutorCreateExecutor(t, srv)

		exc.SetAuthenticationTicket(
			"token1",
			request.AuthenticationMethodCookie,
		)
		helpExecutorMakeRequest(t, exc, http.MethodGet, "/", nil)

		exc.SetAuthenticationTicket(
			"token2",
			request.AuthenticationMethodCookie,
		)
		helpExecutorMakeRequest(t, exc, http.MethodGet, "/", nil)
	})

	t.Run("Header", func(t *testing.T) {
		expectedTokenCount := 1

		srv := helpExecutorCreateServer(
			t,
			func(res http.ResponseWriter, req *http.Request) {
				expectedToken := fmt.Sprintf("token%d", expectedTokenCount)
				headerToken := req.Header.Get("Authorization")
				assert.Equal(t, expectedToken, headerToken)
				expectedTokenCount++
			},
		)

		exc := helpExecutorCreateExecutor(t, srv)

		exc.SetAuthenticationTicket(
			"token1",
			request.AuthenticationMethodHeader,
		)
		helpExecutorMakeRequest(t, exc, http.MethodGet, "/", nil)

		exc.SetAuthenticationTicket(
			"token2",
			request.AuthenticationMethodHeader,
		)
		helpExecutorMakeRequest(t, exc, http.MethodGet, "/", nil)
	})
}

func TestExecutorRequestMixedAuthenticationMethods(t *testing.T) {
	runnerHelper := func(t *testing.T, setAuthenticationFunction func(*request.PVEExecutor)) {
		t.Helper()

		srv := helpExecutorCreateServer(
			t,
			func(res http.ResponseWriter, req *http.Request) {
				var cookieToken string
				cookie, err := req.Cookie("PVEAuthCookie")
				if err == nil {
					cookieToken = cookie.Value
				}

				headerToken := req.Header.Get("Authorization")

				assert.Condition(t, func() bool {
					return (cookieToken == "") != (headerToken == "")
				})
			},
		)

		exc := helpExecutorCreateExecutor(t, srv)

		setAuthenticationFunction(exc)
		helpExecutorMakeRequest(t, exc, http.MethodGet, "/", nil)
	}

	t.Run("CookieFirst", func(t *testing.T) {
		runnerHelper(t, func(exc *request.PVEExecutor) {
			exc.SetAuthenticationTicket(
				"token",
				request.AuthenticationMethodCookie,
			)
			exc.SetAuthenticationTicket(
				"token",
				request.AuthenticationMethodHeader,
			)
		})
	})

	t.Run("HeaderFirst", func(t *testing.T) {
		runnerHelper(t, func(exc *request.PVEExecutor) {
			exc.SetAuthenticationTicket(
				"token",
				request.AuthenticationMethodHeader,
			)
			exc.SetAuthenticationTicket(
				"token",
				request.AuthenticationMethodCookie,
			)
		})
	})
}

func TestExecutorRequestError(t *testing.T) {
	srv := helpExecutorCreateServer(
		t,
		func(res http.ResponseWriter, req *http.Request) {
			res.WriteHeader(http.StatusInternalServerError)
			res.Write([]byte("Internal Server Error"))
		},
	)

	exc := helpExecutorCreateExecutor(t, srv)

	_, err := exc.Request(http.MethodGet, "/", nil)
	assert.Error(t, err)
}
