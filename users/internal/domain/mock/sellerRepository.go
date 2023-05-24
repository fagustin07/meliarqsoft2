// Code generated by MockGen. DO NOT EDIT.
// Source: seller.go

// Package mock is a generated GoMock package.
package mock

import (
	model "meliarqsoft2/internal/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockISellerRepository is a mock of ISellerRepository interface.
type MockISellerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockISellerRepositoryMockRecorder
}

// MockISellerRepositoryMockRecorder is the mock recorder for MockISellerRepository.
type MockISellerRepositoryMockRecorder struct {
	mock *MockISellerRepository
}

// NewMockISellerRepository creates a new mock instance.
func NewMockISellerRepository(ctrl *gomock.Controller) *MockISellerRepository {
	mock := &MockISellerRepository{ctrl: ctrl}
	mock.recorder = &MockISellerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockISellerRepository) EXPECT() *MockISellerRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockISellerRepository) Create(seller *model.Seller) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", seller)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockISellerRepositoryMockRecorder) Create(seller interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockISellerRepository)(nil).Create), seller)
}

// Delete mocks base method.
func (m *MockISellerRepository) Delete(ID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockISellerRepositoryMockRecorder) Delete(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockISellerRepository)(nil).Delete), ID)
}

// DeleteAll mocks base method.
func (m *MockISellerRepository) DeleteAll() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAll")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAll indicates an expected call of DeleteAll.
func (mr *MockISellerRepositoryMockRecorder) DeleteAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAll", reflect.TypeOf((*MockISellerRepository)(nil).DeleteAll))
}

// Find mocks base method.
func (m *MockISellerRepository) Find(businessName string) ([]*model.Seller, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", businessName)
	ret0, _ := ret[0].([]*model.Seller)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockISellerRepositoryMockRecorder) Find(businessName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockISellerRepository)(nil).Find), businessName)
}

// FindById mocks base method.
func (m *MockISellerRepository) FindById(idSeller uuid.UUID) (*model.Seller, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", idSeller)
	ret0, _ := ret[0].(*model.Seller)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockISellerRepositoryMockRecorder) FindById(idSeller interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockISellerRepository)(nil).FindById), idSeller)
}

// Restore mocks base method.
func (m *MockISellerRepository) Restore(ID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Restore", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Restore indicates an expected call of Restore.
func (mr *MockISellerRepositoryMockRecorder) Restore(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Restore", reflect.TypeOf((*MockISellerRepository)(nil).Restore), ID)
}

// Update mocks base method.
func (m *MockISellerRepository) Update(id uuid.UUID, businessName, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, businessName, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockISellerRepositoryMockRecorder) Update(id, businessName, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockISellerRepository)(nil).Update), id, businessName, email)
}
