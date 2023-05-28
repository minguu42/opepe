// Code generated by ogen, DO NOT EDIT.

package ogen

import (
	"context"
)

// Handler handles operations described by OpenAPI v3 specification.
type Handler interface {
	// CreateTasks implements CreateTasks operation.
	//
	// 新しいタスクを作成する.
	//
	// POST /tasks
	CreateTasks(ctx context.Context, req *CreateTasksReq) (CreateTasksRes, error)
	// DeleteTask implements deleteTask operation.
	//
	// タスクを削除する.
	//
	// DELETE /tasks/{taskID}
	DeleteTask(ctx context.Context, params DeleteTaskParams) (DeleteTaskRes, error)
	// GetHealth implements getHealth operation.
	//
	// サーバの状態を取得する.
	//
	// GET /health
	GetHealth(ctx context.Context) (GetHealthRes, error)
	// GetTasks implements getTasks operation.
	//
	// 作成日時の降順で最大25件まで取得する。.
	//
	// GET /tasks
	GetTasks(ctx context.Context) (GetTasksRes, error)
	// PatchTask implements patchTask operation.
	//
	// タスクを更新する.
	//
	// PATCH /tasks/{taskID}
	PatchTask(ctx context.Context, req *PatchTaskReq, params PatchTaskParams) (PatchTaskRes, error)
}

// Server implements http server based on OpenAPI v3 specification and
// calls Handler to handle requests.
type Server struct {
	h Handler
	baseServer
}

// NewServer creates new Server.
func NewServer(h Handler, opts ...ServerOption) (*Server, error) {
	s, err := newServerConfig(opts...).baseServer()
	if err != nil {
		return nil, err
	}
	return &Server{
		h:          h,
		baseServer: s,
	}, nil
}