// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/gardener/pkg/component/etcd/etcd (interfaces: Interface)
//
// Generated by this command:
//
//	mockgen -package mock -destination=mocks.go github.com/gardener/gardener/pkg/component/etcd/etcd Interface
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	v1alpha1 "github.com/gardener/etcd-druid/api/v1alpha1"
	etcd "github.com/gardener/gardener/pkg/component/etcd/etcd"
	gomock "go.uber.org/mock/gomock"
	rest "k8s.io/client-go/rest"
)

// MockInterface is a mock of Interface interface.
type MockInterface struct {
	ctrl     *gomock.Controller
	recorder *MockInterfaceMockRecorder
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

// Deploy mocks base method.
func (m *MockInterface) Deploy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deploy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Deploy indicates an expected call of Deploy.
func (mr *MockInterfaceMockRecorder) Deploy(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deploy", reflect.TypeOf((*MockInterface)(nil).Deploy), arg0)
}

// Destroy mocks base method.
func (m *MockInterface) Destroy(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Destroy", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Destroy indicates an expected call of Destroy.
func (mr *MockInterfaceMockRecorder) Destroy(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Destroy", reflect.TypeOf((*MockInterface)(nil).Destroy), arg0)
}

// Get mocks base method.
func (m *MockInterface) Get(arg0 context.Context) (*v1alpha1.Etcd, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(*v1alpha1.Etcd)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockInterfaceMockRecorder) Get(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockInterface)(nil).Get), arg0)
}

// GetReplicas mocks base method.
func (m *MockInterface) GetReplicas() *int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetReplicas")
	ret0, _ := ret[0].(*int32)
	return ret0
}

// GetReplicas indicates an expected call of GetReplicas.
func (mr *MockInterfaceMockRecorder) GetReplicas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetReplicas", reflect.TypeOf((*MockInterface)(nil).GetReplicas))
}

// GetValues mocks base method.
func (m *MockInterface) GetValues() etcd.Values {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValues")
	ret0, _ := ret[0].(etcd.Values)
	return ret0
}

// GetValues indicates an expected call of GetValues.
func (mr *MockInterfaceMockRecorder) GetValues() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValues", reflect.TypeOf((*MockInterface)(nil).GetValues))
}

// RolloutPeerCA mocks base method.
func (m *MockInterface) RolloutPeerCA(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RolloutPeerCA", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RolloutPeerCA indicates an expected call of RolloutPeerCA.
func (mr *MockInterfaceMockRecorder) RolloutPeerCA(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RolloutPeerCA", reflect.TypeOf((*MockInterface)(nil).RolloutPeerCA), arg0)
}

// Scale mocks base method.
func (m *MockInterface) Scale(arg0 context.Context, arg1 int32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scale", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scale indicates an expected call of Scale.
func (mr *MockInterfaceMockRecorder) Scale(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scale", reflect.TypeOf((*MockInterface)(nil).Scale), arg0, arg1)
}

// SetBackupConfig mocks base method.
func (m *MockInterface) SetBackupConfig(arg0 *etcd.BackupConfig) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetBackupConfig", arg0)
}

// SetBackupConfig indicates an expected call of SetBackupConfig.
func (mr *MockInterfaceMockRecorder) SetBackupConfig(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetBackupConfig", reflect.TypeOf((*MockInterface)(nil).SetBackupConfig), arg0)
}

// SetClientServiceDNSNames mocks base method.
func (m *MockInterface) SetClientServiceDNSNames(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetClientServiceDNSNames", arg0)
}

// SetClientServiceDNSNames indicates an expected call of SetClientServiceDNSNames.
func (mr *MockInterfaceMockRecorder) SetClientServiceDNSNames(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetClientServiceDNSNames", reflect.TypeOf((*MockInterface)(nil).SetClientServiceDNSNames), arg0)
}

// SetPeerServiceDNSNames mocks base method.
func (m *MockInterface) SetPeerServiceDNSNames(arg0 []string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetPeerServiceDNSNames", arg0)
}

// SetPeerServiceDNSNames indicates an expected call of SetPeerServiceDNSNames.
func (mr *MockInterfaceMockRecorder) SetPeerServiceDNSNames(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPeerServiceDNSNames", reflect.TypeOf((*MockInterface)(nil).SetPeerServiceDNSNames), arg0)
}

// SetReplicas mocks base method.
func (m *MockInterface) SetReplicas(arg0 *int32) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetReplicas", arg0)
}

// SetReplicas indicates an expected call of SetReplicas.
func (mr *MockInterfaceMockRecorder) SetReplicas(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetReplicas", reflect.TypeOf((*MockInterface)(nil).SetReplicas), arg0)
}

// SetServiceEndpoint mocks base method.
func (m *MockInterface) SetServiceEndpoint(arg0 *string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetServiceEndpoint", arg0)
}

// SetServiceEndpoint indicates an expected call of SetServiceEndpoint.
func (mr *MockInterfaceMockRecorder) SetServiceEndpoint(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetServiceEndpoint", reflect.TypeOf((*MockInterface)(nil).SetServiceEndpoint), arg0)
}

// SetSkipClientSANVerify mocks base method.
func (m *MockInterface) SetSkipClientSANVerify(arg0 *bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetSkipClientSANVerify", arg0)
}

// SetSkipClientSANVerify indicates an expected call of SetSkipClientSANVerify.
func (mr *MockInterfaceMockRecorder) SetSkipClientSANVerify(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSkipClientSANVerify", reflect.TypeOf((*MockInterface)(nil).SetSkipClientSANVerify), arg0)
}

// SetURLs mocks base method.
func (m *MockInterface) SetURLs(arg0, arg1, arg2 []v1alpha1.MemberInfo) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "SetURLs", arg0, arg1, arg2)
}

// SetURLs indicates an expected call of SetURLs.
func (mr *MockInterfaceMockRecorder) SetURLs(arg0, arg1, arg2 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetURLs", reflect.TypeOf((*MockInterface)(nil).SetURLs), arg0, arg1, arg2)
}

// Snapshot mocks base method.
func (m *MockInterface) Snapshot(arg0 context.Context, arg1 rest.HTTPClient) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Snapshot", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Snapshot indicates an expected call of Snapshot.
func (mr *MockInterfaceMockRecorder) Snapshot(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockInterface)(nil).Snapshot), arg0, arg1)
}

// Wait mocks base method.
func (m *MockInterface) Wait(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Wait", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Wait indicates an expected call of Wait.
func (mr *MockInterfaceMockRecorder) Wait(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Wait", reflect.TypeOf((*MockInterface)(nil).Wait), arg0)
}

// WaitCleanup mocks base method.
func (m *MockInterface) WaitCleanup(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WaitCleanup", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// WaitCleanup indicates an expected call of WaitCleanup.
func (mr *MockInterfaceMockRecorder) WaitCleanup(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WaitCleanup", reflect.TypeOf((*MockInterface)(nil).WaitCleanup), arg0)
}
