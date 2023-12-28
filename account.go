package cloudns

type AccountLoginResponse struct {
	Status            string `json:"status"`
	StatusDescription string `json:"statusDescription"`
}

type AccountIpResponse struct {
	Ip string `json:"ip"`
}

type AccountBalanceResponse struct {
	Funds string `json:"funds"`
}
