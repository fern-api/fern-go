// This file was auto-generated by Fern from our API Definition.

package api

import (
	json "encoding/json"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/error/fixtures/core"
)

type UpgradeError struct {
	*core.APIError
	Body string
}

func (u *UpgradeError) UnmarshalJSON(data []byte) error {
	var body string
	if err := json.Unmarshal(data, &body); err != nil {
		return err
	}
	if body != "upgrade" {
		return fmt.Errorf("expected literal %q, but found %q", "upgrade", body)
	}
	u.StatusCode = 426
	u.Body = body
	return nil
}

func (u *UpgradeError) MarshalJSON() ([]byte, error) {
	return json.Marshal("upgrade")
}
