// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/openshift/managed-upgrade-operator/pkg/clusterversion (interfaces: ClusterVersion)

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	v1 "github.com/openshift/api/config/v1"
	v1alpha1 "github.com/openshift/managed-upgrade-operator/pkg/apis/upgrade/v1alpha1"
	clusterversion "github.com/openshift/managed-upgrade-operator/pkg/clusterversion"
	reflect "reflect"
)

// MockClusterVersion is a mock of ClusterVersion interface
type MockClusterVersion struct {
	ctrl     *gomock.Controller
	recorder *MockClusterVersionMockRecorder
}

// MockClusterVersionMockRecorder is the mock recorder for MockClusterVersion
type MockClusterVersionMockRecorder struct {
	mock *MockClusterVersion
}

// NewMockClusterVersion creates a new mock instance
func NewMockClusterVersion(ctrl *gomock.Controller) *MockClusterVersion {
	mock := &MockClusterVersion{ctrl: ctrl}
	mock.recorder = &MockClusterVersionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterVersion) EXPECT() *MockClusterVersionMockRecorder {
	return m.recorder
}

// EnsureDesiredConfig mocks base method
func (m *MockClusterVersion) EnsureDesiredConfig(arg0 *v1alpha1.UpgradeConfig) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureDesiredConfig", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EnsureDesiredConfig indicates an expected call of EnsureDesiredConfig
func (mr *MockClusterVersionMockRecorder) EnsureDesiredConfig(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureDesiredConfig", reflect.TypeOf((*MockClusterVersion)(nil).EnsureDesiredConfig), arg0)
}

// GetClusterVersion mocks base method
func (m *MockClusterVersion) GetClusterVersion() (*v1.ClusterVersion, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClusterVersion")
	ret0, _ := ret[0].(*v1.ClusterVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClusterVersion indicates an expected call of GetClusterVersion
func (mr *MockClusterVersionMockRecorder) GetClusterVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClusterVersion", reflect.TypeOf((*MockClusterVersion)(nil).GetClusterVersion))
}

// HasDegradedOperators mocks base method
func (m *MockClusterVersion) HasDegradedOperators() (*clusterversion.HasDegradedOperatorsResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasDegradedOperators")
	ret0, _ := ret[0].(*clusterversion.HasDegradedOperatorsResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasDegradedOperators indicates an expected call of HasDegradedOperators
func (mr *MockClusterVersionMockRecorder) HasDegradedOperators() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasDegradedOperators", reflect.TypeOf((*MockClusterVersion)(nil).HasDegradedOperators))
}

// HasUpgradeCommenced mocks base method
func (m *MockClusterVersion) HasUpgradeCommenced(arg0 *v1alpha1.UpgradeConfig) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUpgradeCommenced", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// HasUpgradeCommenced indicates an expected call of HasUpgradeCommenced
func (mr *MockClusterVersionMockRecorder) HasUpgradeCommenced(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUpgradeCommenced", reflect.TypeOf((*MockClusterVersion)(nil).HasUpgradeCommenced), arg0)
}

// HasUpgradeCompleted mocks base method
func (m *MockClusterVersion) HasUpgradeCompleted(arg0 *v1.ClusterVersion, arg1 *v1alpha1.UpgradeConfig) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasUpgradeCompleted", arg0, arg1)
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasUpgradeCompleted indicates an expected call of HasUpgradeCompleted
func (mr *MockClusterVersionMockRecorder) HasUpgradeCompleted(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasUpgradeCompleted", reflect.TypeOf((*MockClusterVersion)(nil).HasUpgradeCompleted), arg0, arg1)
}
