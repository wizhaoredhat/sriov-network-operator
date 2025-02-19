// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mock_utils is a generated GoMock package.
package mock_utils

import (
	reflect "reflect"

	v1 "github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockStoreManagerInterface is a mock of StoreManagerInterface interface.
type MockStoreManagerInterface struct {
	ctrl     *gomock.Controller
	recorder *MockStoreManagerInterfaceMockRecorder
}

// MockStoreManagerInterfaceMockRecorder is the mock recorder for MockStoreManagerInterface.
type MockStoreManagerInterfaceMockRecorder struct {
	mock *MockStoreManagerInterface
}

// NewMockStoreManagerInterface creates a new mock instance.
func NewMockStoreManagerInterface(ctrl *gomock.Controller) *MockStoreManagerInterface {
	mock := &MockStoreManagerInterface{ctrl: ctrl}
	mock.recorder = &MockStoreManagerInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStoreManagerInterface) EXPECT() *MockStoreManagerInterfaceMockRecorder {
	return m.recorder
}

// ClearPCIAddressFolder mocks base method.
func (m *MockStoreManagerInterface) ClearPCIAddressFolder() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ClearPCIAddressFolder")
	ret0, _ := ret[0].(error)
	return ret0
}

// ClearPCIAddressFolder indicates an expected call of ClearPCIAddressFolder.
func (mr *MockStoreManagerInterfaceMockRecorder) ClearPCIAddressFolder() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ClearPCIAddressFolder", reflect.TypeOf((*MockStoreManagerInterface)(nil).ClearPCIAddressFolder))
}

// LoadPfsStatus mocks base method.
func (m *MockStoreManagerInterface) LoadPfsStatus(pciAddress string) (*v1.Interface, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoadPfsStatus", pciAddress)
	ret0, _ := ret[0].(*v1.Interface)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LoadPfsStatus indicates an expected call of LoadPfsStatus.
func (mr *MockStoreManagerInterfaceMockRecorder) LoadPfsStatus(pciAddress interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadPfsStatus", reflect.TypeOf((*MockStoreManagerInterface)(nil).LoadPfsStatus), pciAddress)
}

// SaveLastPfAppliedStatus mocks base method.
func (m *MockStoreManagerInterface) SaveLastPfAppliedStatus(PfInfo *v1.Interface) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveLastPfAppliedStatus", PfInfo)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveLastPfAppliedStatus indicates an expected call of SaveLastPfAppliedStatus.
func (mr *MockStoreManagerInterfaceMockRecorder) SaveLastPfAppliedStatus(PfInfo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveLastPfAppliedStatus", reflect.TypeOf((*MockStoreManagerInterface)(nil).SaveLastPfAppliedStatus), PfInfo)
}
