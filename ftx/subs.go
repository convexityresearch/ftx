package ftx

import (
	"encoding/json"
	"log"

	"github.com/convexity-research/ftx/ftx/api"
)

// SubAccountsList is an alias.
type SubAccountsList api.SubAccountsList

// SubAccount is an alias.
type SubAccount api.SubAccount

// Response is an alias.
type Response api.Response

// SubAccountBalances is an alias.
type SubAccountBalances api.SubAccountBalances

// TransferSubAccounts is an alias.
type TransferSubAccounts api.TransferSubAccounts

// GetSubAccounts will fetch the list of subaccounts.
func (client *Client) GetSubAccounts() (SubAccountsList, error) {
	var SubAccounts SubAccountsList
	resp, err := client._get("SubAccounts", []byte(""))
	if err != nil {
		log.Printf("Error GetSubAccounts", err)
		return SubAccounts, err
	}
	err = _processResponse(resp, &SubAccounts)
	return SubAccounts, err
}

// CreateSubAccount will create a new subaccount using a nickname.
func (client *Client) CreateSubAccount(nickname string) (SubAccount, error) {
	var SubAccount SubAccount
	requestBody, err := json.Marshal(map[string]string{"nickname": nickname})
	if err != nil {
		log.Printf("Error CreateSubAccount", err)
		return SubAccount, err
	}
	resp, err := client._post("SubAccounts", requestBody)
	if err != nil {
		log.Printf("Error CreateSubAccount", err)
		return SubAccount, err
	}
	err = _processResponse(resp, &SubAccount)
	return SubAccount, err
}

// ChangeSubAccountName will change the subaccount's nickname.
func (client *Client) ChangeSubAccountName(nickname string, newNickname string) (Response, error) {
	var changeSubAccount Response
	requestBody, err := json.Marshal(map[string]string{"nickname": nickname, "newNickname": newNickname})
	if err != nil {
		log.Printf("Error ChangeSubAccountName", err)
		return changeSubAccount, err
	}
	resp, err := client._post("SubAccounts/update_name", requestBody)
	if err != nil {
		log.Printf("Error ChangeSubAccountName", err)
		return changeSubAccount, err
	}
	err = _processResponse(resp, &changeSubAccount)
	return changeSubAccount, err
}

// DeleteSubAccount will delete a subaccount.
func (client *Client) DeleteSubAccount(nickname string) (Response, error) {
	var deleteSubAccount Response
	requestBody, err := json.Marshal(map[string]string{"nickname": nickname})
	if err != nil {
		log.Printf("Error DeleteSubAccount", err)
		return deleteSubAccount, err
	}
	resp, err := client._delete("SubAccounts", requestBody)
	if err != nil {
		log.Printf("Error DeleteSubAccount", err)
		return deleteSubAccount, err
	}
	err = _processResponse(resp, &deleteSubAccount)
	return deleteSubAccount, err
}

// GetSubAccountBalances will fetch the list of balances for each of your subaccounts.
func (client *Client) GetSubAccountBalances(nickname string) (SubAccountBalances, error) {
	var SubAccountBalances SubAccountBalances
	resp, err := client._get("SubAccounts/"+nickname+"/balances", []byte(""))
	if err != nil {
		log.Printf("Error SubAccountBalances", err)
		return SubAccountBalances, err
	}
	err = _processResponse(resp, &SubAccountBalances)
	return SubAccountBalances, err
}

// TransferSubAccounts will transfer balances accross your subaccounts.
func (client *Client) TransferSubAccounts(coin string, size float64, source string, destination string) (TransferSubAccounts, error) {
	var transferSubAccounts TransferSubAccounts
	requestBody, err := json.Marshal(map[string]interface{}{
		"coin":        coin,
		"size":        size,
		"source":      source,
		"destination": destination,
	})
	if err != nil {
		log.Printf("Error TransferSubAccounts", err)
		return transferSubAccounts, err
	}
	resp, err := client._post("SubAccounts/transfer", requestBody)
	if err != nil {
		log.Printf("Error TransferSubAccounts", err)
		return transferSubAccounts, err
	}
	err = _processResponse(resp, &transferSubAccounts)
	return transferSubAccounts, err
}
