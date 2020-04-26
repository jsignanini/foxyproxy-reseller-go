package foxyproxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	username, password string
	domainHeader       string
	endpointBaseURL    string
}

type NewClientParams struct {
	Username, Password string
	DomainHeader       string
	EndpointBaseURL    string
}

func NewClient(params *NewClientParams) *Client {
	// TODO handle missing parameters
	return &Client{
		username:        params.Username,
		password:        params.Password,
		domainHeader:    params.DomainHeader,
		endpointBaseURL: params.EndpointBaseURL,
	}
}

func (c *Client) GetActiveNodeConnectionsByAccount(nodeName string) ([]*NodeConnection, error) {
	return c.getActiveNodeConnectionsByAccount(nodeName)
}

func (c *Client) GetActiveNodeConnectionTotals(nodeName string) (int, error) {
	return c.getActiveNodeConnectionTotals(nodeName)
}

func (c *Client) GetAllNodes(index, size int) ([]*Node, error) {
	// validate input
	if index < 0 {
		return nil, fmt.Errorf("index cannot be less than 0")
	}
	if size > 100 {
		return nil, fmt.Errorf("size cannot be larger than 100")
	}

	// get nodes
	res, err := c.doRequest(http.MethodGet, fmt.Sprintf("/nodes/?index=%d&size=%d", index, size), nil)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	nodes := []*Node{}
	if err := json.Unmarshal(bodyBytes, &nodes); err != nil {
		return nil, err
	}

	// populate client
	for _, n := range nodes {
		n.client = c
	}
	return nodes, nil
}

func (c *Client) GetHistoricalNodeConnectionsByAccount(nodeName string, startTime, endTime time.Time) ([]*NodeConnection, error) {
	return c.getHistoricalNodeConnectionsByAccount(nodeName, startTime, endTime)
}

func (c *Client) GetHistoricalNodeConnectionTotals(nodeName string, startTime, endTime time.Time) (int, error) {
	return c.getHistoricalNodeConnectionTotals(nodeName, startTime, endTime)
}

func (c *Client) GetNode(nodeName string) (*Node, error) {
	res, err := c.doRequest(http.MethodGet, fmt.Sprintf("/nodes/%s", nodeName), nil)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	node := NewNode(c)
	if err := json.Unmarshal(bodyBytes, node); err != nil {
		return nil, err
	}
	return node, nil
}

