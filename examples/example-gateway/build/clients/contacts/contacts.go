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

package contactsclient

import (
	"context"
	"fmt"
	"time"

	"github.com/afex/hystrix-go/hystrix"

	zanzibar "github.com/uber/zanzibar/runtime"
	"github.com/uber/zanzibar/runtime/jsonwrapper"

	module "github.com/uber/zanzibar/examples/example-gateway/build/clients/contacts/module"
	clientsIDlClientsContactsContacts "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients-idl/clients/contacts/contacts"
)

// Client defines contacts client interface.
type Client interface {
	HTTPClient() *zanzibar.HTTPClient
	SaveContacts(
		ctx context.Context,
		reqHeaders map[string]string,
		args *clientsIDlClientsContactsContacts.Contacts_SaveContacts_Args,
	) (*clientsIDlClientsContactsContacts.SaveContactsResponse, map[string]string, error)
	TestURLURL(
		ctx context.Context,
		reqHeaders map[string]string,
	) (string, map[string]string, error)
}

// contactsClient is the http client.
type contactsClient struct {
	clientID                  string
	httpClient                *zanzibar.HTTPClient
	jsonWrapper               jsonwrapper.JSONWrapper
	circuitBreakerDisabled    bool
	requestUUIDHeaderKey      string
	requestProcedureHeaderKey string
}

// NewClient returns a new http client.
func NewClient(deps *module.Dependencies) Client {
	ip := deps.Default.Config.MustGetString("clients.contacts.ip")
	port := deps.Default.Config.MustGetInt("clients.contacts.port")
	baseURL := fmt.Sprintf("http://%s:%d", ip, port)
	timeoutVal := int(deps.Default.Config.MustGetInt("clients.contacts.timeout"))
	timeout := time.Millisecond * time.Duration(
		timeoutVal,
	)
	defaultHeaders := make(map[string]string)
	if deps.Default.Config.ContainsKey("http.defaultHeaders") {
		deps.Default.Config.MustGetStruct("http.defaultHeaders", &defaultHeaders)
	}
	if deps.Default.Config.ContainsKey("clients.contacts.defaultHeaders") {
		deps.Default.Config.MustGetStruct("clients.contacts.defaultHeaders", &defaultHeaders)
	}
	var requestUUIDHeaderKey string
	if deps.Default.Config.ContainsKey("http.clients.requestUUIDHeaderKey") {
		requestUUIDHeaderKey = deps.Default.Config.MustGetString("http.clients.requestUUIDHeaderKey")
	}
	var requestProcedureHeaderKey string
	if deps.Default.Config.ContainsKey("http.clients.requestProcedureHeaderKey") {
		requestProcedureHeaderKey = deps.Default.Config.MustGetString("http.clients.requestProcedureHeaderKey")
	}
	followRedirect := true
	if deps.Default.Config.ContainsKey("clients.contacts.followRedirect") {
		followRedirect = deps.Default.Config.MustGetBoolean("clients.contacts.followRedirect")
	}

	methodNames := map[string]string{
		"SaveContacts": "Contacts::saveContacts",
		"TestURLURL":   "Contacts::testUrlUrl",
	}
	// circuitBreakerDisabled sets whether circuit-breaker should be disabled
	circuitBreakerDisabled := false
	if deps.Default.Config.ContainsKey("clients.contacts.circuitBreakerDisabled") {
		circuitBreakerDisabled = deps.Default.Config.MustGetBoolean("clients.contacts.circuitBreakerDisabled")
	}
	qpsLevels := map[string]string{}
	if !circuitBreakerDisabled {
		for methodName := range methodNames {
			circuitBreakerName := "contacts" + "-" + methodName
			qpsLevel := ""
			if level, ok := qpsLevels[circuitBreakerName]; ok {
				qpsLevel = level
			}
			configureCircuitBreaker(deps, timeoutVal, circuitBreakerName, qpsLevel)
		}
	}

	levels := map[string]string{}

	return &contactsClient{
		clientID: "contacts",
		httpClient: zanzibar.NewHTTPClientContext(
			deps.Default.ContextLogger, deps.Default.ContextMetrics, deps.Default.JSONWrapper,
			"contacts",
			methodNames,
			baseURL,
			defaultHeaders,
			timeout,
			followRedirect,
		),
		circuitBreakerDisabled:    circuitBreakerDisabled,
		requestUUIDHeaderKey:      requestUUIDHeaderKey,
		requestProcedureHeaderKey: requestProcedureHeaderKey,
	}
}

