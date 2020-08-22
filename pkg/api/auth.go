package api

import (
	"fmt"
	"net/http"

	"github.com/xabinapal/gopve/internal/pkg/utils"
)

type ticketResponseJSON struct {
	Username  string `json:"username"`
	Ticket    string `json:"ticket"`
	CSRFToken string `json:"CSRFPreventionToken"`

	ClusterName string `json:"clustername"`
}

func (api *API) AuthenticateWithCredentials(username string, password string) error {
	var res ticketResponseJSON
	err := api.client.Request(http.MethodPost, "access/ticket", utils.RequestValues{
		"username": {username},
		"password": {password},
	}, &res)
	if err != nil {
		return err
	}

	api.client.setAuthenticationTicket(res.Ticket)
	api.client.csrf = res.CSRFToken

	return nil
}

func (api *API) AuthenticateWithToken(id string, secret string) error {
	api.client.ticket = fmt.Sprintf("PVEAPIToken=%s!TOKENID=%s", id, secret)
	api.client.unsetAuthenticationTicket()

	// TODO: add request to validate token

	return nil
}

func (c *client) setAuthenticationTicket(ticket string) {
	var authCookie *http.Cookie

	cookies := c.client.Jar.Cookies(c.url)
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

	c.client.Jar.SetCookies(c.url, cookies)
}

func (c *client) unsetAuthenticationTicket() {
	cookies := c.client.Jar.Cookies(c.url)
	for i, cookie := range cookies {
		if cookie.Name == "PVEAuthCookie" {
			cookies = append(cookies[:i], cookies[i+1:]...)
			c.client.Jar.SetCookies(c.url, cookies)
			break
		}
	}
}
