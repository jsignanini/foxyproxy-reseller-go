package foxyproxy

import (
	"fmt"
	"net/http"
)

func (c *Client) doRequest(path string) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("https://ghostery.reseller.api.foxyproxy.com%s", path), nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.Username, c.Password)
	req.Header.Add("Accept", ContentType)
	req.Header.Add("Content-Type", ContentType)
	req.Header.Add("X-DOMAIN", "ghostery")
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	if res.StatusCode != http.StatusOK {
		apiError, err := NewError(res.Body)
		if err != nil {
			return nil, err
		}
		return nil, apiError
	}
	return res, nil
}