func configureCircuitBreaker(deps *module.Dependencies, timeoutVal int, circuitBreakerName string, qpsLevel string) {
	// sleepWindowInMilliseconds sets the amount of time, after tripping the circuit,
	// to reject requests before allowing attempts again to determine if the circuit should again be closed
	sleepWindowInMilliseconds := 5000
	// maxConcurrentRequests sets how many requests can be run at the same time, beyond which requests are rejected
	maxConcurrentRequests := 20
	// errorPercentThreshold sets the error percentage at or above which the circuit should trip open
	errorPercentThreshold := 20
	// requestVolumeThreshold sets a minimum number of requests that will trip the circuit in a rolling window of 10s
	// For example, if the value is 20, then if only 19 requests are received in the rolling window of 10 seconds
	// the circuit will not trip open even if all 19 failed.
	requestVolumeThreshold := 20
	// first checks if level exists in configurations then assigns parameters
	if deps.Default.Config.ContainsKey(qpsLevel) {
		var params map[string]int = make(map[string]int)
		deps.Default.Config.MustGetStruct(qpsLevel, &params)
		if sleepWindow, ok := params["sleepWindowInMilliseconds"]; ok {
			sleepWindowInMilliseconds = sleepWindow
		}
		if max, ok := params["maxConcurrentRequests"]; ok {
			maxConcurrentRequests = max
		}
		if errorPercent, ok := params["errorPercentThreshold"]; ok {
			errorPercentThreshold = errorPercent
		}
		if requestVolume, ok := params["requestVolumeThreshold"]; ok {
			requestVolumeThreshold = requestVolume
		}
	}
	// client settings override parameters
	if deps.Default.Config.ContainsKey("clients.contacts.sleepWindowInMilliseconds") {
		sleepWindowInMilliseconds = int(deps.Default.Config.MustGetInt("clients.contacts.sleepWindowInMilliseconds"))
	}
	if deps.Default.Config.ContainsKey("clients.contacts.maxConcurrentRequests") {
		maxConcurrentRequests = int(deps.Default.Config.MustGetInt("clients.contacts.maxConcurrentRequests"))
	}
	if deps.Default.Config.ContainsKey("clients.contacts.errorPercentThreshold") {
		errorPercentThreshold = int(deps.Default.Config.MustGetInt("clients.contacts.errorPercentThreshold"))
	}
	if deps.Default.Config.ContainsKey("clients.contacts.requestVolumeThreshold") {
		requestVolumeThreshold = int(deps.Default.Config.MustGetInt("clients.contacts.requestVolumeThreshold"))
	}
	hystrix.ConfigureCommand(circuitBreakerName, hystrix.CommandConfig{
		MaxConcurrentRequests:  maxConcurrentRequests,
		ErrorPercentThreshold:  errorPercentThreshold,
		SleepWindow:            sleepWindowInMilliseconds,
		RequestVolumeThreshold: requestVolumeThreshold,
		Timeout:                timeoutVal,
	})
}

// HTTPClient returns the underlying HTTP client, should only be
// used for internal testing.
func (c *contactsClient) HTTPClient() *zanzibar.HTTPClient {
	return c.httpClient
}

