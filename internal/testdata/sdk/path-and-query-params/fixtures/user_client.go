// Generated by Fern. Do not edit.

package api

import (
	context "context"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/path-and-query-params/fixtures/core"
	strings "strings"
)

type UserClient interface {
	GetUser(ctx context.Context, userId string, request *GetUserRequest) (string, error)
}

func NewUserClient(baseURL string, httpClient core.HTTPClient, opts ...core.ClientOption) UserClient {
	options := new(core.ClientOptions)
	for _, opt := range opts {
		opt(options)
	}
	baseURL = strings.TrimRight(baseURL, "/")
	return &userClient{
		getUserEndpoint: newGetUserEndpoint(baseURL+"/"+"users/%v", httpClient, options),
	}
}

type userClient struct {
	getUserEndpoint *getUserEndpoint
}

func (u *userClient) GetUser(ctx context.Context, userId string, request *GetUserRequest) (string, error) {
	return u.getUserEndpoint.Call(ctx, userId, request)
}