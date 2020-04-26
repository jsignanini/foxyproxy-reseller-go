package foxyproxy

type NodeService struct {
	Config string `json:"config,omitempty"`
	Name   string `json:"name"`
	Ports  []int  `json:"ports,omitempty"`
}
