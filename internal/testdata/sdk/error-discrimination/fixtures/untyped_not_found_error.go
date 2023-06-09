// This file was auto-generated by Fern from our API Definition.

package api

import (
	core "github.com/fern-api/fern-go/internal/testdata/sdk/error/fixtures/core"
)

type UntypedNotFoundError struct {
	*core.APIError
}

func (u *UntypedNotFoundError) UnmarshalJSON(data []byte) error {
	u.StatusCode = 404
	return nil
}

func (u *UntypedNotFoundError) MarshalJSON() ([]byte, error) {
	return nil, nil
}
