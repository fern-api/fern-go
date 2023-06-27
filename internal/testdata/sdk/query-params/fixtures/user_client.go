// Generated by Fern. Do not edit.

package api

import (
	context "context"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/query-params/fixtures/core"
	strings "strings"
)

type UserClient interface {
	GetAllUsers(ctx context.Context, request *GetAllUsersRequest) (string, error)
}

func NewUserClient(baseURL string, httpClient core.HTTPClient, opts ...core.ClientOption) UserClient {
	options := new(core.ClientOptions)
	for _, opt := range opts {
		opt(options)
	}
	baseURL = strings.TrimRight(baseURL, "/")
	return &userClient{
		getAllUsersEndpoint: newGetAllUsersEndpoint(baseURL+"/"+"users/all", httpClient, options),
	}
}

type userClient struct {
	getAllUsersEndpoint *getAllUsersEndpoint
}

func (u *userClient) GetAllUsers(ctx context.Context, request *GetAllUsersRequest) (string, error) {
	return u.getAllUsersEndpoint.Call(ctx, request)
}