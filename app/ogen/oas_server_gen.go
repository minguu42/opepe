// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateProject implements CreateProject operation.
	//
	// 新しいプロジェクトを作成する.
	//
	// POST /projects
	CreateProject(ctx context.Context, req *CreateProjectReq) (*Project, error)
	// CreateTask implements CreateTask operation.
	//
	// 新しいタスクを作成する.
	//
	// POST /projects/{projectID}/tasks
	CreateTask(ctx context.Context, req *CreateTaskReq, params CreateTaskParams) (*Task, error)
	// DeleteProject implements DeleteProject operation.
	//
	// プロジェクトを削除する.
	//
	// DELETE /projects/{projectID}
	DeleteProject(ctx context.Context, params DeleteProjectParams) error
	// DeleteTask implements DeleteTask operation.
	//
	// タスクを削除する.
	//
	// DELETE /projects/{projectID}/tasks/{taskID}
	DeleteTask(ctx context.Context, params DeleteTaskParams) error
	// GetHealth implements GetHealth operation.
	//
	// サーバの状態を取得する.
	//
	// GET /health
	GetHealth(ctx context.Context) (*GetHealthOK, error)
	// ListProjects implements ListProjects operation.
	//
	// 作成日時の降順で取得する。.
	//
	// GET /projects
	ListProjects(ctx context.Context, params ListProjectsParams) (*Projects, error)
	// ListTasks implements ListTasks operation.
	//
	// 作成日時の降順で取得する。.
	//
	// GET /projects/{projectID}/tasks
	ListTasks(ctx context.Context, params ListTasksParams) (*Tasks, error)
	// UpdateProject implements UpdateProject operation.
	//
	// プロジェクトを更新する.
	//
	// PATCH /projects/{projectID}
	UpdateProject(ctx context.Context, req *UpdateProjectReq, params UpdateProjectParams) (*Project, error)
	// UpdateTask implements UpdateTask operation.
	//
	// タスクを更新する.
	//
	// PATCH /projects/{projectID}/tasks/{taskID}
	UpdateTask(ctx context.Context, req *UpdateTaskReq, params UpdateTaskParams) (*Task, error)
	// NewError creates *ErrorStatusCode from error returned by handler.
	//
	// Used for common default response.
	NewError(ctx context.Context, err error) *ErrorStatusCode
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h   Handler
	sec SecurityHandler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, sec SecurityHandler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		sec:        sec,
		baseServer: s,
	}, nil
}
