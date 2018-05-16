// Code generated by MockGen. DO NOT EDIT.
// Source: differ.go

package differ

import (
	gomock "github.com/golang/mock/gomock"
	api "github.com/turbinelabs/api"
	reflect "reflect"
)

// MockDiffer is a mock of Differ interface
type MockDiffer struct {
	ctrl     *gomock.Controller
	recorder *MockDifferMockRecorder
}

// MockDifferMockRecorder is the mock recorder for MockDiffer
type MockDifferMockRecorder struct {
	mock *MockDiffer
}

// NewMockDiffer creates a new mock instance
func NewMockDiffer(ctrl *gomock.Controller) *MockDiffer {
	mock := &MockDiffer{ctrl: ctrl}
	mock.recorder = &MockDifferMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDiffer) EXPECT() *MockDifferMockRecorder {
	return m.recorder
}

// Diff mocks base method
func (m *MockDiffer) Diff(proposed []api.Cluster, opts DiffOpts) ([]Diff, error) {
	ret := m.ctrl.Call(m, "Diff", proposed, opts)
	ret0, _ := ret[0].([]Diff)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Diff indicates an expected call of Diff
func (mr *MockDifferMockRecorder) Diff(proposed, opts interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Diff", reflect.TypeOf((*MockDiffer)(nil).Diff), proposed, opts)
}

// Patch mocks base method
func (m *MockDiffer) Patch(diffs []Diff) error {
	ret := m.ctrl.Call(m, "Patch", diffs)
	ret0, _ := ret[0].(error)
	return ret0
}

// Patch indicates an expected call of Patch
func (mr *MockDifferMockRecorder) Patch(diffs interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Patch", reflect.TypeOf((*MockDiffer)(nil).Patch), diffs)
}