// Code generated by MockGen. DO NOT EDIT.
// Source: client/client_interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	client "XMLResReq/client"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockXMLReq is a mock of XMLReq interface.
type MockXMLReq struct {
	ctrl     *gomock.Controller
	recorder *MockXMLReqMockRecorder
}

// MockXMLReqMockRecorder is the mock recorder for MockXMLReq.
type MockXMLReqMockRecorder struct {
	mock *MockXMLReq
}

// NewMockXMLReq creates a new mock instance.
func NewMockXMLReq(ctrl *gomock.Controller) *MockXMLReq {
	mock := &MockXMLReq{ctrl: ctrl}
	mock.recorder = &MockXMLReqMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockXMLReq) EXPECT() *MockXMLReqMockRecorder {
	return m.recorder
}

// XMLRR mocks base method.
func (m *MockXMLReq) XMLRR() (client.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "XMLRR")
	ret0, _ := ret[0].(client.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// XMLRR indicates an expected call of XMLRR.
func (mr *MockXMLReqMockRecorder) XMLRR() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "XMLRR", reflect.TypeOf((*MockXMLReq)(nil).XMLRR))
}
