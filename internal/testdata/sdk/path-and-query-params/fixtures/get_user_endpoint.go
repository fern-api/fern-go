// Generated by Fern. Do not edit.

package api

import (
	context "context"
	errors "errors"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/path-and-query-params/fixtures/core"
	io "io"
	http "net/http"
)

type getUserEndpoint struct {
	url        string
	httpClient core.HTTPClient
	header     http.Header
}

func newGetUserEndpoint(url string, httpClient core.HTTPClient, clientOptions *core.ClientOptions) *getUserEndpoint {
	return &getUserEndpoint{
		url:        url,
		httpClient: httpClient,
		header:     clientOptions.ToHeader(),
	}
}

func (g *getUserEndpoint) decodeError(statusCode int, body io.Reader) error {
	bytes, err := io.ReadAll(body)
	if err != nil {
		return err
	}
	return errors.New(string(bytes))
}

func (g *getUserEndpoint) Call(ctx context.Context, userId string, request *GetUserRequest) (string, error) {
	endpointURL := fmt.Sprintf(g.url, userId)
	var response string
	if err := core.DoRequest(
		ctx,
		g.httpClient,
		endpointURL,
		http.MethodGet,
		request,
		response,
		g.header,
		g.decodeError,
	); err != nil {
		return response, err
	}
	return response, nil
}