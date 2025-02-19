// Code generated by MockGen. DO NOT EDIT.
// Source: ovs.go
//
// Generated by this command:
//
//	mockgen -destination mock/mock_ovs.go -source ovs.go
//

// Package mock_ovs is a generated GoMock package.
package mock_ovs

import (
	context "context"
	reflect "reflect"

	v1 "github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1"
	gomock "go.uber.org/mock/gomock"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
	isgomock struct{}
}

// MockInterfaceMockRecorder is the mock recorder for MockInterface.
type MockInterfaceMockRecorder struct {
	mock *MockInterface
}

// NewMockInterface creates a new mock instance.
func NewMockInterface(ctrl *gomock.Controller) *MockInterface {
	mock := &MockInterface{ctrl: ctrl}
	mock.recorder = &MockInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInterface) EXPECT() *MockInterfaceMockRecorder {
	return m.recorder
}

// CreateOVSBridge mocks base method.
func (m *MockInterface) CreateOVSBridge(ctx context.Context, conf *v1.OVSConfigExt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOVSBridge", ctx, conf)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateOVSBridge indicates an expected call of CreateOVSBridge.
func (mr *MockInterfaceMockRecorder) CreateOVSBridge(ctx, conf any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOVSBridge", reflect.TypeOf((*MockInterface)(nil).CreateOVSBridge), ctx, conf)
}

// GetOVSBridges mocks base method.
func (m *MockInterface) GetOVSBridges(ctx context.Context) ([]v1.OVSConfigExt, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOVSBridges", ctx)
	ret0, _ := ret[0].([]v1.OVSConfigExt)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetOVSBridges indicates an expected call of GetOVSBridges.
func (mr *MockInterfaceMockRecorder) GetOVSBridges(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOVSBridges", reflect.TypeOf((*MockInterface)(nil).GetOVSBridges), ctx)
}

// RemoveInterfaceFromOVSBridge mocks base method.
func (m *MockInterface) RemoveInterfaceFromOVSBridge(ctx context.Context, ifaceAddr string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveInterfaceFromOVSBridge", ctx, ifaceAddr)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveInterfaceFromOVSBridge indicates an expected call of RemoveInterfaceFromOVSBridge.
func (mr *MockInterfaceMockRecorder) RemoveInterfaceFromOVSBridge(ctx, ifaceAddr any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveInterfaceFromOVSBridge", reflect.TypeOf((*MockInterface)(nil).RemoveInterfaceFromOVSBridge), ctx, ifaceAddr)
}

// RemoveOVSBridge mocks base method.
func (m *MockInterface) RemoveOVSBridge(ctx context.Context, bridgeName string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveOVSBridge", ctx, bridgeName)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveOVSBridge indicates an expected call of RemoveOVSBridge.
func (mr *MockInterfaceMockRecorder) RemoveOVSBridge(ctx, bridgeName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveOVSBridge", reflect.TypeOf((*MockInterface)(nil).RemoveOVSBridge), ctx, bridgeName)
}
