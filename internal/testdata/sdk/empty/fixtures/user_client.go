// Generated by Fern. Do not edit.

package api

import (
	core "github.com/fern-api/fern-go/internal/testdata/sdk/empty/fixtures/core"
	strings "strings"
)

type UserClient interface {
}

func NewUserClient(baseURL string, httpClient core.HTTPClient, opts ...core.ClientOption) UserClient {
	options := new(core.ClientOptions)
	for _, opt := range opts {
		opt(options)
	}
	baseURL = strings.TrimRight(baseURL, "/")
	return &userClient{}
}

type userClient struct {
}