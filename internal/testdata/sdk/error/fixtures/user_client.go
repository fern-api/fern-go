// Generated by Fern. Do not edit.

package api

import (
	context "context"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/error/fixtures/core"
	strings "strings"
)

type UserClient interface {
	Get(ctx context.Context, id string) (string, error)
	Update(ctx context.Context, id string, request string) (string, error)
}

func NewUserClient(baseURL string, httpClient core.HTTPClient, opts ...core.ClientOption) UserClient {
	options := new(core.ClientOptions)
	for _, opt := range opts {
		opt(options)
	}
	baseURL = strings.TrimRight(baseURL, "/")
	return &userClient{
		getEndpoint:    newGetEndpoint(baseURL+"/"+"%v", httpClient, options),
		updateEndpoint: newUpdateEndpoint(baseURL+"/"+"%v", httpClient, options),
	}
}

type userClient struct {
	getEndpoint    *getEndpoint
	updateEndpoint *updateEndpoint
}

func (u *userClient) Get(ctx context.Context, id string) (string, error) {
	return u.getEndpoint.Call(ctx, id)
}

func (u *userClient) Update(ctx context.Context, id string, request string) (string, error) {
	return u.updateEndpoint.Call(ctx, id, request)
}