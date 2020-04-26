package foxyproxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Account struct {
	Active   bool   `json:"active"`
	Node     *Node  `json:"node,omitempty"`
	UID      string `json:"uid"`
	Username string `json:"username"`
}

// https://reseller.api.foxyproxy.com/#_username_exists
func (c *Client) UsernameExists(username string) (bool, error) {
	res, err := c.doRequest(fmt.Sprintf("/accounts/exists/%s/", username))
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
func (c *Client) DeactivateAccounts(username string) (int, error) {
	type countRes struct {
		Count int `json:"count"`
	}
	res, err := c.doRequest2(http.MethodPatch, fmt.Sprintf("/accounts/deactivate/%s/", username), nil)
	if err != nil {
		return 0, err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	resJSON := countRes{}
	if err := json.Unmarshal(bodyBytes, &resJSON); err != nil {
		return 0, err
	}
	return resJSON.Count, nil
}
