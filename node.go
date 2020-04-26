package foxyproxy

import (
	"time"
)

// Node represents a node object.
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

// NewNode generates a new node object.
func NewNode(c *Client) *Node {
	return &Node{
		client: c,
	}
}

// GetActiveConnectionsByAccount gets the active connections for the node. This does not include
// closed connections.
// See https://reseller.api.foxyproxy.com/#_active_node_connections_by_account.
func (n *Node) GetActiveConnectionsByAccount() ([]*NodeConnection, error) {
	return n.client.getActiveNodeConnectionsByAccount(n.Name)
}

// GetActiveConnectionTotals gets a count of active connections for the node.
// See https://reseller.api.foxyproxy.com/#_active_node_connection_totals.
func (n *Node) GetActiveConnectionTotals() (int, error) {
	return n.client.getActiveNodeConnectionTotals(n.Name)
}

// GetHistoricalConnectionsByAccount gets the connections for the node between startTime and
// endTime, inclusive. This does not include active connections.
// See https://reseller.api.foxyproxy.com/#_historical_node_connections_by_account.
func (n *Node) GetHistoricalConnectionsByAccount(startTime, endTime time.Time) ([]*NodeConnection, error) {
	return n.client.getHistoricalNodeConnectionsByAccount(n.Name, startTime, endTime)
}

// GetHistoricalConnectionTotals gets a count of connections for the node between startTime and
// endtTime, inclusive. This does not include active connections.
// See https://reseller.api.foxyproxy.com/#_historical_node_connection_totals.
func (n *Node) GetHistoricalConnectionTotals(startTime, endTime time.Time) (int, error) {
	return n.client.getHistoricalNodeConnectionTotals(n.Name, startTime, endTime)
}

// GetTrafficByAccount gets various traffic counts and last authentication info for all accounts on
// the node between two startTime and endTime, inclusive.
// See https://reseller.api.foxyproxy.com/#_node_traffic_by_account.
func (n *Node) GetTrafficByAccount(startTime, endTime time.Time) ([]*NodeTrafficAccount, error) {
	return n.client.getNodeTrafficByAccount(n.Name, startTime, endTime)
}

// GetTrafficTotals gets various traffic counts for the node between startTime and endTime,
// inclusive.
// See https://reseller.api.foxyproxy.com/#_node_traffic_totals.
func (n *Node) GetTrafficTotals(startTime, endTime time.Time) (*NodeTrafficTotals, error) {
	return n.client.getNodeTrafficTotals(n.Name, startTime, endTime)
}

// GetAccountsByNode gets all accounts for the node.
// See https://reseller.api.foxyproxy.com/#_get_accounts_by_node.
func (n *Node) GetAccountsByNode(index, size int) ([]*Account, error) {
	return n.client.getAccountsByNode(n.Name, index, size)
}
