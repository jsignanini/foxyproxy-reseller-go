package foxyproxy

// CommonProperties is optional JSON that account operations which write/change data (create,
// update, delete) accept.
// See https://reseller.api.foxyproxy.com/#_common_api_properties.
type CommonProperties struct {
	Comment   string   `json:"comment,omitempty"`
	NodeNames []string `json:"nodeNames,omitempty"`
}
