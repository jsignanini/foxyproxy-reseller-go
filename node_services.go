package foxyproxy

// NodeService is a node service.
// See https://reseller.api.foxyproxy.com/#_available_services.
type NodeService struct {
	Config string `json:"config,omitempty"`
	Name   string `json:"name"`
	Ports  []int  `json:"ports,omitempty"`
}
