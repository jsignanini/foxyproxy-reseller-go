package foxyproxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Client struct {
	Username, Password string
	fpURL              string
}

func NewClient(environment, username, password string) *Client {
	cl := &Client{
		Username: username,
		Password: password,
		fpURL:    "https://reseller.test.api.foxyproxy.com",
	}
	if environment == "production" {
		cl.fpURL = "https://ghostery.reseller.api.foxyproxy.com"
	}

	return cl
}

func (c *Client) GetActiveNodeConnectionsByAccount(nodeName string) ([]*NodeConnection, error) {
	return c.getActiveNodeConnectionsByAccount(nodeName)
}

func (c *Client) GetActiveNodeConnectionTotals(nodeName string) (int, error) {
	return c.getActiveNodeConnectionTotals(nodeName)
}

func (c *Client) GetAllNodes(index, size int) ([]*Node, error) {
	if index < 0 {
		return nil, fmt.Errorf("index cannot be less than 0")
	}
	if size > 100 {
		return nil, fmt.Errorf("size cannot be larger than 100")
	}
	res, err := c.doRequest(fmt.Sprintf("/nodes/?index=%d&size=%d", index, size))
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
	res, err := c.doRequest(fmt.Sprintf("/nodes/%s", nodeName))
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
	res, err := c.doRequest("/nodes/count/")
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
	res, err := c.doRequest(fmt.Sprintf("/nodes/%s/connections/", nodeName))
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
	res, err := c.doRequest("/nodes/dns-suffixes/")
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
	res, err := c.doRequest(fmt.Sprintf("/nodes/%s/connections/%d/%d/", nodeName, startTime.Unix(), endTime.Unix()))
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
