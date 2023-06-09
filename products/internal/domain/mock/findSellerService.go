// Code generated by MockGen. DO NOT EDIT.
// Source: findSellerById.go

// Package mock is a generated GoMock package.
package mock

import (
	model "meliarqsoft2/internal/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockIFindSellerByIdService is a mock of IFindSellerByIdService interface.
type MockIFindSellerByIdService struct {
	ctrl     *gomock.Controller
	recorder *MockIFindSellerByIdServiceMockRecorder
}

// MockIFindSellerByIdServiceMockRecorder is the mock recorder for MockIFindSellerByIdService.
type MockIFindSellerByIdServiceMockRecorder struct {
	mock *MockIFindSellerByIdService
}

// NewMockIFindSellerByIdService creates a new mock instance.
func NewMockIFindSellerByIdService(ctrl *gomock.Controller) *MockIFindSellerByIdService {
	mock := &MockIFindSellerByIdService{ctrl: ctrl}
	mock.recorder = &MockIFindSellerByIdServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIFindSellerByIdService) EXPECT() *MockIFindSellerByIdServiceMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIFindSellerByIdService) Execute(id uuid.UUID) (model.SellerDTO, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", id)
	ret0, _ := ret[0].(model.SellerDTO)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Execute indicates an expected call of Execute.
func (mr *MockIFindSellerByIdServiceMockRecorder) Execute(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIFindSellerByIdService)(nil).Execute), id)
}
