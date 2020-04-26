package foxyproxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// NodeConnection is a node connection.
type NodeConnection struct {
	UID         string
	Active      bool
	Username    string
	Connections int
}

func (c *Client) getActiveNodeConnectionsByAccount(nodeName string) ([]*NodeConnection, error) {
	res, err := c.doRequest(http.MethodGet, fmt.Sprintf("/nodes/%s/connections-by-account/", nodeName), nil)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	fmt.Println(string(bodyBytes))
	connections := []*NodeConnection{}
	if err := json.Unmarshal(bodyBytes, &connections); err != nil {
		return nil, err
	}
	return connections, nil
}

func (c *Client) getHistoricalNodeConnectionsByAccount(nodeName string, startTime, endTime time.Time) ([]*NodeConnection, error) {
	res, err := c.doRequest(http.MethodGet, fmt.Sprintf("/nodes/%s/connections-by-account/%d/%d/", nodeName, startTime.Unix(), endTime.Unix()), nil)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	connections := []*NodeConnection{}
	if err := json.Unmarshal(bodyBytes, &connections); err != nil {
		return nil, err
	}
	return connections, nil
}
