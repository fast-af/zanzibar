// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/uber/zanzibar/examples/example-gateway/build/clients/withexceptions (interfaces: Client)

// Package clientmock is a generated GoMock package.
package clientmock

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	withexceptions "github.com/uber/zanzibar/examples/example-gateway/build/gen-code/clients-idl/clients/withexceptions/withexceptions"
	runtime "github.com/uber/zanzibar/runtime"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Func1 mocks base method.
func (m *MockClient) Func1(arg0 context.Context, arg1 map[string]string) (context.Context, *withexceptions.Response, map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Func1", arg0, arg1)
	ret0, _ := ret[0].(context.Context)
	ret1, _ := ret[1].(*withexceptions.Response)
	ret2, _ := ret[2].(map[string]string)
	ret3, _ := ret[3].(error)
	return ret0, ret1, ret2, ret3
}

// Func1 indicates an expected call of Func1.
func (mr *MockClientMockRecorder) Func1(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Func1", reflect.TypeOf((*MockClient)(nil).Func1), arg0, arg1)
}

// HTTPClient mocks base method.
func (m *MockClient) HTTPClient() *runtime.HTTPClient {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HTTPClient")
	ret0, _ := ret[0].(*runtime.HTTPClient)
	return ret0
}

// HTTPClient indicates an expected call of HTTPClient.
func (mr *MockClientMockRecorder) HTTPClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HTTPClient", reflect.TypeOf((*MockClient)(nil).HTTPClient))
}
