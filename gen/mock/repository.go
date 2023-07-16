// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/minguu42/mtasks/pkg/entity"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateProject mocks base method.
func (m *MockRepository) CreateProject(ctx context.Context, userID, name, color string) (*entity.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProject", ctx, userID, name, color)
	ret0, _ := ret[0].(*entity.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateProject indicates an expected call of CreateProject.
func (mr *MockRepositoryMockRecorder) CreateProject(ctx, userID, name, color interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProject", reflect.TypeOf((*MockRepository)(nil).CreateProject), ctx, userID, name, color)
}

// CreateTask mocks base method.
func (m *MockRepository) CreateTask(ctx context.Context, projectID, title, content string, priority int, dueOn *time.Time) (*entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", ctx, projectID, title, content, priority, dueOn)
	ret0, _ := ret[0].(*entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockRepositoryMockRecorder) CreateTask(ctx, projectID, title, content, priority, dueOn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockRepository)(nil).CreateTask), ctx, projectID, title, content, priority, dueOn)
}

// DeleteProject mocks base method.
func (m *MockRepository) DeleteProject(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProject", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProject indicates an expected call of DeleteProject.
func (mr *MockRepositoryMockRecorder) DeleteProject(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProject", reflect.TypeOf((*MockRepository)(nil).DeleteProject), ctx, id)
}

// DeleteTask mocks base method.
func (m *MockRepository) DeleteTask(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteTask", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockRepositoryMockRecorder) DeleteTask(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockRepository)(nil).DeleteTask), ctx, id)
}

// GetProjectByID mocks base method.
func (m *MockRepository) GetProjectByID(ctx context.Context, id string) (*entity.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectByID", ctx, id)
	ret0, _ := ret[0].(*entity.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectByID indicates an expected call of GetProjectByID.
func (mr *MockRepositoryMockRecorder) GetProjectByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectByID", reflect.TypeOf((*MockRepository)(nil).GetProjectByID), ctx, id)
}

// GetProjectsByUserID mocks base method.
func (m *MockRepository) GetProjectsByUserID(ctx context.Context, userID, sort string, limit, offset int) ([]*entity.Project, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProjectsByUserID", ctx, userID, sort, limit, offset)
	ret0, _ := ret[0].([]*entity.Project)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProjectsByUserID indicates an expected call of GetProjectsByUserID.
func (mr *MockRepositoryMockRecorder) GetProjectsByUserID(ctx, userID, sort, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProjectsByUserID", reflect.TypeOf((*MockRepository)(nil).GetProjectsByUserID), ctx, userID, sort, limit, offset)
}

// GetTaskByID mocks base method.
func (m *MockRepository) GetTaskByID(ctx context.Context, id string) (*entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskByID", ctx, id)
	ret0, _ := ret[0].(*entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskByID indicates an expected call of GetTaskByID.
func (mr *MockRepositoryMockRecorder) GetTaskByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByID", reflect.TypeOf((*MockRepository)(nil).GetTaskByID), ctx, id)
}

// GetTasksByProjectID mocks base method.
func (m *MockRepository) GetTasksByProjectID(ctx context.Context, projectID, sort string, limit, offset int) ([]*entity.Task, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasksByProjectID", ctx, projectID, sort, limit, offset)
	ret0, _ := ret[0].([]*entity.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTasksByProjectID indicates an expected call of GetTasksByProjectID.
func (mr *MockRepositoryMockRecorder) GetTasksByProjectID(ctx, projectID, sort, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasksByProjectID", reflect.TypeOf((*MockRepository)(nil).GetTasksByProjectID), ctx, projectID, sort, limit, offset)
}

// GetUserByAPIKey mocks base method.
func (m *MockRepository) GetUserByAPIKey(ctx context.Context, apiKey string) (*entity.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByAPIKey", ctx, apiKey)
	ret0, _ := ret[0].(*entity.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByAPIKey indicates an expected call of GetUserByAPIKey.
func (mr *MockRepositoryMockRecorder) GetUserByAPIKey(ctx, apiKey interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByAPIKey", reflect.TypeOf((*MockRepository)(nil).GetUserByAPIKey), ctx, apiKey)
}

// UpdateProject mocks base method.
func (m *MockRepository) UpdateProject(ctx context.Context, p *entity.Project) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProject", ctx, p)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProject indicates an expected call of UpdateProject.
func (mr *MockRepositoryMockRecorder) UpdateProject(ctx, p interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProject", reflect.TypeOf((*MockRepository)(nil).UpdateProject), ctx, p)
}

// UpdateTask mocks base method.
func (m *MockRepository) UpdateTask(ctx context.Context, t *entity.Task) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", ctx, t)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockRepositoryMockRecorder) UpdateTask(ctx, t interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockRepository)(nil).UpdateTask), ctx, t)
}
