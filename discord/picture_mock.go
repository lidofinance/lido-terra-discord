// Code generated by MockGen. DO NOT EDIT.
// Source: ./picture.go

// Package discord is a generated GoMock package.
package discord

import (
	io "io"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPicture is a mock of Picture interface.
type MockPicture struct {
	ctrl     *gomock.Controller
	recorder *MockPictureMockRecorder
}

// MockPictureMockRecorder is the mock recorder for MockPicture.
type MockPictureMockRecorder struct {
	mock *MockPicture
}

// NewMockPicture creates a new mock instance.
func NewMockPicture(ctrl *gomock.Controller) *MockPicture {
	mock := &MockPicture{ctrl: ctrl}
	mock.recorder = &MockPictureMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPicture) EXPECT() *MockPictureMockRecorder {
	return m.recorder
}

// Body mocks base method.
func (m *MockPicture) Body() io.Reader {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Body")
	ret0, _ := ret[0].(io.Reader)
	return ret0
}

// Body indicates an expected call of Body.
func (mr *MockPictureMockRecorder) Body() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Body", reflect.TypeOf((*MockPicture)(nil).Body))
}

// Name mocks base method.
func (m *MockPicture) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name.
func (mr *MockPictureMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockPicture)(nil).Name))
}
