// This file was auto-generated by Fern from our API Definition.

package notification

import (
	context "context"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/core"
	http "net/http"
)

type NotificationClient interface {
	List(ctx context.Context, userId string) ([]*Notification, error)
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

func (n *notificationClient) List(ctx context.Context, userId string) ([]*Notification, error) {
	baseURL := "https://api.foo.io/v1"
	if n.baseURL != "" {
		baseURL = n.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"users/%v/notifications", userId)

	var response []*Notification
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
