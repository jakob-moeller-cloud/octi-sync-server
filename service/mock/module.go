// Code generated by MockGen. DO NOT EDIT.
// Source: module.go
//
// Generated by this command:
//
//	mockgen -source module.go -package mock -destination mock/module.go Module
//
// Package mock is a generated GoMock package.
package mock

import (
	io "io"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockModule is a mock of Module interface.
type MockModule struct {
	ctrl     *gomock.Controller
	recorder *MockModuleMockRecorder
}

// MockModuleMockRecorder is the mock recorder for MockModule.
type MockModuleMockRecorder struct {
	mock *MockModule
}

// NewMockModule creates a new mock instance.
func NewMockModule(ctrl *gomock.Controller) *MockModule {
	mock := &MockModule{ctrl: ctrl}
	mock.recorder = &MockModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockModule) EXPECT() *MockModuleMockRecorder {
	return m.recorder
}

// Raw mocks base method.
func (m *MockModule) Raw() io.Reader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Raw")
	ret0, _ := ret[0].(io.Reader)
	return ret0
}

// Raw indicates an expected call of Raw.
func (mr *MockModuleMockRecorder) Raw() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Raw", reflect.TypeOf((*MockModule)(nil).Raw))
}

// Size mocks base method.
func (m *MockModule) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockModuleMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockModule)(nil).Size))
}
