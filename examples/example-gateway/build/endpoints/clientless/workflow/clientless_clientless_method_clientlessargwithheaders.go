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

package clientlessworkflow

import (
	"context"

	zanzibar "github.com/uber/zanzibar/runtime"

	endpointsIDlEndpointsClientlessClientless "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/endpoints-idl/endpoints/clientless/clientless"

	module "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/clientless/module"
	"go.uber.org/zap"
)

// ClientlessClientlessArgWithHeadersWorkflow defines the interface for ClientlessClientlessArgWithHeaders workflow
type ClientlessClientlessArgWithHeadersWorkflow interface {
	Handle(
		ctx context.Context,
		reqHeaders zanzibar.Header,
		r *endpointsIDlEndpointsClientlessClientless.Clientless_ClientlessArgWithHeaders_Args,
	) (context.Context, *endpointsIDlEndpointsClientlessClientless.Response, zanzibar.Header, error)
}

// NewClientlessClientlessArgWithHeadersWorkflow creates a workflow
func NewClientlessClientlessArgWithHeadersWorkflow(deps *module.Dependencies) ClientlessClientlessArgWithHeadersWorkflow {

	return &clientlessClientlessArgWithHeadersWorkflow{
		Logger: deps.Default.Logger,
	}
}

// clientlessClientlessArgWithHeadersWorkflow calls thrift client .
type clientlessClientlessArgWithHeadersWorkflow struct {
	Logger *zap.Logger
}

// Handle processes the request without a downstream
func (w clientlessClientlessArgWithHeadersWorkflow) Handle(
	ctx context.Context,
	reqHeaders zanzibar.Header,
	r *endpointsIDlEndpointsClientlessClientless.Clientless_ClientlessArgWithHeaders_Args,
) (context.Context, *endpointsIDlEndpointsClientlessClientless.Response, zanzibar.Header, error) {
	response := convertClientlessArgWithHeadersDummyResponse(r)

	clientlessHeaders := map[string]string{}

	var ok bool
	var h string
	h, ok = reqHeaders.Get("X-Deputy-Forwarded")
	if ok {
		clientlessHeaders["X-Deputy-Forwarded"] = h
	}
	h, ok = reqHeaders.Get("X-Uuid")
	if ok {
		clientlessHeaders["X-Uuid"] = h
	}

	// Filter and map response headers from client to server response.
	resHeaders := zanzibar.ServerHTTPHeader{}
	h, ok = clientlessHeaders["X-Token"]
	if ok {
		resHeaders.Set("X-Token", h)
	}
	h, ok = clientlessHeaders["X-Uuid"]
	if ok {
		resHeaders.Set("X-Uuid", h)
	}

	resHeaders.Set(zanzibar.ClientTypeKey, "clientless")
	return ctx, response, resHeaders, nil
}

func convertClientlessArgWithHeadersDummyResponse(in *endpointsIDlEndpointsClientlessClientless.Clientless_ClientlessArgWithHeaders_Args) *endpointsIDlEndpointsClientlessClientless.Response {
	out := &endpointsIDlEndpointsClientlessClientless.Response{}

	return out
}
