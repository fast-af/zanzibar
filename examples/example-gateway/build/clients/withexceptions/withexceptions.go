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

package withexceptionsclient

import (
	"context"
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix"

	zanzibar "github.com/uber/zanzibar/runtime"

	module "github.com/uber/zanzibar/examples/example-gateway/build/clients/withexceptions/module"
	clientsWithexceptionsWithexceptions "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients/withexceptions/withexceptions"
)

// Client defines withexceptions client interface.
type Client interface {
	HTTPClient() *zanzibar.HTTPClient
	Func1(
		ctx context.Context,
		reqHeaders map[string]string,
	) (*clientsWithexceptionsWithexceptions.Response, map[string]string, error)
}

// withexceptionsClient is the http client.
type withexceptionsClient struct {
	clientID               string
	httpClient             *zanzibar.HTTPClient
	circuitBreakerDisabled bool
	requestUUIDHeaderKey   string
}

// NewClient returns a new http client.
func NewClient(deps *module.Dependencies) Client {
	ip := deps.Default.Config.MustGetString("clients.withexceptions.ip")
	port := deps.Default.Config.MustGetInt("clients.withexceptions.port")
	baseURL := fmt.Sprintf("http://%s:%d", ip, port)
	timeoutVal := int(deps.Default.Config.MustGetInt("clients.withexceptions.timeout"))
	timeout := time.Millisecond * time.Duration(
		timeoutVal,
	)
	defaultHeaders := make(map[string]string)
	if deps.Default.Config.ContainsKey("clients.withexceptions.defaultHeaders") {
		deps.Default.Config.MustGetStruct("clients.withexceptions.defaultHeaders", &defaultHeaders)
	}
	var requestUUIDHeaderKey string
	if deps.Default.Config.ContainsKey("http.clients.requestUUIDHeaderKey") {
		requestUUIDHeaderKey = deps.Default.Config.MustGetString("http.clients.requestUUIDHeaderKey")
	}

	circuitBreakerDisabled := configureCicruitBreaker(deps, timeoutVal)

	return &withexceptionsClient{
		clientID: "withexceptions",
		httpClient: zanzibar.NewHTTPClientContext(
			deps.Default.Logger, deps.Default.ContextMetrics,
			"withexceptions",
			map[string]string{
				"Func1": "WithExceptions::Func1",
			},
			baseURL,
			defaultHeaders,
			timeout,
		),
		circuitBreakerDisabled: circuitBreakerDisabled,
		requestUUIDHeaderKey:   requestUUIDHeaderKey,
	}
}

func configureCicruitBreaker(deps *module.Dependencies, timeoutVal int) bool {
	// circuitBreakerDisabled sets whether circuit-breaker should be disabled
	circuitBreakerDisabled := false
	if deps.Default.Config.ContainsKey("clients.withexceptions.circuitBreakerDisabled") {
		circuitBreakerDisabled = deps.Default.Config.MustGetBoolean("clients.withexceptions.circuitBreakerDisabled")
	}
	// sleepWindowInMilliseconds sets the amount of time, after tripping the circuit,
	// to reject requests before allowing attempts again to determine if the circuit should again be closed
	sleepWindowInMilliseconds := 5000
	if deps.Default.Config.ContainsKey("clients.withexceptions.sleepWindowInMilliseconds") {
		sleepWindowInMilliseconds = int(deps.Default.Config.MustGetInt("clients.withexceptions.sleepWindowInMilliseconds"))
	}
	// maxConcurrentRequests sets how many requests can be run at the same time, beyond which requests are rejected
	maxConcurrentRequests := 20
	if deps.Default.Config.ContainsKey("clients.withexceptions.maxConcurrentRequests") {
		maxConcurrentRequests = int(deps.Default.Config.MustGetInt("clients.withexceptions.maxConcurrentRequests"))
	}
	// errorPercentThreshold sets the error percentage at or above which the circuit should trip open
	errorPercentThreshold := 20
	if deps.Default.Config.ContainsKey("clients.withexceptions.errorPercentThreshold") {
		errorPercentThreshold = int(deps.Default.Config.MustGetInt("clients.withexceptions.errorPercentThreshold"))
	}
	// requestVolumeThreshold sets a minimum number of requests that will trip the circuit in a rolling window of 10s
	// For example, if the value is 20, then if only 19 requests are received in the rolling window of 10 seconds
	// the circuit will not trip open even if all 19 failed.
	requestVolumeThreshold := 20
	if deps.Default.Config.ContainsKey("clients.withexceptions.requestVolumeThreshold") {
		requestVolumeThreshold = int(deps.Default.Config.MustGetInt("clients.withexceptions.requestVolumeThreshold"))
	}
	if !circuitBreakerDisabled {
		hystrix.ConfigureCommand("withexceptions", hystrix.CommandConfig{
			MaxConcurrentRequests:  maxConcurrentRequests,
			ErrorPercentThreshold:  errorPercentThreshold,
			SleepWindow:            sleepWindowInMilliseconds,
			RequestVolumeThreshold: requestVolumeThreshold,
			Timeout:                timeoutVal,
		})
	}
	return circuitBreakerDisabled
}

// HTTPClient returns the underlying HTTP client, should only be
// used for internal testing.
func (c *withexceptionsClient) HTTPClient() *zanzibar.HTTPClient {
	return c.httpClient
}

// Func1 calls "/withexceptions/func1" endpoint.
func (c *withexceptionsClient) Func1(
	ctx context.Context,
	headers map[string]string,
) (*clientsWithexceptionsWithexceptions.Response, map[string]string, error) {
	reqUUID := zanzibar.RequestUUIDFromCtx(ctx)
	if reqUUID != "" {
		if headers == nil {
			headers = make(map[string]string)
		}
		headers[c.requestUUIDHeaderKey] = reqUUID
	}

	var defaultRes *clientsWithexceptionsWithexceptions.Response
	req := zanzibar.NewClientHTTPRequest(ctx, c.clientID, "Func1", "WithExceptions::Func1", c.httpClient)

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/withexceptions" + "/func1"

	err := req.WriteJSON("GET", fullURL, headers, nil)
	if err != nil {
		return defaultRes, nil, err
	}

	var res *zanzibar.ClientHTTPResponse
	if c.circuitBreakerDisabled {
		res, err = req.Do()
	} else {
		err = hystrix.DoC(ctx, "withexceptions", func(ctx context.Context) error {
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

	res.CheckOKResponse([]int{200, 401})

	switch res.StatusCode {
	case 200:
		var responseBody clientsWithexceptionsWithexceptions.Response
		err = res.ReadAndUnmarshalBody(&responseBody)
		if err != nil {
			return defaultRes, respHeaders, err
		}
		return &responseBody, respHeaders, nil

	case 401:
		allOptions := []interface{}{
			&clientsWithexceptionsWithexceptions.ExceptionType1{}, &clientsWithexceptionsWithexceptions.ExceptionType2{},
		}
		v, err := res.ReadAndUnmarshalBodyMultipleOptions(allOptions)
		if err != nil {
			return defaultRes, respHeaders, err
		}
		return defaultRes, respHeaders, v.(error)

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
