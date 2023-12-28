// Package accounts provides the /accounts APIs
package account

import (
	"net/http"

	cloudns "github.com/Eyup-Devop/cloudns"
	"github.com/Eyup-Devop/cloudns/auth"
)

// Client is used to invoke /accounts APIs.
type Client struct {
	B cloudns.Backend
}

// New creates a login response.
func NewLogin() (*cloudns.AccountLoginResponse, error) {
	return getC().NewLogin()
}

// Get returns login response.
func (c Client) NewLogin() (*cloudns.AccountLoginResponse, error) {
	loginResponse := &cloudns.AccountLoginResponse{}
	params := &auth.Auth{}
	err := c.B.Call(http.MethodPost, "/login/login.json", params, loginResponse, false)
	return loginResponse, err
}

func GetMyIp() (*cloudns.AccountIpResponse, error) {
	return getC().GetMyIp()
}

func (c Client) GetMyIp() (*cloudns.AccountIpResponse, error) {
	ipResponse := &cloudns.AccountIpResponse{}
	params := &auth.Auth{}
	err := c.B.Call(http.MethodPost, "/ip/get-my-ip.json", params, ipResponse, false)
	return ipResponse, err
}

func GetAccountBalance() (*cloudns.AccountBalanceResponse, error) {
	return getC().GetAccountBalance()
}

func (c Client) GetAccountBalance() (*cloudns.AccountBalanceResponse, error) {
	balanceResponse := &cloudns.AccountBalanceResponse{}
	params := &auth.Auth{}
	err := c.B.Call(http.MethodPost, "/account/get-balance.json", params, balanceResponse, false)
	return balanceResponse, err
}

func getC() Client {
	return Client{cloudns.GetBackend()}
}
