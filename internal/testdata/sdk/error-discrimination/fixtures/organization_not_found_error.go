// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/error/fixtures/core"
)

type OrganizationNotFoundError struct {
	*core.APIError
	Body *OrganizationNotFoundErrorBody
}

func (o *OrganizationNotFoundError) UnmarshalJSON(data []byte) error {
	body := new(OrganizationNotFoundErrorBody)
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	o.StatusCode = 404
	o.Body = body
	return nil
}

func (o *OrganizationNotFoundError) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.Body)
}
