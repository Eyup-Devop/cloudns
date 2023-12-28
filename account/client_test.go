package account

import (
	"testing"

	cloudns "github.com/Eyup-Devop/cloudns"
	assert "github.com/stretchr/testify/require"
)

var (
	authId       *string = cloudns.String("****")
	authPassword *string = cloudns.String("*********")
)

func TestAccountLogin(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	loginResponse, err := NewLogin()
	assert.Nil(t, err)
	assert.NotNil(t, loginResponse)
}

func TestGetMyIp(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	ipResponse, err := GetMyIp()
	assert.Nil(t, err)
	assert.NotNil(t, ipResponse)
}

func TestGetAccountBalance(t *testing.T) {
	cloudns.AuthId = authId
	cloudns.AuthPassword = authPassword
	accountBalance, err := GetAccountBalance()
	assert.Nil(t, err)
	assert.NotNil(t, accountBalance)
}
