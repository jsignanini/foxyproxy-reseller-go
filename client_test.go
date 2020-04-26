package foxyproxy

import "testing"

func TestNewClient(t *testing.T) {
	const (
		domainHeader    = "example-inc"
		endpointBaseURL = "https://reseller.example-inc.api.foxyproxy.com"
		username        = "admin"
		password        = "12345"
	)
	ncp := NewClientParams{
		DomainHeader:    domainHeader,
		EndpointBaseURL: endpointBaseURL,
		Username:        username,
		Password:        password,
	}
	c := NewClient(&ncp)
	if c.domainHeader != domainHeader {
		t.Errorf("expected client domain header: %s, got %s", domainHeader, c.domainHeader)
	}
	if c.endpointBaseURL != endpointBaseURL {
		t.Errorf("expected client endpoint base url: %s, got %s", endpointBaseURL, c.endpointBaseURL)
	}
	if c.username != username {
		t.Errorf("expected client username: %s, got %s", username, c.username)
	}
	if c.password != password {
		t.Errorf("expected client password: %s, got %s", password, c.password)
	}
}
