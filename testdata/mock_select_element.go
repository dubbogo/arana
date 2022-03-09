// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/runtime/xxast/ast_select_element2.go

// Package mock_xxast is a generated GoMock package.
package testdata

import (
	reflect "reflect"
)

import (
	xxast "github.com/dubbogo/arana/pkg/runtime/xxast"

	gomock "github.com/golang/mock/gomock"
)

// MockSelectElementFunction is a mock of SelectElement interface.
type MockSelectElementFunction struct {
	ctrl     *gomock.Controller
	recorder *MockSelectElementFunctionMockRecorder
}

// MockSelectElementFunctionMockRecorder is the mock recorder for MockSelectElementFunction.
type MockSelectElementFunctionMockRecorder struct {
	mock *MockSelectElementFunction
}

// NewMockSelectElementFunction creates a new mock instance.
func NewMockSelectElementFunction(ctrl *gomock.Controller) *MockSelectElementFunction {
	mock := &MockSelectElementFunction{ctrl: ctrl}
	mock.recorder = &MockSelectElementFunctionMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSelectElementFunction) EXPECT() *MockSelectElementFunctionMockRecorder {
	return m.recorder
}

// Alias mocks base method.
func (m *MockSelectElementFunction) Alias() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Alias")
	ret0, _ := ret[0].(string)
	return ret0
}

// Alias indicates an expected call of Alias.
func (mr *MockSelectElementFunctionMockRecorder) Alias() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Alias", reflect.TypeOf((*MockSelectElementFunction)(nil).Alias))
}

// Function mocks base method.
func (m *MockSelectElementFunction) Function() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Function")
	ret0, _ := ret[0].(string)
	return ret0
}

// Function indicates an expected call of Alias.
func (mr *MockSelectElementFunctionMockRecorder) Function() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Function", reflect.TypeOf((*MockSelectElementFunction)(nil).Alias))
}

// InTables mocks base method.
func (m *MockSelectElementFunction) InTables(tables map[string]struct{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InTables", tables)
	ret0, _ := ret[0].(error)
	return ret0
}

// InTables indicates an expected call of InTables.
func (mr *MockSelectElementFunctionMockRecorder) InTables(tables interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InTables", reflect.TypeOf((*MockSelectElementFunction)(nil).InTables), tables)
}

// Mode mocks base method.
func (m *MockSelectElementFunction) Mode() xxast.SelMode {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Mode")
	ret0, _ := ret[0].(xxast.SelMode)
	return ret0
}

// Mode indicates an expected call of Mode.
func (mr *MockSelectElementFunctionMockRecorder) Mode() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Mode", reflect.TypeOf((*MockSelectElementFunction)(nil).Mode))
}

// ToSelectString mocks base method.
func (m *MockSelectElementFunction) ToSelectString() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ToSelectString")
	ret0, _ := ret[0].(string)
	return ret0
}

// ToSelectString indicates an expected call of ToSelectString.
func (mr *MockSelectElementFunctionMockRecorder) ToSelectString() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ToSelectString", reflect.TypeOf((*MockSelectElementFunction)(nil).ToSelectString))
}
