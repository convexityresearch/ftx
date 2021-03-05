package api

import (
	"time"
)

// SubAccountsList represents the reponse of the sub accounts list api.
type SubAccountsList struct {
	Success bool `json:"success"`
	Result  []struct {
		Nickname  string `json:"nickname"`
		Deletable bool   `json:"deletable"`
		Editable  bool   `json:"editable"`
	} `json:"result"`
}

// SubAccount is how subs are represented in ftx.
type SubAccount struct {
	Success bool `json:"success"`
	Result  struct {
		Nickname  string `json:"nickname"`
		Deletable bool   `json:"deletable"`
		Editable  bool   `json:"editable"`
	} `json:"result"`
}

// SubAccountBalances represent the active balance of a subaccount.
type SubAccountBalances struct {
	Success bool `json:"success"`
	Result  []struct {
		Coin  string  `json:"coin"`
		Free  float64 `json:"free"`
		Total float64 `json:"total"`
	} `json:"result"`
}

// TransferSubAccounts represents a transfer request.
type TransferSubAccounts struct {
	Success bool `json:"success"`
	Result  struct {
		ID     int       `json:"id"`
		Coin   string    `json:"coin"`
		Size   float64   `json:"size"`
		Time   time.Time `json:"time"`
		Notes  string    `json:"notes"`
		Status string    `json:"status"`
	} `json:"result"`
}
