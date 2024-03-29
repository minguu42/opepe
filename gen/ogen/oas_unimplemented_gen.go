// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateProject implements createProject operation.
//
// 新しいプロジェクトを作成する.
//
// POST /projects
func (UnimplementedHandler) CreateProject(ctx context.Context, req *CreateProjectReq) (r *Project, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateTask implements createTask operation.
//
// 新しいタスクを作成する.
//
// POST /projects/{projectID}/tasks
func (UnimplementedHandler) CreateTask(ctx context.Context, req *CreateTaskReq, params CreateTaskParams) (r *Task, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteProject implements deleteProject operation.
//
// プロジェクトを削除する.
//
// DELETE /projects/{projectID}
func (UnimplementedHandler) DeleteProject(ctx context.Context, params DeleteProjectParams) error {
	return ht.ErrNotImplemented
}

// DeleteTask implements deleteTask operation.
//
// タスクを削除する.
//
// DELETE /projects/{projectID}/tasks/{taskID}
func (UnimplementedHandler) DeleteTask(ctx context.Context, params DeleteTaskParams) error {
	return ht.ErrNotImplemented
}

// GetHealth implements getHealth operation.
//
// サーバの状態を取得する.
//
// GET /health
func (UnimplementedHandler) GetHealth(ctx context.Context) (r *GetHealthOK, _ error) {
	return r, ht.ErrNotImplemented
}

// ListProjects implements listProjects operation.
//
// プロジェクト一覧を取得する.
//
// GET /projects
func (UnimplementedHandler) ListProjects(ctx context.Context, params ListProjectsParams) (r *Projects, _ error) {
	return r, ht.ErrNotImplemented
}

// ListTasks implements listTasks operation.
//
// タスク一覧を取得する.
//
// GET /projects/{projectID}/tasks
func (UnimplementedHandler) ListTasks(ctx context.Context, params ListTasksParams) (r *Tasks, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateProject implements updateProject operation.
//
// プロジェクトを更新する.
//
// PATCH /projects/{projectID}
func (UnimplementedHandler) UpdateProject(ctx context.Context, req *UpdateProjectReq, params UpdateProjectParams) (r *Project, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateTask implements updateTask operation.
//
// タスクを更新する.
//
// PATCH /projects/{projectID}/tasks/{taskID}
func (UnimplementedHandler) UpdateTask(ctx context.Context, req *UpdateTaskReq, params UpdateTaskParams) (r *Task, _ error) {
	return r, ht.ErrNotImplemented
}

// NewError creates *ErrorStatusCode from error returned by handler.
//
// Used for common default response.
func (UnimplementedHandler) NewError(ctx context.Context, err error) (r *ErrorStatusCode) {
	r = new(ErrorStatusCode)
	return r
}
