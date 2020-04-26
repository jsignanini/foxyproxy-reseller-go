package foxyproxy

// Account represents a customer who can be assigned to one or more nodes (vpn/proxy servers).
// See https://reseller.api.foxyproxy.com/#_accounts.
type Account struct {
	Active   bool   `json:"active"`
	Node     *Node  `json:"node,omitempty"`
	UID      string `json:"uid"`
	Username string `json:"username"`

	client *Client
}

// NewAccount generates a new account object.
func NewAccount(c *Client) *Account {
	return &Account{
		client: c,
	}
}

// GetNodeNames returns a slice of length 1 containing the account's node name.
// Returns an empty slice if node is nil.
func (a *Account) GetNodeNames() []string {
	nodeNames := []string{}
	if a.Node != nil {
		nodeNames = append(nodeNames, a.Node.Name)
	}
	return nodeNames
}

// Deactivate deactivates the account on it's node and returns a count of affected accounts.
// See https://reseller.api.foxyproxy.com/#_deactivate_accounts.
func (a *Account) Deactivate() (int, error) {
	return a.client.deactivateAccount(a.Username, &CommonProperties{
		NodeNames: a.GetNodeNames(),
	})
}

// Activate activates the account on it's node and returns a count of affected accounts.
// See https://reseller.api.foxyproxy.com/#_activate_accounts.
func (a *Account) Activate() (int, error) {
	return a.client.activateAccount(a.Username, &CommonProperties{
		NodeNames: a.GetNodeNames(),
	})
}

// UpdatePassword updates the password on it's node and returns a count of affected accounts.
// See https://reseller.api.foxyproxy.com/#_update_passwords.
func (a *Account) UpdatePassword(password string) (int, error) {
	return a.client.updatePassword(a.Username, password, &CommonProperties{
		NodeNames: a.GetNodeNames(),
	})
}

// DeleteAccountsParams is an object of optional account deletion parameters.
type DeleteAccountsParams struct {
	IncludeHistory bool `json:"includeHistory"`
	CommonProperties
}

// Delete deletes the account on it's node. If includeHistory is set to true, account history is
// also deleted on it's node. Returns a count of affected accounts.
// See https://reseller.api.foxyproxy.com/#_delete_accounts.
func (a *Account) Delete(includeHistory bool) (int, error) {
	return a.client.deleteAccounts(a.Username, includeHistory, &CommonProperties{
		NodeNames: a.GetNodeNames(),
	})
}
