// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/arana-db/arana/pkg/proto (interfaces: Field,Row,KeyedRow,Dataset,Result)

// Package testdata is a generated GoMock package.
package testdata

import (
	io "io"
	reflect "reflect"

	proto "github.com/arana-db/arana/pkg/proto"
	gomock "github.com/golang/mock/gomock"
)

// MockField is a mock of Field interface.
type MockField struct {
	ctrl     *gomock.Controller
	recorder *MockFieldMockRecorder
}

// MockFieldMockRecorder is the mock recorder for MockField.
type MockFieldMockRecorder struct {
	mock *MockField
}

// NewMockField creates a new mock instance.
func NewMockField(ctrl *gomock.Controller) *MockField {
	mock := &MockField{ctrl: ctrl}
	mock.recorder = &MockFieldMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockField) EXPECT() *MockFieldMockRecorder {
	return m.recorder
}

// DatabaseTypeName mocks base method.
func (m *MockField) DatabaseTypeName() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DatabaseTypeName")
	ret0, _ := ret[0].(string)
	return ret0
}

// DatabaseTypeName indicates an expected call of DatabaseTypeName.
func (mr *MockFieldMockRecorder) DatabaseTypeName() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DatabaseTypeName", reflect.TypeOf((*MockField)(nil).DatabaseTypeName))
}

// DecimalSize mocks base method.
func (m *MockField) DecimalSize() (int64, int64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecimalSize")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(bool)
	return ret0, ret1, ret2
}

// DecimalSize indicates an expected call of DecimalSize.
func (mr *MockFieldMockRecorder) DecimalSize() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecimalSize", reflect.TypeOf((*MockField)(nil).DecimalSize))
}

// Length mocks base method.
func (m *MockField) Length() (int64, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Length indicates an expected call of Length.
func (mr *MockFieldMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockField)(nil).Length))
}

// Name mocks base method.
func (m *MockField) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockFieldMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockField)(nil).Name))
}

// Nullable mocks base method.
func (m *MockField) Nullable() (bool, bool) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Nullable")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// Nullable indicates an expected call of Nullable.
func (mr *MockFieldMockRecorder) Nullable() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Nullable", reflect.TypeOf((*MockField)(nil).Nullable))
}

// ScanType mocks base method.
func (m *MockField) ScanType() reflect.Type {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ScanType")
	ret0, _ := ret[0].(reflect.Type)
	return ret0
}

// ScanType indicates an expected call of ScanType.
func (mr *MockFieldMockRecorder) ScanType() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ScanType", reflect.TypeOf((*MockField)(nil).ScanType))
}

// MockRow is a mock of Row interface.
type MockRow struct {
	ctrl     *gomock.Controller
	recorder *MockRowMockRecorder
}

// MockRowMockRecorder is the mock recorder for MockRow.
type MockRowMockRecorder struct {
	mock *MockRow
}

// NewMockRow creates a new mock instance.
func NewMockRow(ctrl *gomock.Controller) *MockRow {
	mock := &MockRow{ctrl: ctrl}
	mock.recorder = &MockRowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRow) EXPECT() *MockRowMockRecorder {
	return m.recorder
}

// IsBinary mocks base method.
func (m *MockRow) IsBinary() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBinary")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsBinary indicates an expected call of IsBinary.
func (mr *MockRowMockRecorder) IsBinary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBinary", reflect.TypeOf((*MockRow)(nil).IsBinary))
}

// Length mocks base method.
func (m *MockRow) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockRowMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockRow)(nil).Length))
}

// Scan mocks base method.
func (m *MockRow) Scan(arg0 []proto.Value) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockRowMockRecorder) Scan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockRow)(nil).Scan), arg0)
}

// WriteTo mocks base method.
func (m *MockRow) WriteTo(arg0 io.Writer) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTo", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteTo indicates an expected call of WriteTo.
func (mr *MockRowMockRecorder) WriteTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTo", reflect.TypeOf((*MockRow)(nil).WriteTo), arg0)
}

// MockKeyedRow is a mock of KeyedRow interface.
type MockKeyedRow struct {
	ctrl     *gomock.Controller
	recorder *MockKeyedRowMockRecorder
}

// MockKeyedRowMockRecorder is the mock recorder for MockKeyedRow.
type MockKeyedRowMockRecorder struct {
	mock *MockKeyedRow
}

// NewMockKeyedRow creates a new mock instance.
func NewMockKeyedRow(ctrl *gomock.Controller) *MockKeyedRow {
	mock := &MockKeyedRow{ctrl: ctrl}
	mock.recorder = &MockKeyedRowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockKeyedRow) EXPECT() *MockKeyedRowMockRecorder {
	return m.recorder
}

