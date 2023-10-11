// Code generated by MockGen. DO NOT EDIT.
// Source: devices.go
//
// Generated by this command:
//
//	mockgen -source devices.go -package mock -destination mock/devices.go Devices
//
// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	service "github.com/jakob-moeller-cloud/octi-sync-server/service"
	gomock "go.uber.org/mock/gomock"
)

// MockDevices is a mock of Devices interface.
type MockDevices struct {
	ctrl     *gomock.Controller
	recorder *MockDevicesMockRecorder
}

// MockDevicesMockRecorder is the mock recorder for MockDevices.
type MockDevicesMockRecorder struct {
	mock *MockDevices
}

// NewMockDevices creates a new mock instance.
func NewMockDevices(ctrl *gomock.Controller) *MockDevices {
	mock := &MockDevices{ctrl: ctrl}
	mock.recorder = &MockDevicesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDevices) EXPECT() *MockDevicesMockRecorder {
	return m.recorder
}

// AddDevice mocks base method.
func (m *MockDevices) AddDevice(ctx context.Context, account service.Account, id service.DeviceID, password string) (service.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddDevice", ctx, account, id, password)
	ret0, _ := ret[0].(service.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddDevice indicates an expected call of AddDevice.
func (mr *MockDevicesMockRecorder) AddDevice(ctx, account, id, password any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddDevice", reflect.TypeOf((*MockDevices)(nil).AddDevice), ctx, account, id, password)
}

// DeleteDevice mocks base method.
func (m *MockDevices) DeleteDevice(ctx context.Context, account service.Account, id service.DeviceID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteDevice", ctx, account, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteDevice indicates an expected call of DeleteDevice.
func (mr *MockDevicesMockRecorder) DeleteDevice(ctx, account, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteDevice", reflect.TypeOf((*MockDevices)(nil).DeleteDevice), ctx, account, id)
}

// GetDevice mocks base method.
func (m *MockDevices) GetDevice(ctx context.Context, account service.Account, id service.DeviceID) (service.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevice", ctx, account, id)
	ret0, _ := ret[0].(service.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevice indicates an expected call of GetDevice.
func (mr *MockDevicesMockRecorder) GetDevice(ctx, account, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevice", reflect.TypeOf((*MockDevices)(nil).GetDevice), ctx, account, id)
}

// GetDevices mocks base method.
func (m *MockDevices) GetDevices(ctx context.Context, account service.Account) (map[service.DeviceID]service.Device, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDevices", ctx, account)
	ret0, _ := ret[0].(map[service.DeviceID]service.Device)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDevices indicates an expected call of GetDevices.
func (mr *MockDevicesMockRecorder) GetDevices(ctx, account any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDevices", reflect.TypeOf((*MockDevices)(nil).GetDevices), ctx, account)
}

// HealthCheck mocks base method.
func (m *MockDevices) HealthCheck() service.HealthCheck {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck")
	ret0, _ := ret[0].(service.HealthCheck)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockDevicesMockRecorder) HealthCheck() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockDevices)(nil).HealthCheck))
}
