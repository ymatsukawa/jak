// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/http/client.go
//
// Generated by this command:
//
//	mockgen -source=./internal/http/client.go -destination=./test/mock/internal/http/client.go
//

// Package mock_http is a generated GoMock package.
package mock_http

import (
	reflect "reflect"
	time "time"

	http "github.com/ymatsukawa/jak/internal/http"
	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
	isgomock struct{}
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

// Do mocks base method.
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Do", req)
	ret0, _ := ret[0].(*http.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Do indicates an expected call of Do.
func (mr *MockClientMockRecorder) Do(req any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Do", reflect.TypeOf((*MockClient)(nil).Do), req)
}

// SetTimeout mocks base method.
func (m *MockClient) SetTimeout(timeout time.Duration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetTimeout", timeout)
}

// SetTimeout indicates an expected call of SetTimeout.
func (mr *MockClientMockRecorder) SetTimeout(timeout any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetTimeout", reflect.TypeOf((*MockClient)(nil).SetTimeout), timeout)
}
