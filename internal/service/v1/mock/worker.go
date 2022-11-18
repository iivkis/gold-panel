// Code generated by MockGen. DO NOT EDIT.
// Source: gold-panel/internal/service/v1 (interfaces: IWorkerService)

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	service "gold-panel/internal/service/v1"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIWorkerService is a mock of IWorkerService interface.
type MockIWorkerService struct {
	ctrl     *gomock.Controller
	recorder *MockIWorkerServiceMockRecorder
}

// MockIWorkerServiceMockRecorder is the mock recorder for MockIWorkerService.
type MockIWorkerServiceMockRecorder struct {
	mock *MockIWorkerService
}

// NewMockIWorkerService creates a new mock instance.
func NewMockIWorkerService(ctrl *gomock.Controller) *MockIWorkerService {
	mock := &MockIWorkerService{ctrl: ctrl}
	mock.recorder = &MockIWorkerServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIWorkerService) EXPECT() *MockIWorkerServiceMockRecorder {
	return m.recorder
}

// WorkerAdd mocks base method.
func (m *MockIWorkerService) WorkerAdd(arg0 context.Context, arg1 *service.WorkerAddDTO) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkerAdd", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WorkerAdd indicates an expected call of WorkerAdd.
func (mr *MockIWorkerServiceMockRecorder) WorkerAdd(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkerAdd", reflect.TypeOf((*MockIWorkerService)(nil).WorkerAdd), arg0, arg1)
}

// WorkerGetToken mocks base method.
func (m *MockIWorkerService) WorkerGetToken(arg0 context.Context, arg1 int64) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WorkerGetToken", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// WorkerGetToken indicates an expected call of WorkerGetToken.
func (mr *MockIWorkerServiceMockRecorder) WorkerGetToken(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WorkerGetToken", reflect.TypeOf((*MockIWorkerService)(nil).WorkerGetToken), arg0, arg1)
}
