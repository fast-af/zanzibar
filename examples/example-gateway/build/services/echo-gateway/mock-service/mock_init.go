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

package echogatewayservicegeneratedmock

import (
	"github.com/golang/mock/gomock"
	module "github.com/uber/zanzibar/examples/example-gateway/build/services/echo-gateway/module"
	zanzibar "github.com/uber/zanzibar/runtime"

	bazclientgenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/baz/mock-client"
	echoclientgenerated "github.com/uber/zanzibar/examples/example-gateway/build/clients/echo/mock-client"
	bounceendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bounce"
	bounceendpointmodule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/bounce/module"
	echoendpointgenerated "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/echo"
	echoendpointmodule "github.com/uber/zanzibar/examples/example-gateway/build/endpoints/tchannel/echo/module"
	defaultexamplemiddlewaregenerated "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/default/default_example"
	defaultexamplemiddlewaremodule "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/default/default_example/module"
	defaultexample2middlewaregenerated "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/default/default_example2"
	defaultexample2middlewaremodule "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/default/default_example2/module"
	defaultexampletchannelmiddlewaregenerated "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/default/default_example_tchannel"
	defaultexampletchannelmiddlewaremodule "github.com/uber/zanzibar/examples/example-gateway/build/middlewares/default/default_example_tchannel/module"
)

// MockClientNodes contains mock client dependencies
type MockClientNodes struct {
	Baz  *bazclientgenerated.MockClient
	Echo *echoclientgenerated.MockClient
}

// InitializeDependenciesMock fully initializes all dependencies in the dep tree
// for the echo-gateway service with leaf nodes being mocks
func InitializeDependenciesMock(
	g *zanzibar.Gateway,
	ctrl *gomock.Controller,
) (*module.DependenciesTree, *module.Dependencies, *MockClientNodes) {
	tree := &module.DependenciesTree{}

	initializedDefaultDependencies := &zanzibar.DefaultDependencies{
		ContextExtractor:     g.ContextExtractor,
		ContextMetrics:       g.ContextMetrics,
		ContextLogger:        g.ContextLogger,
		Logger:               g.Logger,
		Scope:                g.RootScope,
		Config:               g.Config,
		ServerTChannel:       g.ServerTChannel,
		Tracer:               g.Tracer,
		GRPCClientDispatcher: g.GRPCClientDispatcher,
		JSONWrapper:          g.JSONWrapper,
	}

	mockClientNodes := &MockClientNodes{
		Baz:  bazclientgenerated.NewMockClient(ctrl),
		Echo: echoclientgenerated.NewMockClient(ctrl),
	}
	initializedClientDependencies := &module.ClientDependenciesNodes{}
	tree.Client = initializedClientDependencies
	initializedClientDependencies.Baz = mockClientNodes.Baz
	initializedClientDependencies.Echo = mockClientNodes.Echo

	initializedMiddlewareDependencies := &module.MiddlewareDependenciesNodes{}
	tree.Middleware = initializedMiddlewareDependencies
	initializedMiddlewareDependencies.DefaultExample = defaultexamplemiddlewaregenerated.NewMiddleware(&defaultexamplemiddlewaremodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &defaultexamplemiddlewaremodule.ClientDependencies{
			Baz: initializedClientDependencies.Baz,
		},
	})
	initializedMiddlewareDependencies.DefaultExample2 = defaultexample2middlewaregenerated.NewMiddleware(&defaultexample2middlewaremodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &defaultexample2middlewaremodule.ClientDependencies{
			Baz: initializedClientDependencies.Baz,
		},
	})
	initializedMiddlewareDependencies.DefaultExampleTchannel = defaultexampletchannelmiddlewaregenerated.NewMiddleware(&defaultexampletchannelmiddlewaremodule.Dependencies{
		Default: initializedDefaultDependencies,
	})

	initializedEndpointDependencies := &module.EndpointDependenciesNodes{}
	tree.Endpoint = initializedEndpointDependencies
	initializedEndpointDependencies.Bounce = bounceendpointgenerated.NewEndpoint(&bounceendpointmodule.Dependencies{
		Default: initializedDefaultDependencies,
		Client: &bounceendpointmodule.ClientDependencies{
			Echo: initializedClientDependencies.Echo,
		},
		Middleware: &bounceendpointmodule.MiddlewareDependencies{
			DefaultExample:         initializedMiddlewareDependencies.DefaultExample,
			DefaultExample2:        initializedMiddlewareDependencies.DefaultExample2,
			DefaultExampleTchannel: initializedMiddlewareDependencies.DefaultExampleTchannel,
		},
	})
	initializedEndpointDependencies.Echo = echoendpointgenerated.NewEndpoint(&echoendpointmodule.Dependencies{
		Default: initializedDefaultDependencies,
		Middleware: &echoendpointmodule.MiddlewareDependencies{
			DefaultExample:         initializedMiddlewareDependencies.DefaultExample,
			DefaultExample2:        initializedMiddlewareDependencies.DefaultExample2,
			DefaultExampleTchannel: initializedMiddlewareDependencies.DefaultExampleTchannel,
		},
	})

	dependencies := &module.Dependencies{
		Default: initializedDefaultDependencies,
		Endpoint: &module.EndpointDependencies{
			Bounce: initializedEndpointDependencies.Bounce,
			Echo:   initializedEndpointDependencies.Echo,
		},
	}

	return tree, dependencies, mockClientNodes
}
