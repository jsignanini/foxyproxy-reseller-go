package foxyproxy

import (
	"bytes"
	"fmt"
	"net/http"
)

func (c *Client) doRequest2(method, path string, body []byte) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(method, fmt.Sprintf("%s%s", c.endpointBaseURL, path), bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.username, c.password)
	req.Header.Add("Accept", ContentType)
	req.Header.Add("Content-Type", ContentType)
	req.Header.Add("X-DOMAIN", c.domainHeader)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	switch res.StatusCode {
	case http.StatusOK, http.StatusNotFound:
		return res, nil
	default:
		if res.StatusCode != http.StatusOK {
			apiError, err := NewError(res.Body)
			if err != nil {
				return nil, err
			}
			return nil, apiError
		}
		return res, nil
	}
	return res, nil
}

func (c *Client) doRequest(path string) (*http.Response, error) {
	client := http.Client{}
	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s%s", c.endpointBaseURL, path), nil)
	if err != nil {
		return nil, err
	}
	req.SetBasicAuth(c.username, c.password)
	req.Header.Add("Accept", ContentType)
	req.Header.Add("Content-Type", ContentType)
	req.Header.Add("X-DOMAIN", c.domainHeader)
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	switch res.StatusCode {
	case http.StatusOK, http.StatusNotFound:
		return res, nil
	default:
		if res.StatusCode != http.StatusOK {
			apiError, err := NewError(res.Body)
			if err != nil {
				return nil, err
			}
			return nil, apiError
		}
		return res, nil
	}
}
