// Code generated by zanzibar
// @generated

// Copyright (c) 2017 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package barClient

import (
	"context"
	"strconv"

	"github.com/pkg/errors"
	clientsBarBar "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/bar/bar"
	"github.com/uber/zanzibar/runtime"
)

// Client defines bar client interface.
type Client interface {
	HTTPClient() *zanzibar.HTTPClient
	ArgNotStruct(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBarBar.Bar_ArgNotStruct_Args,
	) (map[string]string, error)
	ArgWithHeaders(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBarBar.Bar_ArgWithHeaders_Args,
	) (*clientsBarBar.BarResponse, map[string]string, error)
	MissingArg(
		ctx context.Context,
		reqHeaders map[string]string,
	) (*clientsBarBar.BarResponse, map[string]string, error)
	NoRequest(
		ctx context.Context,
		reqHeaders map[string]string,
	) (*clientsBarBar.BarResponse, map[string]string, error)
	Normal(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBarBar.Bar_Normal_Args,
	) (*clientsBarBar.BarResponse, map[string]string, error)
	TooManyArgs(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBarBar.Bar_TooManyArgs_Args,
	) (*clientsBarBar.BarResponse, map[string]string, error)
	Echo(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsBarBar.Echo_Echo_Args,
	) (string, map[string]string, error)
}

// NewClient returns a new http client.
func NewClient(gateway *zanzibar.Gateway) Client {
	ip := gateway.Config.MustGetString("clients.bar.ip")
	port := gateway.Config.MustGetInt("clients.bar.port")

	baseURL := "http://" + ip + ":" + strconv.Itoa(int(port))
	return &barClient{
		clientID:   "bar",
		httpClient: zanzibar.NewHTTPClient(gateway, baseURL),
	}
}

// barClient is the http client.
type barClient struct {
	clientID   string
	httpClient *zanzibar.HTTPClient
}

// HTTPClient returns the underlying HTTP client, should only be
// used for internal testing.
func (c *barClient) HTTPClient() *zanzibar.HTTPClient {
	return c.httpClient
}

// ArgNotStruct calls "/arg-not-struct-path" endpoint.
func (c *barClient) ArgNotStruct(
	ctx context.Context,
	headers map[string]string,
	r *clientsBarBar.Bar_ArgNotStruct_Args,
) (map[string]string, error) {
	req := zanzibar.NewClientHTTPRequest(
		c.clientID, "argNotStruct", c.httpClient,
	)

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/arg-not-struct-path"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200, 403})

	switch res.StatusCode {
	case 200:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return respHeaders, err
		}

		return respHeaders, nil

	case 403:
		var exception clientsBarBar.BarException
		err = res.ReadAndUnmarshalBody(&exception)
		if err != nil {
			return respHeaders, err
		}
		return respHeaders, &exception

	default:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return respHeaders, err
		}
	}

	return respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// ArgWithHeaders calls "/bar/argWithHeaders" endpoint.
func (c *barClient) ArgWithHeaders(
	ctx context.Context,
	headers map[string]string,
	r *clientsBarBar.Bar_ArgWithHeaders_Args,
) (*clientsBarBar.BarResponse, map[string]string, error) {
	var defaultRes *clientsBarBar.BarResponse
	req := zanzibar.NewClientHTTPRequest(
		c.clientID, "argWithHeaders", c.httpClient,
	)
	// TODO(jakev): Ensure we validate mandatory headers

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/bar" + "/argWithHeaders"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return defaultRes, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return defaultRes, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200})

	switch res.StatusCode {
	case 200:
		var responseBody clientsBarBar.BarResponse
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return defaultRes, respHeaders, err
		}
		// TODO(jakev): read response headers and put them in body

		return &responseBody, respHeaders, nil
	}

	return defaultRes, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// ArgWithManyQueryParams calls "/bar/argWithManyQueryParams" endpoint.
