// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
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

package multiclient

import (
	"context"
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix"

	zanzibar "github.com/uber/zanzibar/runtime"

	module "github.com/uber/zanzibar/examples/example-gateway/build/clients/multi/module"
)

// Client defines multi client interface.
type Client interface {
	HTTPClient() *zanzibar.HTTPClient
	HelloA(
		ctx context.Context,
		reqHeaders map[string]string,
	) (string, map[string]string, error)
	HelloB(
		ctx context.Context,
		reqHeaders map[string]string,
	) (string, map[string]string, error)
}

// multiClient is the http client.
type multiClient struct {
	clientID               string
	httpClient             *zanzibar.HTTPClient
	circuitBreakerDisabled bool
}

// NewClient returns a new http client.
func NewClient(deps *module.Dependencies) Client {
	ip := deps.Default.Config.MustGetString("clients.multi.ip")
	port := deps.Default.Config.MustGetInt("clients.multi.port")
	baseURL := fmt.Sprintf("http://%s:%d", ip, port)
	timeout := time.Duration(deps.Default.Config.MustGetInt("clients.multi.timeout")) * time.Millisecond
	defaultHeaders := make(map[string]string)
	if deps.Default.Config.ContainsKey("clients.multi.defaultHeaders") {
		deps.Default.Config.MustGetStruct("clients.multi.defaultHeaders", &defaultHeaders)
	}

	maxConcurrentRequests := deps.Default.Config.MustGetInt("clients.multi.maxConcurrentRequests")
	errorPercentThreshold := deps.Default.Config.MustGetInt("clients.multi.errorPercentThreshold")
	hystrix.ConfigureCommand("multi", hystrix.CommandConfig{
		MaxConcurrentRequests: int(maxConcurrentRequests),
		ErrorPercentThreshold: int(errorPercentThreshold),
	})

	circuitBreakerDisabled := deps.Default.Config.ContainsKey("clients.multi.circuitBreakerDisabled") &&
		deps.Default.Config.MustGetBoolean("clients.multi.circuitBreakerDisabled")

	return &multiClient{
		clientID: "multi",
		httpClient: zanzibar.NewHTTPClientContext(
			deps.Default.Logger, deps.Default.ContextMetrics,
			"multi",
			[]string{
				"HelloA",
				"HelloB",
			},
			baseURL,
			defaultHeaders,
			timeout,
		),
		circuitBreakerDisabled: circuitBreakerDisabled,
	}
}

// HTTPClient returns the underlying HTTP client, should only be
// used for internal testing.
func (c *multiClient) HTTPClient() *zanzibar.HTTPClient {
	return c.httpClient
}

// HelloA calls "/multi/serviceA_b/hello" endpoint.
func (c *multiClient) HelloA(
	ctx context.Context,
	headers map[string]string,
) (string, map[string]string, error) {
	var defaultRes string
	req := zanzibar.NewClientHTTPRequest(ctx, c.clientID, "HelloA", c.httpClient)

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/multi" + "/serviceA_b" + "/hello"

	err := req.WriteJSON("GET", fullURL, headers, nil)
	if err != nil {
		return defaultRes, nil, err
	}

	var res *zanzibar.ClientHTTPResponse
	if c.circuitBreakerDisabled {
		res, err = req.Do()
	} else {
		err = hystrix.DoC(ctx, "multi", func(ctx context.Context) error {
			res, err = req.Do()
			return err
		}, nil)
	}
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
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return defaultRes, respHeaders, err
		}

		return responseBody, respHeaders, nil
	default:
		_, err = res.ReadAll()
		if err != nil {
			return defaultRes, respHeaders, err
		}
	}

	return defaultRes, respHeaders, &zanzibar.UnexpectedHTTPError{
		StatusCode: res.StatusCode,
		RawBody:    res.GetRawBody(),
	}
}

// HelloB calls "/multi/serviceB_b/hello" endpoint.
func (c *multiClient) HelloB(
	ctx context.Context,
	headers map[string]string,
) (string, map[string]string, error) {
	var defaultRes string
	req := zanzibar.NewClientHTTPRequest(ctx, c.clientID, "HelloB", c.httpClient)

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/multi" + "/serviceB_b" + "/hello"

	err := req.WriteJSON("GET", fullURL, headers, nil)
	if err != nil {
		return defaultRes, nil, err
	}

	var res *zanzibar.ClientHTTPResponse
	if c.circuitBreakerDisabled {
		res, err = req.Do()
	} else {
		err = hystrix.DoC(ctx, "multi", func(ctx context.Context) error {
			res, err = req.Do()
			return err
		}, nil)
	}
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
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return defaultRes, respHeaders, err
		}

		return responseBody, respHeaders, nil
	default:
		_, err = res.ReadAll()
		if err != nil {
			return defaultRes, respHeaders, err
		}
	}

	return defaultRes, respHeaders, &zanzibar.UnexpectedHTTPError{
		StatusCode: res.StatusCode,
		RawBody:    res.GetRawBody(),
	}
}
