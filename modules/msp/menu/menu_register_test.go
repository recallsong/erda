// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/erda-project/erda-infra/pkg/transport (interfaces: Register)

// Package exporter is a generated GoMock package.
package menu

import (
	reflect "reflect"

	http "github.com/erda-project/erda-infra/pkg/transport/http"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockRegister is a mock of Register interface.
type MockRegister struct {
	ctrl     *gomock.Controller
	recorder *MockRegisterMockRecorder
}

// MockRegisterMockRecorder is the mock recorder for MockRegister.
type MockRegisterMockRecorder struct {
	mock *MockRegister
}

// NewMockRegister creates a new mock instance.
func NewMockRegister(ctrl *gomock.Controller) *MockRegister {
	mock := &MockRegister{ctrl: ctrl}
	mock.recorder = &MockRegisterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRegister) EXPECT() *MockRegisterMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockRegister) Add(arg0, arg1 string, arg2 http.HandlerFunc) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0, arg1, arg2)
}

// Add indicates an expected call of Add.
func (mr *MockRegisterMockRecorder) Add(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRegister)(nil).Add), arg0, arg1, arg2)
}

// RegisterService mocks base method.
func (m *MockRegister) RegisterService(arg0 *grpc.ServiceDesc, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterService", arg0, arg1)
}

// RegisterService indicates an expected call of RegisterService.
func (mr *MockRegisterMockRecorder) RegisterService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterService", reflect.TypeOf((*MockRegister)(nil).RegisterService), arg0, arg1)
}