func (c *BarClient) ArgWithManyQueryParams(
	ctx context.Context,
	headers map[string]string,
	r *clientsBarBar.Bar_ArgWithManyQueryParams_Args,
) (*clientsBarBar.BarResponse, map[string]string, error) {
	var defaultRes *clientsBarBar.BarResponse
	req := zanzibar.NewClientHTTPRequest(
		c.ClientID, "argWithManyQueryParams", c.HTTPClient,
	)

	// Generate full URL.
	fullURL := c.HTTPClient.BaseURL + "/bar" + "/argWithManyQueryParams"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return defaultRes, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return defaultRes, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200})

	switch res.StatusCode {
	case 200:
		var responseBody clientsBarBar.BarResponse
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return defaultRes, respHeaders, err
		}
		// TODO(jakev): read response headers and put them in body

		return &responseBody, respHeaders, nil
	}

	return defaultRes, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// ArgWithQueryHeader calls "/bar/argWithQueryHeader" endpoint.
func (c *BarClient) ArgWithQueryHeader(
	ctx context.Context,
	headers map[string]string,
	r *clientsBarBar.Bar_ArgWithQueryHeader_Args,
) (*clientsBarBar.BarResponse, map[string]string, error) {
	var defaultRes *clientsBarBar.BarResponse
	req := zanzibar.NewClientHTTPRequest(
		c.ClientID, "argWithQueryHeader", c.HTTPClient,
	)

	// Generate full URL.
	fullURL := c.HTTPClient.BaseURL + "/bar" + "/argWithQueryHeader"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return defaultRes, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return defaultRes, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200})

	switch res.StatusCode {
	case 200:
		var responseBody clientsBarBar.BarResponse
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return defaultRes, respHeaders, err
		}
		// TODO(jakev): read response headers and put them in body

		return &responseBody, respHeaders, nil
	}

	return defaultRes, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// ArgWithQueryParams calls "/bar/argWithQueryParams" endpoint.
func (c *BarClient) ArgWithQueryParams(
	ctx context.Context,
	headers map[string]string,
	r *clientsBarBar.Bar_ArgWithQueryParams_Args,
) (*clientsBarBar.BarResponse, map[string]string, error) {
	var defaultRes *clientsBarBar.BarResponse
	req := zanzibar.NewClientHTTPRequest(
		c.ClientID, "argWithQueryParams", c.HTTPClient,
	)

	// Generate full URL.
	fullURL := c.HTTPClient.BaseURL + "/bar" + "/argWithQueryParams"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return defaultRes, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return defaultRes, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200})

	switch res.StatusCode {
	case 200:
		var responseBody clientsBarBar.BarResponse
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return defaultRes, respHeaders, err
		}
		// TODO(jakev): read response headers and put them in body

		return &responseBody, respHeaders, nil
	}

	return defaultRes, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// MissingArg calls "/missing-arg-path" endpoint.
func (c *barClient) MissingArg(
	ctx context.Context,
	headers map[string]string,
) (*clientsBarBar.BarResponse, map[string]string, error) {
	var defaultRes *clientsBarBar.BarResponse
	req := zanzibar.NewClientHTTPRequest(
		c.clientID, "missingArg", c.httpClient,
	)

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/missing-arg-path"

	err := req.WriteJSON("GET", fullURL, headers, nil)
	if err != nil {
		return defaultRes, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return defaultRes, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200, 403})

	switch res.StatusCode {
	case 200:
		var responseBody clientsBarBar.BarResponse
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return defaultRes, respHeaders, err
		}
		// TODO(jakev): read response headers and put them in body

		return &responseBody, respHeaders, nil

	case 403:
		var exception clientsBarBar.BarException
		err = res.ReadAndUnmarshalBody(&exception)
		if err != nil {
			return defaultRes, respHeaders, err
		}
		return defaultRes, respHeaders, &exception

	default:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return defaultRes, respHeaders, err
		}
	}

	return defaultRes, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// NoRequest calls "/no-request-path" endpoint.
