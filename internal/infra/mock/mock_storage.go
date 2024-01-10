// Source: internal/domain/storage/storage.go

// Package mock_storage is a generated GoMock package.
package mock_storage

import (
	context "context"
	reflect "reflect"

	domain "github.com/STUD-IT-team/bmstu-stud-web-backend/internal/domain"
	gomock "github.com/golang/mock/gomock"
)

// MockStorage is a mock of Storage interface.
type MockStorage struct {
	ctrl     *gomock.Controller
	recorder *MockStorageMockRecorder
}

// MockStorageMockRecorder is the mock recorder for MockStorage.
type MockStorageMockRecorder struct {
	mock *MockStorage
}

// NewMockStorage creates a new mock instance.
func NewMockStorage(ctrl *gomock.Controller) *MockStorage {
	mock := &MockStorage{ctrl: ctrl}
	mock.recorder = &MockStorageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStorage) EXPECT() *MockStorageMockRecorder {
	return m.recorder
}

// GetAllFeed mocks base method.
func (m *MockStorage) GetAllFeed(ctx context.Context) ([]domain.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllFeed", ctx)
	ret0, _ := ret[0].([]domain.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllFeed indicates an expected call of GetAllFeed.
func (mr *MockStorageMockRecorder) GetAllFeed(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllFeed", reflect.TypeOf((*MockStorage)(nil).GetAllFeed), ctx)
}

// GetFeed mocks base method.
func (m *MockStorage) GetFeed(ctx context.Context, id int) (domain.Feed, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFeed", ctx, id)
	ret0, _ := ret[0].(domain.Feed)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFeed indicates an expected call of GetFeed.
func (mr *MockStorageMockRecorder) GetFeed(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFeed", reflect.TypeOf((*MockStorage)(nil).GetFeed), ctx, id)
}
