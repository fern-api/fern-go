// This file was auto-generated by Fern from our API Definition.

package notificationclient

import (
	context "context"
	fmt "fmt"
	fixtures "github.com/fern-api/fern-go/internal/testdata/sdk/client-cycle/fixtures"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/client-cycle/fixtures/core"
	notification "github.com/fern-api/fern-go/internal/testdata/sdk/client-cycle/fixtures/user/notification"
	http "net/http"
)

type NotificationClient interface {
	List(ctx context.Context, userId string) ([]*notification.Notification, error)
	Foo(ctx context.Context, userId string, fooId string) (*fixtures.Foo, error)
}

func NewNotificationClient(opts ...core.ClientOption) NotificationClient {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &notificationClient{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type notificationClient struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func (n *notificationClient) List(ctx context.Context, userId string) ([]*notification.Notification, error) {
	baseURL := "https://api.foo.io/v1"
	if n.baseURL != "" {
		baseURL = n.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"users/%v/notifications", userId)

	var response []*notification.Notification
	if err := core.DoRequest(
		ctx,
		n.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		n.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (n *notificationClient) Foo(ctx context.Context, userId string, fooId string) (*fixtures.Foo, error) {
	baseURL := "https://api.foo.io/v1"
	if n.baseURL != "" {
		baseURL = n.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"users/%v/notifications/foo/%v", userId, fooId)

	response := new(fixtures.Foo)
	if err := core.DoRequest(
		ctx,
		n.httpClient,
		endpointURL,
		http.MethodPost,
		nil,
		&response,
		n.header,
		nil,
	); err != nil {
		return response, err
	}
	return response, nil
}
