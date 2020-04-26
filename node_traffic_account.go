package foxyproxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// NodeTrafficAccount is a count of a node's traffic.
type NodeTrafficAccount struct {
	UID         string
	Active      bool
	Username    string
	TrafficDown float64
	TrafficUp   float64
	TrafficAll  float64
}

func (c *Client) getNodeTrafficByAccount(nodeName string, startTime, endTime time.Time) ([]*NodeTrafficAccount, error) {
	res, err := c.doRequest(http.MethodGet, fmt.Sprintf("/nodes/%s/traffic-by-account/%d/%d", nodeName, startTime.Unix(), endTime.Unix()), nil)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(string(bodyBytes))
	}
	traffics := []*NodeTrafficAccount{}
	if err := json.Unmarshal(bodyBytes, &traffics); err != nil {
		return nil, err
	}
	return traffics, nil
}
