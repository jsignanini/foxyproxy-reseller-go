package foxyproxy

import (
	"fmt"
	"net/http"
)

type Account struct {
	Active   bool   `json:"active"`
	Node     *Node  `json:"node,omitempty"`
	UID      string `json:"uid"`
	Username string `json:"username"`

	client *Client
}

func NewAccount(c *Client) *Account {
	return &Account{
		client: c,
	}
}

// https://reseller.api.foxyproxy.com/#_username_exists
func (c *Client) UsernameExists(username string) (bool, error) {
	res, err := c.doRequest(http.MethodGet, fmt.Sprintf("/accounts/exists/%s/", username), nil)
	if err != nil {
		return false, err
	}
	switch res.StatusCode {
	case http.StatusOK:
		return true, nil
	case http.StatusNotFound:
		return false, nil
	default:
		return false, fmt.Errorf("unexpected response")
	}
}

// https://reseller.api.foxyproxy.com/#_deactivate_accounts
func (a *Account) Deactivate() (int, error) {
	return a.client.deactivateAccount(a.Username, &commonAPIProperties{
		NodeNames: []string{a.Node.Name},
	})
}

// https://reseller.api.foxyproxy.com/#_activate_accounts
func (a *Account) Activate() (int, error) {
	return a.client.activateAccount(a.Username, &commonAPIProperties{
		NodeNames: []string{a.Node.Name},
	})
}

// https://reseller.api.foxyproxy.com/#_update_passwords
func (a *Account) UpdatePassword(password string) (int, error) {
	return a.client.updatePassword(a.Username, password, &commonAPIProperties{
		NodeNames: []string{a.Node.Name},
	})
}

// https://reseller.api.foxyproxy.com/#_common_api_properties
type commonAPIProperties struct {
	Comment   string   `json:"comment,omitempty"`
	NodeNames []string `json:"nodeNames,omitempty"`
}

type DeleteAccountsParams struct {
	IncludeHistory bool `json:"includeHistory"`
	commonAPIProperties
}

// https://reseller.api.foxyproxy.com/#_delete_accounts
func (a *Account) Delete(includeHistory bool) (int, error) {
	return a.client.deleteAccounts(a.Username, includeHistory, &commonAPIProperties{
		NodeNames: []string{a.Node.Name},
	})
}
