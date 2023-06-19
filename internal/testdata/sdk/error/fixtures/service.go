// Generated by Fern. Do not edit.

package api

import (
	context "context"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/error/fixtures/core"
)

type Service interface {
	Get(ctx context.Context, id string) (string, error)
	Update(ctx context.Context, id string, request string) (string, error)
}

func NewClient(baseURL string, httpClient core.HTTPClient, opts ...core.ClientOption) (Service, error) {
	options := new(core.ClientOptions)
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		getEndpoint:    newGetEndpoint(baseURL, httpClient, options),
		updateEndpoint: newUpdateEndpoint(baseURL, httpClient, options),
	}, nil
}

type client struct {
	getEndpoint    *getEndpoint
	updateEndpoint *updateEndpoint
}

func (g *client) Get(ctx context.Context, id string) (string, error) {
	return g.getEndpoint.Call(ctx, id)
}

func (u *client) Update(ctx context.Context, id string, request string) (string, error) {
	return u.updateEndpoint.Call(ctx, id, request)
}