package client_test

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xabinapal/gopve/pkg/client"
	"github.com/xabinapal/gopve/pkg/request"
	"github.com/xabinapal/gopve/pkg/request/mocks"
)

func TestClientUserAuthentication(t *testing.T) {
	exc := new(mocks.Executor)
	cli := client.NewClientWithExecutor(exc, 0)

	values := request.Values{
		"username": {"testUsername"},
		"password": {"testPassword"},
	}

	response, err := ioutil.ReadFile("./testdata/access_ticket.json")
	assert.NoError(t, err)

	exc.
		On("Request", http.MethodPost, "access/ticket", url.Values(values)).
		Return(response, nil).
		Once()

	exc.On("SetAuthenticationTicket", "authenticationToken", request.AuthenticationMethodCookie).Return(response).Once()
	exc.On("SetCSRFToken", "csrfToken").Return().Once()

	err = cli.AuthenticateWithCredentials("testUsername", "testPassword")
	assert.NoError(t, err)

	exc.AssertExpectations(t)
}
