// Code generated by MockGen. DO NOT EDIT.
// Source: username.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockUsernameGenerator is a mock of UsernameGenerator interface.
type MockUsernameGenerator struct {
	ctrl     *gomock.Controller
	recorder *MockUsernameGeneratorMockRecorder
}

// MockUsernameGeneratorMockRecorder is the mock recorder for MockUsernameGenerator.
type MockUsernameGeneratorMockRecorder struct {
	mock *MockUsernameGenerator
}

// NewMockUsernameGenerator creates a new mock instance.
func NewMockUsernameGenerator(ctrl *gomock.Controller) *MockUsernameGenerator {
	mock := &MockUsernameGenerator{ctrl: ctrl}
	mock.recorder = &MockUsernameGeneratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsernameGenerator) EXPECT() *MockUsernameGeneratorMockRecorder {
	return m.recorder
}

// Generate mocks base method.
func (m *MockUsernameGenerator) Generate() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Generate indicates an expected call of Generate.
func (mr *MockUsernameGeneratorMockRecorder) Generate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockUsernameGenerator)(nil).Generate))
}
