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

// https://reseller.api.foxyproxy.com/#_activate_accounts
func (c *Client) ActivateAccounts(username string) (int, error) {
	type countRes struct {
		Count int `json:"count"`
	}
	res, err := c.doRequest2(http.MethodPatch, fmt.Sprintf("/accounts/activate/%s/", username), nil)
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

// https://reseller.api.foxyproxy.com/#_update_passwords
func (c *Client) UpdatePassword(username, password string) (int, error) {
	// validate input
	if len(password) < 3 {
		return 0, fmt.Errorf("password must be more than 3 characters long")
	}
	if len(password) > 127 {
		return 0, fmt.Errorf("password must be less than 127 characters long")
	}

	type countRes struct {
		Count int `json:"count"`
	}
	res, err := c.doRequest2(
		http.MethodPatch,
		fmt.Sprintf("/accounts/update-password/%s", username),
		[]byte(fmt.Sprintf(`{ "password": "%s" }`, password)),
	)
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

// https://reseller.api.foxyproxy.com/#_common_api_properties
type commonAPIProperties struct {
	Comment   string   `json:"comment,omitempty"`
	NodeNames []string `json:"nodenames,omitempty"`
}

type DeleteAccountsParams struct {
	IncludeHistory bool `json:"includeHistory"`
	commonAPIProperties
}

// https://reseller.api.foxyproxy.com/#_delete_accounts
func (c *Client) DeleteAccounts(username string, params *DeleteAccountsParams) (int, error) {
	body := []byte{}
	if params != nil {
		var err error
		body, err = json.Marshal(params)
		if err != nil {
			return 0, err
		}
	}
	type countRes struct {
		Count int `json:"count"`
	}
	res, err := c.doRequest2(http.MethodPatch, fmt.Sprintf("/accounts/activate/%s/", username), body)
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