// Fields mocks base method.
func (m *MockKeyedRow) Fields() []proto.Field {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fields")
	ret0, _ := ret[0].([]proto.Field)
	return ret0
}

// Fields indicates an expected call of Fields.
func (mr *MockKeyedRowMockRecorder) Fields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fields", reflect.TypeOf((*MockKeyedRow)(nil).Fields))
}

// Get mocks base method.
func (m *MockKeyedRow) Get(arg0 string) (proto.Value, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].(proto.Value)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockKeyedRowMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockKeyedRow)(nil).Get), arg0)
}

// IsBinary mocks base method.
func (m *MockKeyedRow) IsBinary() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsBinary")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsBinary indicates an expected call of IsBinary.
func (mr *MockKeyedRowMockRecorder) IsBinary() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsBinary", reflect.TypeOf((*MockKeyedRow)(nil).IsBinary))
}

// Length mocks base method.
func (m *MockKeyedRow) Length() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Length")
	ret0, _ := ret[0].(int)
	return ret0
}

// Length indicates an expected call of Length.
func (mr *MockKeyedRowMockRecorder) Length() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Length", reflect.TypeOf((*MockKeyedRow)(nil).Length))
}

// Scan mocks base method.
func (m *MockKeyedRow) Scan(arg0 []proto.Value) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Scan indicates an expected call of Scan.
func (mr *MockKeyedRowMockRecorder) Scan(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockKeyedRow)(nil).Scan), arg0)
}

// WriteTo mocks base method.
func (m *MockKeyedRow) WriteTo(arg0 io.Writer) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WriteTo", arg0)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WriteTo indicates an expected call of WriteTo.
func (mr *MockKeyedRowMockRecorder) WriteTo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WriteTo", reflect.TypeOf((*MockKeyedRow)(nil).WriteTo), arg0)
}

// MockDataset is a mock of Dataset interface.
type MockDataset struct {
	ctrl     *gomock.Controller
	recorder *MockDatasetMockRecorder
}

// MockDatasetMockRecorder is the mock recorder for MockDataset.
type MockDatasetMockRecorder struct {
	mock *MockDataset
}

// NewMockDataset creates a new mock instance.
func NewMockDataset(ctrl *gomock.Controller) *MockDataset {
	mock := &MockDataset{ctrl: ctrl}
	mock.recorder = &MockDatasetMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataset) EXPECT() *MockDatasetMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockDataset) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockDatasetMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDataset)(nil).Close))
}

// Fields mocks base method.
func (m *MockDataset) Fields() ([]proto.Field, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Fields")
	ret0, _ := ret[0].([]proto.Field)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Fields indicates an expected call of Fields.
func (mr *MockDatasetMockRecorder) Fields() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Fields", reflect.TypeOf((*MockDataset)(nil).Fields))
}

// Next mocks base method.
func (m *MockDataset) Next() (proto.Row, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(proto.Row)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next.
func (mr *MockDatasetMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockDataset)(nil).Next))
}

// MockResult is a mock of Result interface.
type MockResult struct {
	ctrl     *gomock.Controller
	recorder *MockResultMockRecorder
}

// MockResultMockRecorder is the mock recorder for MockResult.
type MockResultMockRecorder struct {
	mock *MockResult
}

// NewMockResult creates a new mock instance.
func NewMockResult(ctrl *gomock.Controller) *MockResult {
	mock := &MockResult{ctrl: ctrl}
	mock.recorder = &MockResultMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResult) EXPECT() *MockResultMockRecorder {
	return m.recorder
}

// Dataset mocks base method.
func (m *MockResult) Dataset() (proto.Dataset, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Dataset")
	ret0, _ := ret[0].(proto.Dataset)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Dataset indicates an expected call of Dataset.
func (mr *MockResultMockRecorder) Dataset() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Dataset", reflect.TypeOf((*MockResult)(nil).Dataset))
}

// LastInsertId mocks base method.
func (m *MockResult) LastInsertId() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastInsertId")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LastInsertId indicates an expected call of LastInsertId.
func (mr *MockResultMockRecorder) LastInsertId() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastInsertId", reflect.TypeOf((*MockResult)(nil).LastInsertId))
}

// RowsAffected mocks base method.
func (m *MockResult) RowsAffected() (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RowsAffected")
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RowsAffected indicates an expected call of RowsAffected.
func (mr *MockResultMockRecorder) RowsAffected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RowsAffected", reflect.TypeOf((*MockResult)(nil).RowsAffected))
}
