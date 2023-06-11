// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"

	ht "github.com/ogen-go/ogen/http"
)

// UnimplementedHandler is no-op Handler which returns http.ErrNotImplemented.
type UnimplementedHandler struct{}

var _ Handler = UnimplementedHandler{}

// CreateProject implements CreateProject operation.
//
// 新しいプロジェクトを作成する.
//
// POST /projects
func (UnimplementedHandler) CreateProject(ctx context.Context, req *CreateProjectReq, params CreateProjectParams) (r CreateProjectRes, _ error) {
	return r, ht.ErrNotImplemented
}

// CreateTask implements CreateTask operation.
//
// 新しいタスクを作成する.
//
// POST /projects/{projectID}/tasks
func (UnimplementedHandler) CreateTask(ctx context.Context, req *CreateTaskReq, params CreateTaskParams) (r CreateTaskRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteProject implements DeleteProject operation.
//
// プロジェクトを削除する.
//
// DELETE /projects/{projectID}
func (UnimplementedHandler) DeleteProject(ctx context.Context, params DeleteProjectParams) (r DeleteProjectRes, _ error) {
	return r, ht.ErrNotImplemented
}

// DeleteTask implements DeleteTask operation.
//
// タスクを削除する.
//
// DELETE /projects/{projectID}/tasks/{taskID}
func (UnimplementedHandler) DeleteTask(ctx context.Context, params DeleteTaskParams) (r DeleteTaskRes, _ error) {
	return r, ht.ErrNotImplemented
}

// GetHealth implements GetHealth operation.
//
// サーバの状態を取得する.
//
// GET /health
func (UnimplementedHandler) GetHealth(ctx context.Context) (r GetHealthRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListProjects implements ListProjects operation.
//
// 作成日時の降順で取得する。.
//
// GET /projects
func (UnimplementedHandler) ListProjects(ctx context.Context, params ListProjectsParams) (r ListProjectsRes, _ error) {
	return r, ht.ErrNotImplemented
}

// ListTasks implements ListTasks operation.
//
// 作成日時の降順で取得する。.
//
// GET /projects/{projectID}/tasks
func (UnimplementedHandler) ListTasks(ctx context.Context, params ListTasksParams) (r ListTasksRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateProject implements UpdateProject operation.
//
// プロジェクトを更新する.
//
// PATCH /projects/{projectID}
func (UnimplementedHandler) UpdateProject(ctx context.Context, req *UpdateProjectReq, params UpdateProjectParams) (r UpdateProjectRes, _ error) {
	return r, ht.ErrNotImplemented
}

// UpdateTask implements UpdateTask operation.
//
// タスクを更新する.
//
// PATCH /projects/{projectID}/tasks/{taskID}
func (UnimplementedHandler) UpdateTask(ctx context.Context, req *UpdateTaskReq, params UpdateTaskParams) (r UpdateTaskRes, _ error) {
	return r, ht.ErrNotImplemented
}
