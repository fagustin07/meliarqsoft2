// Code generated by MockGen. DO NOT EDIT.
// Source: deletePurchase.go

// Package mock is a generated GoMock package.
package mock

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockIDeletePurchasesByProductsService is a mock of IDeletePurchasesByProductsService interface.
type MockIDeletePurchasesByProductsService struct {
	ctrl     *gomock.Controller
	recorder *MockIDeletePurchasesByProductsServiceMockRecorder
}

// MockIDeletePurchasesByProductsServiceMockRecorder is the mock recorder for MockIDeletePurchasesByProductsService.
type MockIDeletePurchasesByProductsServiceMockRecorder struct {
	mock *MockIDeletePurchasesByProductsService
}

// NewMockIDeletePurchasesByProductsService creates a new mock instance.
func NewMockIDeletePurchasesByProductsService(ctrl *gomock.Controller) *MockIDeletePurchasesByProductsService {
	mock := &MockIDeletePurchasesByProductsService{ctrl: ctrl}
	mock.recorder = &MockIDeletePurchasesByProductsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIDeletePurchasesByProductsService) EXPECT() *MockIDeletePurchasesByProductsServiceMockRecorder {
	return m.recorder
}

// Execute mocks base method.
func (m *MockIDeletePurchasesByProductsService) Execute(IDs []uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Execute", IDs)
	ret0, _ := ret[0].(error)
	return ret0
}

// Execute indicates an expected call of Execute.
func (mr *MockIDeletePurchasesByProductsServiceMockRecorder) Execute(IDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Execute", reflect.TypeOf((*MockIDeletePurchasesByProductsService)(nil).Execute), IDs)
}