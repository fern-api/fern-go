// This file was auto-generated by Fern from our API Definition.

package userclient

import (
	context "context"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/client-cycle/fixtures/core"
	user "github.com/fern-api/fern-go/internal/testdata/sdk/client-cycle/fixtures/user"
	notificationclient "github.com/fern-api/fern-go/internal/testdata/sdk/client-cycle/fixtures/user/notification/notificationclient"
	http "net/http"
)

type Client interface {
	GetUser(ctx context.Context, userId string) (*user.User, error)
	Notification() notificationclient.Client
	User() UserClient
}

func NewClient(opts ...core.ClientOption) Client {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &client{
		baseURL:            options.BaseURL,
		httpClient:         options.HTTPClient,
		header:             options.ToHeader(),
		notificationClient: notificationclient.NewClient(opts...),
		userClient:         NewUserClient(opts...),
	}
}

type client struct {
	baseURL            string
	httpClient         core.HTTPClient
	header             http.Header
	notificationClient notificationclient.Client
	userClient         UserClient
}

func (c *client) GetUser(ctx context.Context, userId string) (*user.User, error) {
	baseURL := "https://api.foo.io/v1"
	if c.baseURL != "" {
		baseURL = c.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"users/%v", userId)

	response := new(user.User)
	if err := core.DoRequest(
		ctx,
		c.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		c.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (c *client) Notification() notificationclient.Client {
	return c.notificationClient
}

func (c *client) User() UserClient {
	return c.userClient
}
