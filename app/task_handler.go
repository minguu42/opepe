package app

import (
	"context"
	"time"

	"github.com/minguu42/mtasks/app/logging"
	"github.com/minguu42/mtasks/app/ogen"
)

// CreateTask は POST /projects/{projectID}/tasks に対応するハンドラ
func (h *handler) CreateTask(ctx context.Context, req *ogen.CreateTaskReq, params ogen.CreateTaskParams) (*ogen.Task, error) {
	u, ok := ctx.Value(userKey{}).(*User)
	if !ok {
		logging.Errorf("ctx.Value(userKey{}).(*User) failed")
		return nil, errUnauthorized
	}

	p, err := h.repository.GetProjectByID(ctx, params.ProjectID)
	if err != nil {
		logging.Errorf("repository.GetProjectByID failed: %v", err)
		return nil, errInternalServerError
	}
	if u.ID != p.UserID {
		logging.Errorf("u.ID != p.UserID")
		return nil, errTaskNotFound
	}

	t, err := h.repository.CreateTask(ctx, params.ProjectID, req.Title)
	if err != nil {
		logging.Errorf("repository.Create failed: %v", err)
		return nil, errInternalServerError
	}

	return newTaskResponse(t), nil
}

// ListTasks は GET /projects/{projectID}/tasks に対応するハンドラ
func (h *handler) ListTasks(ctx context.Context, params ogen.ListTasksParams) (*ogen.Tasks, error) {
	u, ok := ctx.Value(userKey{}).(*User)
	if !ok {
		logging.Errorf("ctx.Value(userKey{}).(*User) failed")
		return nil, errUnauthorized
	}

	p, err := h.repository.GetProjectByID(ctx, params.ProjectID)
	if err != nil {
		logging.Errorf("repository.GetProjectByID failed: %v", err)
		return nil, errInternalServerError
	}
	if u.ID != p.UserID {
		logging.Errorf("u.ID != p.UserID")
		return nil, errTaskNotFound
	}

	ts, err := h.repository.GetTasksByProjectID(ctx, p.ID, string(params.Sort.Or(ogen.ListTasksSortMinusCreatedAt)), params.Limit.Or(10)+1, params.Offset.Or(0))
	if err != nil {
		logging.Errorf("repository.GetTasksByProjectID failed: %v", err)
		return nil, errInternalServerError
	}

	hasNext := false
	if len(ts) == params.Limit.Or(10)+1 {
		hasNext = true
		ts = ts[:params.Limit.Or(10)]
	}

	return &ogen.Tasks{
		Tasks:   newTasksResponse(ts),
		HasNext: hasNext,
	}, nil
}

// UpdateTask は PATCH /projects/{projectID}/tasks/{taskID} に対応するハンドラ
func (h *handler) UpdateTask(ctx context.Context, req *ogen.UpdateTaskReq, params ogen.UpdateTaskParams) (*ogen.Task, error) {
	u, ok := ctx.Value(userKey{}).(*User)
	if !ok {
		logging.Errorf("ctx.Value(userKey{}).(*User) failed")
		return nil, errUnauthorized
	}

	p, err := h.repository.GetProjectByID(ctx, params.ProjectID)
	if err != nil {
		logging.Errorf("repository.GetProjectByID failed: %v", err)
		return nil, errInternalServerError
	}
	if u.ID != p.UserID {
		logging.Errorf("u.ID != p.UserID")
		return nil, errProjectNotFound
	}
	t, err := h.repository.GetTaskByID(ctx, params.TaskID)
	if err != nil {
		logging.Errorf("repository.GetTaskByID failed: %v", err)
		return nil, errInternalServerError
	}
	if p.ID != t.ProjectID {
		logging.Errorf("p.ID != t.ProjectID")
		return nil, errTaskNotFound
	}

	if !req.IsCompleted.IsSet() {
		logging.Errorf("value contains nothing")
		return nil, errBadRequest
	}
	now := time.Now()
	if req.IsCompleted.Value {
		t.CompletedAt = &now
	} else {
		t.CompletedAt = nil
	}
	t.UpdatedAt = now
	if err := h.repository.UpdateTask(ctx, params.TaskID, t.CompletedAt, t.UpdatedAt); err != nil {
		logging.Errorf("repository.UpdateTask failed: %v", err)
		return nil, errInternalServerError
	}

	return newTaskResponse(t), nil
}

// DeleteTask は DELETE /projects/{projectID}/tasks/{taskID} に対応するハンドラ
func (h *handler) DeleteTask(ctx context.Context, params ogen.DeleteTaskParams) error {
	u, ok := ctx.Value(userKey{}).(*User)
	if !ok {
		logging.Errorf("ctx.Value(userKey{}).(*User) failed")
		return errUnauthorized
	}

	p, err := h.repository.GetProjectByID(ctx, params.ProjectID)
	if err != nil {
		logging.Errorf("repository.GetProjectByID failed: %v", err)
		return errInternalServerError
	}
	if u.ID != p.UserID {
		logging.Errorf("u.ID != p.UserID")
		return errProjectNotFound
	}
	t, err := h.repository.GetTaskByID(ctx, params.TaskID)
	if err != nil {
		logging.Errorf("repository.GetTaskByID failed: %v", err)
		return errInternalServerError
	}
	if p.ID != t.ProjectID {
		logging.Errorf("p.ID != t.ProjectID")
		return errTaskNotFound
	}

	if err := h.repository.DeleteTask(ctx, t.ID); err != nil {
		logging.Errorf("repository.DeleteTask failed: %v", err)
		return errInternalServerError
	}

	return nil
}
