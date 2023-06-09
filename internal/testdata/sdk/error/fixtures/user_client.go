// This file was auto-generated by Fern from our API Definition.

package api

import (
	bytes "bytes"
	context "context"
	json "encoding/json"
	errors "errors"
	fmt "fmt"
	core "github.com/fern-api/fern-go/internal/testdata/sdk/error/fixtures/core"
	io "io"
	http "net/http"
)

type UserClient interface {
	Get(ctx context.Context, id string) (string, error)
	Update(ctx context.Context, id string, request string) (string, error)
}

func NewUserClient(opts ...core.ClientOption) UserClient {
	options := core.NewClientOptions()
	for _, opt := range opts {
		opt(options)
	}
	return &userClient{
		baseURL:    options.BaseURL,
		httpClient: options.HTTPClient,
		header:     options.ToHeader(),
	}
}

type userClient struct {
	baseURL    string
	httpClient core.HTTPClient
	header     http.Header
}

func (u *userClient) Get(ctx context.Context, id string) (string, error) {
	baseURL := ""
	if u.baseURL != "" {
		baseURL = u.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"%v", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 404:
			value := new(UserNotFoundError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 501:
			value := new(NotImplementedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 418:
			value := new(TeapotError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 426:
			value := new(UpgradeError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 400:
			value := new(UntypedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response string
	if err := core.DoRequest(
		ctx,
		u.httpClient,
		endpointURL,
		http.MethodGet,
		nil,
		&response,
		u.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}

func (u *userClient) Update(ctx context.Context, id string, request string) (string, error) {
	baseURL := ""
	if u.baseURL != "" {
		baseURL = u.baseURL
	}
	endpointURL := fmt.Sprintf(baseURL+"/"+"%v", id)

	errorDecoder := func(statusCode int, body io.Reader) error {
		raw, err := io.ReadAll(body)
		if err != nil {
			return err
		}
		apiError := core.NewAPIError(statusCode, errors.New(string(raw)))
		decoder := json.NewDecoder(bytes.NewReader(raw))
		switch statusCode {
		case 426:
			value := new(UpgradeError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		case 400:
			value := new(UntypedError)
			value.APIError = apiError
			if err := decoder.Decode(value); err != nil {
				return err
			}
			return value
		}
		return apiError
	}

	var response string
	if err := core.DoRequest(
		ctx,
		u.httpClient,
		endpointURL,
		http.MethodPost,
		request,
		&response,
		u.header,
		errorDecoder,
	); err != nil {
		return response, err
	}
	return response, nil
}
