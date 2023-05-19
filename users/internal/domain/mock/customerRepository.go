// Code generated by MockGen. DO NOT EDIT.
// Source: customer.go

// Package mock is a generated GoMock package.
package mock

import (
	model "meliarqsoft2/internal/domain/model"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
)

// MockICustomerRepository is a mock of ICustomerRepository interface.
type MockICustomerRepository struct {
	ctrl     *gomock.Controller
	recorder *MockICustomerRepositoryMockRecorder
}

// MockICustomerRepositoryMockRecorder is the mock recorder for MockICustomerRepository.
type MockICustomerRepositoryMockRecorder struct {
	mock *MockICustomerRepository
}

// NewMockICustomerRepository creates a new mock instance.
func NewMockICustomerRepository(ctrl *gomock.Controller) *MockICustomerRepository {
	mock := &MockICustomerRepository{ctrl: ctrl}
	mock.recorder = &MockICustomerRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICustomerRepository) EXPECT() *MockICustomerRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockICustomerRepository) Create(customer *model.Customer) (uuid.UUID, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", customer)
	ret0, _ := ret[0].(uuid.UUID)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockICustomerRepositoryMockRecorder) Create(customer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockICustomerRepository)(nil).Create), customer)
}

// Delete mocks base method.
func (m *MockICustomerRepository) Delete(ID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockICustomerRepositoryMockRecorder) Delete(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockICustomerRepository)(nil).Delete), ID)
}

// Find mocks base method.
func (m *MockICustomerRepository) Find(emailPattern string) ([]*model.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", emailPattern)
	ret0, _ := ret[0].([]*model.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Find indicates an expected call of Find.
func (mr *MockICustomerRepositoryMockRecorder) Find(emailPattern interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockICustomerRepository)(nil).Find), emailPattern)
}

// FindById mocks base method.
func (m *MockICustomerRepository) FindById(idCustomer uuid.UUID) (*model.Customer, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", idCustomer)
	ret0, _ := ret[0].(*model.Customer)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockICustomerRepositoryMockRecorder) FindById(idCustomer interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockICustomerRepository)(nil).FindById), idCustomer)
}

// Update mocks base method.
func (m *MockICustomerRepository) Update(ID uuid.UUID, name, surname, email string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ID, name, surname, email)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockICustomerRepositoryMockRecorder) Update(ID, name, surname, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockICustomerRepository)(nil).Update), ID, name, surname, email)
}
