// Generated by Fern. Do not edit.

package user

import (
	context "context"
	errors "errors"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/packages/fixtures/core"
	io "io"
	http "net/http"
)

type listEndpoint struct {
	url        string
	httpClient core.HTTPClient
	header     http.Header
}

func newListEndpoint(url string, httpClient core.HTTPClient, clientOptions *core.ClientOptions) *listEndpoint {
	return &listEndpoint{
		url:        url,
		httpClient: httpClient,
		header:     clientOptions.ToHeader(),
	}
}

func (l *listEndpoint) decodeError(statusCode int, body io.Reader) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return errors.New(string(bytes))
}

func (l *listEndpoint) Call(ctx context.Context) ([]*User, error) {
	endpointURL := l.url
	var response []*User
	if err := core.DoRequest(
		ctx,
		l.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		l.header,
		l.decodeError,
	); err != nil {
		return response, err
	}
	return response, nil
}