func (c *Client) GetNodeCount() (int, error) {
	type total struct {
		Count int
	}
	res, err := c.doRequest(http.MethodGet, "/nodes/count/", nil)
	if err != nil {
		return 0, err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	t := &total{}
	if err := json.Unmarshal(bodyBytes, t); err != nil {
		return 0, err
	}
	return t.Count, nil
}

func (c *Client) GetNodeTrafficByAccount(nodeName string, startTime, endTime time.Time) ([]*NodeTrafficAccount, error) {
	return c.getNodeTrafficByAccount(nodeName, startTime, endTime)
}

func (c *Client) GetNodeTrafficTotals(nodeName string, startTime, endTime time.Time) (*NodeTrafficTotals, error) {
	return c.getNodeTrafficTotals(nodeName, startTime, endTime)
}

func (c *Client) getActiveNodeConnectionTotals(nodeName string) (int, error) {
	type total struct {
		Count int
	}
	res, err := c.doRequest(http.MethodGet, fmt.Sprintf("/nodes/%s/connections/", nodeName), nil)
	if err != nil {
		return 0, err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	t := &total{}
	if err := json.Unmarshal(bodyBytes, t); err != nil {
		return 0, err
	}
	return t.Count, nil
}

func (c *Client) getDNSSuffixes() ([]string, error) {
	res, err := c.doRequest(http.MethodGet, "/nodes/dns-suffixes/", nil)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	suffixes := []string{}
	if err := json.Unmarshal(bodyBytes, &suffixes); err != nil {
		return nil, err
	}
	return suffixes, nil
}

func (c *Client) getHistoricalNodeConnectionTotals(nodeName string, startTime, endTime time.Time) (int, error) {
	type total struct {
		Count int
	}
	res, err := c.doRequest(http.MethodGet, fmt.Sprintf("/nodes/%s/connections/%d/%d/", nodeName, startTime.Unix(), endTime.Unix()), nil)
	if err != nil {
		return 0, err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}
	t := &total{}
	if err := json.Unmarshal(bodyBytes, t); err != nil {
		return 0, err
	}
	return t.Count, nil
}

// https://reseller.api.foxyproxy.com/#_get_accounts
func (c *Client) GetAccounts(index, size int) ([]*Account, error) {
	// validate input
	if index < 0 {
		return nil, fmt.Errorf("index cannot be less than 0")
	}
	if size > 100 {
		return nil, fmt.Errorf("size cannot be larger than 100")
	}

	// get accounts
	res, err := c.doRequest(http.MethodGet, fmt.Sprintf("/accounts/?index=%d&size=%d", index, size), nil)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	if err := json.Unmarshal(bodyBytes, &accounts); err != nil {
		return nil, err
	}

	// populate client
	for _, a := range accounts {
		a.client = c
	}
	return accounts, nil
}

// https://reseller.api.foxyproxy.com/#_get_accounts_by_username
func (c *Client) GetAccountsByUsername(username string, index, size int) ([]*Account, error) {
	// validate input
	if index < 0 {
		return nil, fmt.Errorf("index cannot be less than 0")
	}
	if size > 100 {
		return nil, fmt.Errorf("size cannot be larger than 100")
	}

	// get accounts
	res, err := c.doRequest(http.MethodGet, fmt.Sprintf("/accounts/%s/?index=%d&size=%d", username, index, size), nil)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	if err := json.Unmarshal(bodyBytes, &accounts); err != nil {
		return nil, err
	}
	return accounts, nil
}

// https://reseller.api.foxyproxy.com/#_get_accounts_by_node
func (c *Client) GetAccountsByNode(nodeName string, index, size int) ([]*Account, error) {
	return c.getAccountsByNode(nodeName, index, size)
}

func (c *Client) getAccountsByNode(nodeName string, index, size int) ([]*Account, error) {
	// validate input
	if index < 0 {
		return nil, fmt.Errorf("index cannot be less than 0")
	}
	if size > 100 {
		return nil, fmt.Errorf("size cannot be larger than 100")
	}

	// get accounts
	res, err := c.doRequest(http.MethodGet, fmt.Sprintf("/nodes/%s/accounts/?index=%d&size=%d", nodeName, index, size), nil)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	if err := json.Unmarshal(bodyBytes, &accounts); err != nil {
		return nil, err
	}
	return accounts, nil
}

// https://reseller.api.foxyproxy.com/#_count_accounts
func (c *Client) CountAccounts() (int, error) {
	type countRes struct {
		Count int `json:"count"`
	}
	res, err := c.doRequest(http.MethodGet, "/accounts/count/", nil)
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

// // https://reseller.api.foxyproxy.com/#_deactivate_accounts
func (c *Client) DeactivateAccount(username string) (int, error) {
	return c.deactivateAccount(username, nil)
}

func (c *Client) deactivateAccount(username string, params *commonAPIProperties) (int, error) {
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
	res, err := c.doRequest(http.MethodPatch, fmt.Sprintf("/accounts/deactivate/%s/", username), body)
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
func (c *Client) ActivateAccount(username string) (int, error) {
	return c.activateAccount(username, nil)
}

func (c *Client) activateAccount(username string, params *commonAPIProperties) (int, error) {
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
	res, err := c.doRequest(http.MethodPatch, fmt.Sprintf("/accounts/activate/%s/", username), body)
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
	return c.updatePassword(username, password, nil)
}

func (c *Client) updatePassword(username, password string, params *commonAPIProperties) (int, error) {
	// validate input
	if len(password) < 3 {
		return 0, fmt.Errorf("password must be more than 3 characters long")
	}
	if len(password) > 127 {
		return 0, fmt.Errorf("password must be less than 127 characters long")
	}

	type updatePasswordBody struct {
		Password string `json:"password"`
		*commonAPIProperties
	}
	upb := updatePasswordBody{
		Password:            password,
		commonAPIProperties: params,
	}
	type countRes struct {
		Count int `json:"count"`
	}
	jsonBody, err := json.Marshal(upb)
	if err != nil {
		return 0, err
	}
	res, err := c.doRequest(
		http.MethodPatch,
		fmt.Sprintf("/accounts/update-password/%s", username),
		jsonBody,
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

// https://reseller.api.foxyproxy.com/#_delete_accounts
func (c *Client) DeleteAccounts(username string, includeHistory bool) (int, error) {
	return c.deleteAccounts(username, includeHistory, nil)
}
func (c *Client) deleteAccounts(username string, includeHistory bool, params *commonAPIProperties) (int, error) {
	type body struct {
		IncludeHistory bool `json:"includeHistory"`
		*commonAPIProperties
	}
	b := body{
		IncludeHistory:      includeHistory,
		commonAPIProperties: params,
	}
	bJSON, err := json.Marshal(b)
	if err != nil {
		return 0, err
	}
	type countRes struct {
		Count int `json:"count"`
	}
	res, err := c.doRequest(http.MethodPatch, fmt.Sprintf("/accounts/activate/%s/", username), bJSON)
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

// https://reseller.api.foxyproxy.com/#_copy_accounts_from_one_node_to_others
func (c *Client) CopyAccounts(fromNode string, toNodes []string) (int, error) {
	params := commonAPIProperties{
		NodeNames: toNodes,
	}
	body, err := json.Marshal(params)
	if err != nil {
		return 0, err
	}
	type countRes struct {
		Count int `json:"count"`
	}
	res, err := c.doRequest(http.MethodPost, fmt.Sprintf("/accounts/copy-all/%s/", fromNode), body)
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
