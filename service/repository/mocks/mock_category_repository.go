// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/saipulmuiz/go-project-starter/service (interfaces: CategoryRepository)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
	models "github.com/saipulmuiz/go-project-starter/models"
	serror "github.com/saipulmuiz/go-project-starter/pkg/serror"
)

// MockCategoryRepository is a mock of CategoryRepository interface.
type MockCategoryRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCategoryRepositoryMockRecorder
}

// MockCategoryRepositoryMockRecorder is the mock recorder for MockCategoryRepository.
type MockCategoryRepositoryMockRecorder struct {
	mock *MockCategoryRepository
}

// NewMockCategoryRepository creates a new mock instance.
func NewMockCategoryRepository(ctrl *gomock.Controller) *MockCategoryRepository {
	mock := &MockCategoryRepository{ctrl: ctrl}
	mock.recorder = &MockCategoryRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCategoryRepository) EXPECT() *MockCategoryRepositoryMockRecorder {
	return m.recorder
}

// CreateCategory mocks base method.
func (m *MockCategoryRepository) CreateCategory(arg0 context.Context, arg1 models.CreateCategoryRequest) (int64, serror.SError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateCategory", arg0, arg1)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(serror.SError)
	return ret0, ret1
}

// CreateCategory indicates an expected call of CreateCategory.
func (mr *MockCategoryRepositoryMockRecorder) CreateCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateCategory", reflect.TypeOf((*MockCategoryRepository)(nil).CreateCategory), arg0, arg1)
}

// DeleteCategory mocks base method.
func (m *MockCategoryRepository) DeleteCategory(arg0 context.Context, arg1 int64) serror.SError {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteCategory", arg0, arg1)
	ret0, _ := ret[0].(serror.SError)
	return ret0
}

// DeleteCategory indicates an expected call of DeleteCategory.
func (mr *MockCategoryRepositoryMockRecorder) DeleteCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteCategory", reflect.TypeOf((*MockCategoryRepository)(nil).DeleteCategory), arg0, arg1)
}

// GetCategories mocks base method.
func (m *MockCategoryRepository) GetCategories(arg0 context.Context, arg1 models.GetCategoryRequest) ([]models.Category, serror.SError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategories", arg0, arg1)
	ret0, _ := ret[0].([]models.Category)
	ret1, _ := ret[1].(serror.SError)
	return ret0, ret1
}

// GetCategories indicates an expected call of GetCategories.
func (mr *MockCategoryRepositoryMockRecorder) GetCategories(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategories", reflect.TypeOf((*MockCategoryRepository)(nil).GetCategories), arg0, arg1)
}

// GetCategoryByID mocks base method.
func (m *MockCategoryRepository) GetCategoryByID(arg0 context.Context, arg1 int64) (models.Category, serror.SError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCategoryByID", arg0, arg1)
	ret0, _ := ret[0].(models.Category)
	ret1, _ := ret[1].(serror.SError)
	return ret0, ret1
}

// GetCategoryByID indicates an expected call of GetCategoryByID.
func (mr *MockCategoryRepositoryMockRecorder) GetCategoryByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCategoryByID", reflect.TypeOf((*MockCategoryRepository)(nil).GetCategoryByID), arg0, arg1)
}

// UpdateCategoryByID mocks base method.
func (m *MockCategoryRepository) UpdateCategoryByID(arg0 context.Context, arg1 *sqlx.DB, arg2 models.UpdateCategoryRequest) (models.Category, serror.SError) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateCategoryByID", arg0, arg1, arg2)
	ret0, _ := ret[0].(models.Category)
	ret1, _ := ret[1].(serror.SError)
	return ret0, ret1
}

// UpdateCategoryByID indicates an expected call of UpdateCategoryByID.
func (mr *MockCategoryRepositoryMockRecorder) UpdateCategoryByID(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateCategoryByID", reflect.TypeOf((*MockCategoryRepository)(nil).UpdateCategoryByID), arg0, arg1, arg2)
}
