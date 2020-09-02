package client_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func TestClientUserAuthentication(t *testing.T) {
	exc := new(mocks.Executor)
	cli := helpClientCreate(t, exc)

	values := request.Values{
		"username": {"testUsername"},
		"password": {"testPassword"},
	}

	response, err := ioutil.ReadFile("./testdata/access_ticket.json")
	if err != nil {
		t.Fatalf("Unexpected ioutil.ReadFile error: %s", err.Error())
	}

	exc.
		On("Request", http.MethodPost, "access/ticket", url.Values(values)).
		Return(response, nil).
		Once()

	exc.On("SetAuthenticationTicket", "authenticationToken", request.AuthenticationMethodCookie).Return(response).Once()
	exc.On("SetCSRFToken", "csrfToken").Return().Once()

	if err := cli.AuthenticateWithCredentials("testUsername", "testPassword"); err != nil {
		t.Fatalf("Unexpected client.Client.AuthenticateWithCredentials error: %s", err.Error())
	}

	exc.AssertExpectations(t)
}
