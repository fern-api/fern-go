package client

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	// acceptHeader is the Accept header.
	acceptHeader = "Accept"

	// contentType specifies the JSON Content-Type header value.
	contentType       = "application/json"
	contentTypeHeader = "Content-Type"

	// fernLanguage specifies the value of the X-Fern-Language header.
	fernLanguage       = "go"
	fernLanguageHeader = "X-Fern-Language"

	// fernSDKName specifies the name of this Fern SDK.
	fernSDKName       = "fern-go-sdk"
	fernSDKNameHeader = "X-Fern-SDK-Name"

	// fernSDKVersion specifies the version of this Fern SDK.
	fernSDKVersion       = "0.0.1"
	fernSDKVersionHeader = "X-Fern-SDK-Version"
)

// fernHeaders specifies all of the standard Fern headers in
// a set so that they're easier to access and reference.
var fernHeaders = map[string]string{
	acceptHeader:         contentType,
	contentTypeHeader:    contentType,
	fernLanguageHeader:   fernLanguage,
	fernSDKNameHeader:    fernSDKName,
	fernSDKVersionHeader: fernSDKVersion,
}

// HTTPClient is an interface for a subset of the *http.Client.
type HTTPClient interface {
	Do(*http.Request) (*http.Response, error)
}

// ClientOption adapts the behavior of a Fern client.
type ClientOption func(*clientOptions)

type clientOptions struct{}

// doRequest issues a JSON request to the given url.
func doRequest(
	ctx context.Context,
	client HTTPClient,
	url string,
	method string,
	request any,
	response any,
	endpointHeaders http.Header,
	errorDecoder func(int, io.Reader) error,
) error {
	var requestBody io.Reader
	if request != nil {
		requestBytes, err := json.Marshal(request)
		if err != nil {
			return err
		}
		requestBody = bytes.NewReader(requestBytes)
	}
	req, err := newRequest(ctx, url, method, endpointHeaders, requestBody)
	if err != nil {
		return err
	}

	// If the call has been cancelled, don't issue the request.
	if err := ctx.Err(); err != nil {
		return err
	}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	// Close the response body after we're done.
	defer resp.Body.Close()

	// Check if the call was cancelled before we return the error
	// associated with the call and/or unmarshal the response data.
	if err = ctx.Err(); err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		if errorDecoder != nil {
			// This endpoint has custom errors, so we'll
			// attempt to unmarshal the error into a structured
			// type based on the status code.
			return errorDecoder(resp.StatusCode, resp.Body)
		}
		// This endpoint doesn't have any custom error
		// types, so we just read the body as-is, and
		// put it into a normal error.
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return errors.New(string(bytes))
	}

	// Mutate the response parameter in-place.
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(response); err != nil {
		return err
	}

	return nil
}

// newRequest returns a new *http.Request with all of the fields
// required to issue the call.
func newRequest(
	ctx context.Context,
	url string,
	method string,
	endpointHeaders http.Header,
	requestBody io.Reader,
) (*http.Request, error) {
	req, err := http.NewRequest(method, url, requestBody)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	for name, value := range fernHeaders {
		req.Header.Set(name, value)
	}
	for name, values := range endpointHeaders {
		req.Header[name] = values
	}
	return req, nil
}
