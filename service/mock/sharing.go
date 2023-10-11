// Code generated by MockGen. DO NOT EDIT.
// Source: sharing.go
//
// Generated by this command:
//
//	mockgen -source sharing.go -package mock -destination mock/sharing.go Sharing
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	service "github.com/jakob-moeller-cloud/octi-sync-server/service"
	gomock "go.uber.org/mock/gomock"
)

// MockSharing is a mock of Sharing interface.
type MockSharing struct {
	ctrl     *gomock.Controller
	recorder *MockSharingMockRecorder
}

// MockSharingMockRecorder is the mock recorder for MockSharing.
type MockSharingMockRecorder struct {
	mock *MockSharing
}

// NewMockSharing creates a new mock instance.
func NewMockSharing(ctrl *gomock.Controller) *MockSharing {
	mock := &MockSharing{ctrl: ctrl}
	mock.recorder = &MockSharingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSharing) EXPECT() *MockSharingMockRecorder {
	return m.recorder
}

// Revoke mocks base method.
func (m *MockSharing) Revoke(ctx context.Context, shareCode service.ShareCode) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revoke", ctx, shareCode)
	ret0, _ := ret[0].(error)
	return ret0
}

// Revoke indicates an expected call of Revoke.
func (mr *MockSharingMockRecorder) Revoke(ctx, shareCode any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revoke", reflect.TypeOf((*MockSharing)(nil).Revoke), ctx, shareCode)
}

// Share mocks base method.
func (m *MockSharing) Share(ctx context.Context, account service.Account) (service.ShareCode, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Share", ctx, account)
	ret0, _ := ret[0].(service.ShareCode)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Share indicates an expected call of Share.
func (mr *MockSharingMockRecorder) Share(ctx, account any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Share", reflect.TypeOf((*MockSharing)(nil).Share), ctx, account)
}

// Shared mocks base method.
func (m *MockSharing) Shared(ctx context.Context, shareCode service.ShareCode) (service.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Shared", ctx, shareCode)
	ret0, _ := ret[0].(service.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Shared indicates an expected call of Shared.
func (mr *MockSharingMockRecorder) Shared(ctx, shareCode any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Shared", reflect.TypeOf((*MockSharing)(nil).Shared), ctx, shareCode)
}