func (c *barClient) NoRequest(
	ctx context.Context,
	headers map[string]string,
) (*clientsBarBar.BarResponse, map[string]string, error) {
	var defaultRes *clientsBarBar.BarResponse
	req := zanzibar.NewClientHTTPRequest(
		c.clientID, "noRequest", c.httpClient,
	)

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/no-request-path"

	err := req.WriteJSON("GET", fullURL, headers, nil)
	if err != nil {
		return defaultRes, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return defaultRes, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200, 403})

	switch res.StatusCode {
	case 200:
		var responseBody clientsBarBar.BarResponse
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return defaultRes, respHeaders, err
		}
		// TODO(jakev): read response headers and put them in body

		return &responseBody, respHeaders, nil

	case 403:
		var exception clientsBarBar.BarException
		err = res.ReadAndUnmarshalBody(&exception)
		if err != nil {
			return defaultRes, respHeaders, err
		}
		return defaultRes, respHeaders, &exception

	default:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return defaultRes, respHeaders, err
		}
	}

	return defaultRes, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// Normal calls "/bar-path" endpoint.
func (c *barClient) Normal(
	ctx context.Context,
	headers map[string]string,
	r *clientsBarBar.Bar_Normal_Args,
) (*clientsBarBar.BarResponse, map[string]string, error) {
	var defaultRes *clientsBarBar.BarResponse
	req := zanzibar.NewClientHTTPRequest(
		c.clientID, "normal", c.httpClient,
	)

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/bar-path"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return defaultRes, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return defaultRes, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200, 403})

	switch res.StatusCode {
	case 200:
		var responseBody clientsBarBar.BarResponse
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return defaultRes, respHeaders, err
		}
		// TODO(jakev): read response headers and put them in body

		return &responseBody, respHeaders, nil

	case 403:
		var exception clientsBarBar.BarException
		err = res.ReadAndUnmarshalBody(&exception)
		if err != nil {
			return defaultRes, respHeaders, err
		}
		return defaultRes, respHeaders, &exception

	default:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return defaultRes, respHeaders, err
		}
	}

	return defaultRes, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// TooManyArgs calls "/too-many-args-path" endpoint.
func (c *barClient) TooManyArgs(
	ctx context.Context,
	headers map[string]string,
	r *clientsBarBar.Bar_TooManyArgs_Args,
) (*clientsBarBar.BarResponse, map[string]string, error) {
	var defaultRes *clientsBarBar.BarResponse
	req := zanzibar.NewClientHTTPRequest(
		c.clientID, "tooManyArgs", c.httpClient,
	)

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/too-many-args-path"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return defaultRes, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return defaultRes, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200, 403})

	switch res.StatusCode {
	case 200:
		var responseBody clientsBarBar.BarResponse
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return defaultRes, respHeaders, err
		}
		// TODO(jakev): read response headers and put them in body

		return &responseBody, respHeaders, nil

	case 403:
		var exception clientsBarBar.BarException
		err = res.ReadAndUnmarshalBody(&exception)
		if err != nil {
			return defaultRes, respHeaders, err
		}
		return defaultRes, respHeaders, &exception

	default:
		// TODO: log about unexpected body bytes?
		_, err = res.ReadAll()
		if err != nil {
			return defaultRes, respHeaders, err
		}
	}

	return defaultRes, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}

// Echo calls "/echo" endpoint.
func (c *barClient) Echo(
	ctx context.Context,
	headers map[string]string,
	r *clientsBarBar.Echo_Echo_Args,
) (string, map[string]string, error) {
	var defaultRes string
	req := zanzibar.NewClientHTTPRequest(
		c.clientID, "echo", c.httpClient,
	)
	// TODO(jakev): Ensure we validate mandatory headers

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/echo"

	err := req.WriteJSON("POST", fullURL, headers, r)
	if err != nil {
		return defaultRes, nil, err
	}
	res, err := req.Do(ctx)
	if err != nil {
		return defaultRes, nil, err
	}

	respHeaders := map[string]string{}
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200})

	switch res.StatusCode {
	case 200:
		var responseBody string
		err = res.ReadAndUnmarshalNonStructBody(&responseBody)
		if err != nil {
			return defaultRes, respHeaders, err
		}

		return responseBody, respHeaders, nil
	}

	return defaultRes, respHeaders, errors.Errorf(
		"Unexpected http client response (%d)", res.StatusCode,
	)
}
