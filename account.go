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
