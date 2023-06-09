// This file was auto-generated by Fern from our API Definition.

package api

import (
	uuid "github.com/gofrs/uuid/v5"
	time "time"
)

// GetUsersRequest is an in-lined request used by the GetUsername endpoint.
type GetUsersRequest struct {
	Id               uuid.UUID  `json:"-"`
	Date             time.Time  `json:"-"`
	Deadline         time.Time  `json:"-"`
	Bytes            []byte     `json:"-"`
	OptionalId       *uuid.UUID `json:"-"`
	OptionalDate     *time.Time `json:"-"`
	OptionalDeadline *time.Time `json:"-"`
	OptionalBytes    *[]byte    `json:"-"`
}
