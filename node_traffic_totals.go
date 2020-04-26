package foxyproxy

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// NodeTrafficTotals is a total traffic count for a node.
type NodeTrafficTotals struct {
	TrafficDown float64
	TrafficUp   float64
	TrafficAll  float64
	Quota       float64
}

func (c *Client) getNodeTrafficTotals(nodeName string, startTime, endTime time.Time) (*NodeTrafficTotals, error) {
	res, err := c.doRequest(http.MethodGet, fmt.Sprintf("/nodes/%s/traffic/%d/%d", nodeName, startTime.Unix(), endTime.Unix()), nil)
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
	traffic := &NodeTrafficTotals{}
	if err := json.Unmarshal(bodyBytes, traffic); err != nil {
		return nil, err
	}
	return traffic, nil
}
