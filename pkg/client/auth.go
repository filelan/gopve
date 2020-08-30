package client

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/pkg/request"
)

type ticketResponseJSON struct {
	Username  string `json:"username"`
	Ticket    string `json:"ticket"`
	CSRFToken string `json:"CSRFPreventionToken"`

	ClusterName string `json:"clustername"`
}

func (cli *Client) AuthenticateWithCredentials(username string, password string) error {
	var res ticketResponseJSON
	err := cli.Request(http.MethodPost, "access/ticket", request.Values{
		"username": {username},
		"password": {password},
	}, &res)
	if err != nil {
		return err
	}

	cli.executor.SetAuthenticationTicket(res.Ticket, request.AuthenticationMethodCookie)
	cli.executor.SetCSRFToken(res.CSRFToken)

	return nil
}

func (cli *Client) AuthenticateWithToken(id string, secret string) error {
	ticket := fmt.Sprintf("PVEAPIToken=%s!TOKENID=%s", id, secret)
	cli.executor.SetAuthenticationTicket(ticket, request.AuthenticationMethodHeader)

	// TODO: add request to validate token

	return nil
}
