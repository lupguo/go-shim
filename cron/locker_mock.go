// Code generated by MockGen. DO NOT EDIT.
// Source: ./Locker.go

// Package cron is a generated GoMock package.
package cron

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockLocker is a mock of DistributeLocker interface
type MockLocker struct {
	ctrl     *gomock.Controller
	recorder *MockLockerMockRecorder
}

// MockLockerMockRecorder is the mock recorder for MockLocker
type MockLockerMockRecorder struct {
	mock *MockLocker
}

// NewMockLocker creates a new mock instance
func NewMockLocker(ctrl *gomock.Controller) *MockLocker {
	mock := &MockLocker{ctrl: ctrl}
	mock.recorder = &MockLockerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockLocker) EXPECT() *MockLockerMockRecorder {
	return m.recorder
}

// Lock mocks base method
func (m *MockLocker) Lock(key string, ttl time.Duration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Lock", key, ttl)
	ret0, _ := ret[0].(error)
	return ret0
}

// Lock indicates an expected call of Lock
func (mr *MockLockerMockRecorder) Lock(key, ttl interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Lock", reflect.TypeOf((*MockLocker)(nil).Lock), key, ttl)
}

// UnLock mocks base method
func (m *MockLocker) UnLock(key string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UnLock", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnLock indicates an expected call of UnLock
func (mr *MockLockerMockRecorder) UnLock(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnLock", reflect.TypeOf((*MockLocker)(nil).UnLock), key)
}