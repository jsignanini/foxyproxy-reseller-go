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

	client *Client

	// TODO
	// Services
}

func NewNode(c *Client) *Node {
	return &Node{
		client: c,
	}
}

func (n *Node) GetClient() *Client {
	if n.client != nil {
		return n.client
	}
	return NewClient(&NewClientParams{})
}

func (n *Node) GetActiveConnectionsByAccount() ([]*NodeConnection, error) {
	return n.GetClient().getActiveNodeConnectionsByAccount(n.Name)
}

func (n *Node) GetActiveConnectionTotals() (int, error) {
	return n.GetClient().getActiveNodeConnectionTotals(n.Name)
}

func (n *Node) GetHistoricalConnectionsByAccount(startTime, endTime time.Time) ([]*NodeConnection, error) {
	return n.GetClient().getHistoricalNodeConnectionsByAccount(n.Name, startTime, endTime)
}

func (n *Node) GetHistoricalConnectionTotals(startTime, endTime time.Time) (int, error) {
	return n.GetClient().getHistoricalNodeConnectionTotals(n.Name, startTime, endTime)
}

func (n *Node) GetTrafficByAccount(startTime, endTime time.Time) ([]*NodeTrafficAccount, error) {
	return n.GetClient().getNodeTrafficByAccount(n.Name, startTime, endTime)
}

func (n *Node) GetTrafficTotals(startTime, endTime time.Time) (*NodeTrafficTotals, error) {
	return n.GetClient().getNodeTrafficTotals(n.Name, startTime, endTime)
}
