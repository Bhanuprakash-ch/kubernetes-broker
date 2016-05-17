// Automatically generated by MockGen. DO NOT EDIT!
// Source: k8s/k8screator_rest_api.go

package k8s

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of K8sCreatorRest interface
type MockK8sCreatorRest struct {
	ctrl     *gomock.Controller
	recorder *_MockK8sCreatorRestRecorder
}

// Recorder for MockK8sCreatorRest (not exported)
type _MockK8sCreatorRestRecorder struct {
	mock *MockK8sCreatorRest
}

func NewMockK8sCreatorRest(ctrl *gomock.Controller) *MockK8sCreatorRest {
	mock := &MockK8sCreatorRest{ctrl: ctrl}
	mock.recorder = &_MockK8sCreatorRestRecorder{mock}
	return mock
}

func (_m *MockK8sCreatorRest) EXPECT() *_MockK8sCreatorRestRecorder {
	return _m.recorder
}

func (_m *MockK8sCreatorRest) DeleteCluster(org string) error {
	ret := _m.ctrl.Call(_m, "DeleteCluster", org)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockK8sCreatorRestRecorder) DeleteCluster(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteCluster", arg0)
}

func (_m *MockK8sCreatorRest) GetCluster(org string) (int, K8sClusterCredential, error) {
	ret := _m.ctrl.Call(_m, "GetCluster", org)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(K8sClusterCredential)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (_mr *_MockK8sCreatorRestRecorder) GetCluster(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetCluster", arg0)
}

func (_m *MockK8sCreatorRest) GetOrCreateCluster(org string) (K8sClusterCredential, error) {
	ret := _m.ctrl.Call(_m, "GetOrCreateCluster", org)
	ret0, _ := ret[0].(K8sClusterCredential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockK8sCreatorRestRecorder) GetOrCreateCluster(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetOrCreateCluster", arg0)
}

func (_m *MockK8sCreatorRest) PostCluster(org string) (int, error) {
	ret := _m.ctrl.Call(_m, "PostCluster", org)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockK8sCreatorRestRecorder) PostCluster(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "PostCluster", arg0)
}

func (_m *MockK8sCreatorRest) GetClusters() ([]K8sClusterCredential, error) {
	ret := _m.ctrl.Call(_m, "GetClusters")
	ret0, _ := ret[0].([]K8sClusterCredential)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockK8sCreatorRestRecorder) GetClusters() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetClusters")
}