// Code generated by MockGen. DO NOT EDIT.
// Source: store.go
//
// Generated by this command:
//
//	mockgen -destination mock/mock_store.go -source store.go
//

// Package mock_store is a generated GoMock package.
package mock_store

import (
	reflect "reflect"

	v1 "github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
	isgomock struct{}
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// AddManagedOVSBridge mocks base method.
func (m *MockStore) AddManagedOVSBridge(br *v1.OVSConfigExt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddManagedOVSBridge", br)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddManagedOVSBridge indicates an expected call of AddManagedOVSBridge.
func (mr *MockStoreMockRecorder) AddManagedOVSBridge(br any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddManagedOVSBridge", reflect.TypeOf((*MockStore)(nil).AddManagedOVSBridge), br)
}

// GetManagedOVSBridge mocks base method.
func (m *MockStore) GetManagedOVSBridge(name string) (*v1.OVSConfigExt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedOVSBridge", name)
	ret0, _ := ret[0].(*v1.OVSConfigExt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManagedOVSBridge indicates an expected call of GetManagedOVSBridge.
func (mr *MockStoreMockRecorder) GetManagedOVSBridge(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedOVSBridge", reflect.TypeOf((*MockStore)(nil).GetManagedOVSBridge), name)
}

// GetManagedOVSBridges mocks base method.
func (m *MockStore) GetManagedOVSBridges() (map[string]*v1.OVSConfigExt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetManagedOVSBridges")
	ret0, _ := ret[0].(map[string]*v1.OVSConfigExt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetManagedOVSBridges indicates an expected call of GetManagedOVSBridges.
func (mr *MockStoreMockRecorder) GetManagedOVSBridges() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetManagedOVSBridges", reflect.TypeOf((*MockStore)(nil).GetManagedOVSBridges))
}

// RemoveManagedOVSBridge mocks base method.
func (m *MockStore) RemoveManagedOVSBridge(name string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveManagedOVSBridge", name)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveManagedOVSBridge indicates an expected call of RemoveManagedOVSBridge.
func (mr *MockStoreMockRecorder) RemoveManagedOVSBridge(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveManagedOVSBridge", reflect.TypeOf((*MockStore)(nil).RemoveManagedOVSBridge), name)
}
