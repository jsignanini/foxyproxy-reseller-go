package foxyproxy

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
)

// Error is an api error object.
type Error struct {
	Timestamp   string `json:"timestamp"`
	Status      int    `json:"status"`
	ErrorString string `json:"error"`
	Message     string `json:"message"`
	Path        string `json:"path"`
}

// Error returns a string representation of Error.
func (e *Error) Error() string {
	errBytes, err := json.MarshalIndent(e, "", "\t")
	if err != nil {
		return err.Error()
	}
	return string(errBytes)
}

// NewError generates a new Error object.
func NewError(body io.ReadCloser) (*Error, error) {
	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, err
	}
	e := &Error{}
	if err := json.Unmarshal(bodyBytes, e); err != nil {
		return nil, fmt.Errorf(string(bodyBytes))
	}
	return e, nil
}
