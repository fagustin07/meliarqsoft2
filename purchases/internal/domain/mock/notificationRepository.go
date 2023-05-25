package mock

import (
	"github.com/golang/mock/gomock"
	"meliarqsoft2/internal/domain/model"
	"reflect"
)

// MockINotificationRepository is a mock of INotificationService interface.
type MockINotificationRepository struct {
	ctrl     *gomock.Controller
	recorder *MockINotificationRepositoryMockRecorder
}

// MockINotificationRepositoryMockRecorder is the mock recorder for MockINotificationRepository.
type MockINotificationRepositoryMockRecorder struct {
	mock *MockINotificationRepository
}

// NewMockINotificationRepository creates a new mock instance.
func NewMockINotificationRepository(ctrl *gomock.Controller) *MockINotificationRepository {
	mock := &MockINotificationRepository{ctrl: ctrl}
	mock.recorder = &MockINotificationRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockINotificationRepository) EXPECT() *MockINotificationRepositoryMockRecorder {
	return m.recorder
}

// Send mocks base method.
func (m *MockINotificationRepository) Send(notification *model.Notification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Send", notification)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Create.
func (mr *MockINotificationRepositoryMockRecorder) Send(notification interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockINotificationRepository)(nil).Send), notification)
}
