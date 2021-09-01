// Code generated by MockGen. DO NOT EDIT.
// Source: ozonva/ova-task-api/internal/repo (interfaces: Repo)

// Package mocks is a generated GoMock package.
package mocks

import (
	tasks "ozonva/ova-task-api/pkg/entities/tasks"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddTasks mocks base method.
func (m *MockRepo) AddTasks(arg0 []tasks.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddTasks", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTasks indicates an expected call of AddTasks.
func (mr *MockRepoMockRecorder) AddTasks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTasks", reflect.TypeOf((*MockRepo)(nil).AddTasks), arg0)
}

// DescribeTasks mocks base method.
func (m *MockRepo) DescribeTasks(arg0 uint64) (*tasks.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeTasks", arg0)
	ret0, _ := ret[0].(*tasks.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeTasks indicates an expected call of DescribeTasks.
func (mr *MockRepoMockRecorder) DescribeTasks(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeTasks", reflect.TypeOf((*MockRepo)(nil).DescribeTasks), arg0)
}

// ListTasks mocks base method.
func (m *MockRepo) ListTasks(arg0, arg1 uint64) ([]tasks.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListTasks", arg0, arg1)
	ret0, _ := ret[0].([]tasks.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasks indicates an expected call of ListTasks.
func (mr *MockRepoMockRecorder) ListTasks(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasks", reflect.TypeOf((*MockRepo)(nil).ListTasks), arg0, arg1)
}

// RemoveTask mocks base method.
func (m *MockRepo) RemoveTask(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveTask indicates an expected call of RemoveTask.
func (mr *MockRepoMockRecorder) RemoveTask(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveTask", reflect.TypeOf((*MockRepo)(nil).RemoveTask), arg0)
}
