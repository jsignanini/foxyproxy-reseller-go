package foxyproxy

import (
	"time"
)

type Node struct {
	Active      bool
	Name        string
	IPAddress   string
	Country     string
	CountryCode string
	City        string
	Services    []*NodeService

	client *Client
}

func NewNode(c *Client) *Node {
	return &Node{
		client: c,
	}
}

func (n *Node) GetActiveConnectionsByAccount() ([]*NodeConnection, error) {
	return n.client.getActiveNodeConnectionsByAccount(n.Name)
}

func (n *Node) GetActiveConnectionTotals() (int, error) {
	return n.client.getActiveNodeConnectionTotals(n.Name)
}

func (n *Node) GetHistoricalConnectionsByAccount(startTime, endTime time.Time) ([]*NodeConnection, error) {
	return n.client.getHistoricalNodeConnectionsByAccount(n.Name, startTime, endTime)
}

func (n *Node) GetHistoricalConnectionTotals(startTime, endTime time.Time) (int, error) {
	return n.client.getHistoricalNodeConnectionTotals(n.Name, startTime, endTime)
}

func (n *Node) GetTrafficByAccount(startTime, endTime time.Time) ([]*NodeTrafficAccount, error) {
	return n.client.getNodeTrafficByAccount(n.Name, startTime, endTime)
}

func (n *Node) GetTrafficTotals(startTime, endTime time.Time) (*NodeTrafficTotals, error) {
	return n.client.getNodeTrafficTotals(n.Name, startTime, endTime)
}
