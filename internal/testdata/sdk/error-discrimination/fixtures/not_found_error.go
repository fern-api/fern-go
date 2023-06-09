// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/error/fixtures/core"
)

type NotFoundError struct {
	*core.APIError
	Body string
}

func (n *NotFoundError) UnmarshalJSON(data []byte) error {
	var body string
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	n.StatusCode = 404
	n.Body = body
	return nil
}

func (n *NotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(n.Body)
}