// SaveContacts calls "/:userUUID/contacts" endpoint.
func (c *contactsClient) SaveContacts(
	ctx context.Context,
	headers map[string]string,
	r *clientsIDlClientsContactsContacts.Contacts_SaveContacts_Args,
) (*clientsIDlClientsContactsContacts.SaveContactsResponse, map[string]string, error) {
	reqUUID := zanzibar.RequestUUIDFromCtx(ctx)
	if headers == nil {
		headers = make(map[string]string)
	}
	if reqUUID != "" {
		headers[c.requestUUIDHeaderKey] = reqUUID
	}
	if c.requestProcedureHeaderKey != "" {
		headers[c.requestProcedureHeaderKey] = "Contacts::saveContacts"
	}

	var defaultRes *clientsIDlClientsContactsContacts.SaveContactsResponse
	req := zanzibar.NewClientHTTPRequest(ctx, c.clientID, "SaveContacts", "Contacts::saveContacts", c.httpClient)

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/" + string(r.SaveContactsRequest.UserUUID) + "/contacts"

	err := req.WriteJSON("POST", fullURL, headers, r.SaveContactsRequest)
	if err != nil {
		return defaultRes, nil, err
	}

	var res *zanzibar.ClientHTTPResponse
	if c.circuitBreakerDisabled {
		res, err = req.Do()
	} else {
		// We want hystrix ckt-breaker to count errors only for system issues
		var clientErr error
		circuitBreakerName := "contacts" + "-" + "SaveContacts"
		err = hystrix.DoC(ctx, circuitBreakerName, func(ctx context.Context) error {
			res, clientErr = req.Do()
			if res != nil {
				// This is not a system error/issue. Downstream responded
				return nil
			}
			return clientErr
		}, nil)
		if err == nil {
			// ckt-breaker was ok, bubble up client error if set
			err = clientErr
		}
	}
	if err != nil {
		return defaultRes, nil, err
	}

	respHeaders := make(map[string]string)
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{202, 400, 404})

	switch res.StatusCode {
	case 202:
		var responseBody clientsIDlClientsContactsContacts.SaveContactsResponse
		rawBody, err := res.ReadAll()
		if err != nil {
			return defaultRes, respHeaders, err
		}
		err = res.UnmarshalBody(&responseBody, rawBody)
		if err != nil {
			return defaultRes, respHeaders, err
		}

		return &responseBody, respHeaders, nil

	case 400:
		return defaultRes, respHeaders, &clientsIDlClientsContactsContacts.BadRequest{}
	case 404:
		return defaultRes, respHeaders, &clientsIDlClientsContactsContacts.NotFound{}

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

// TestURLURL calls "/contacts/testUrl" endpoint.
func (c *contactsClient) TestURLURL(
	ctx context.Context,
	headers map[string]string,
) (string, map[string]string, error) {
	reqUUID := zanzibar.RequestUUIDFromCtx(ctx)
	if headers == nil {
		headers = make(map[string]string)
	}
	if reqUUID != "" {
		headers[c.requestUUIDHeaderKey] = reqUUID
	}
	if c.requestProcedureHeaderKey != "" {
		headers[c.requestProcedureHeaderKey] = "Contacts::testUrlUrl"
	}

	var defaultRes string
	req := zanzibar.NewClientHTTPRequest(ctx, c.clientID, "TestURLURL", "Contacts::testUrlUrl", c.httpClient)

	// Generate full URL.
	fullURL := c.httpClient.BaseURL + "/contacts" + "/testUrl"

	err := req.WriteJSON("GET", fullURL, headers, nil)
	if err != nil {
		return defaultRes, nil, err
	}

	var res *zanzibar.ClientHTTPResponse
	if c.circuitBreakerDisabled {
		res, err = req.Do()
	} else {
		// We want hystrix ckt-breaker to count errors only for system issues
		var clientErr error
		circuitBreakerName := "contacts" + "-" + "TestURLURL"
		err = hystrix.DoC(ctx, circuitBreakerName, func(ctx context.Context) error {
			res, clientErr = req.Do()
			if res != nil {
				// This is not a system error/issue. Downstream responded
				return nil
			}
			return clientErr
		}, nil)
		if err == nil {
			// ckt-breaker was ok, bubble up client error if set
			err = clientErr
		}
	}
	if err != nil {
		return defaultRes, nil, err
	}

	respHeaders := make(map[string]string)
	for k := range res.Header {
		respHeaders[k] = res.Header.Get(k)
	}

	res.CheckOKResponse([]int{200})

	switch res.StatusCode {
	case 200:
		var responseBody string
		rawBody, err := res.ReadAll()
		if err != nil {
			return defaultRes, respHeaders, err
		}
		err = res.UnmarshalBody(&responseBody, rawBody)
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
