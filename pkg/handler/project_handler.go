package handler

import (
	"context"
	"errors"

	"github.com/minguu42/opepe/gen/ogen"
	"github.com/minguu42/opepe/pkg/domain/model"
	"github.com/minguu42/opepe/pkg/domain/repository"
	"github.com/minguu42/opepe/pkg/logging"
	"github.com/minguu42/opepe/pkg/ttime"
)

// CreateProject は POST /projects に対応するハンドラ
func (h *Handler) CreateProject(ctx context.Context, req *ogen.CreateProjectReq) (*ogen.Project, error) {
	u, ok := ctx.Value(userKey{}).(*model.User)
	if !ok {
		return nil, errUnauthorized
	}

	now := ttime.Now(ctx)
	p := model.Project{
		ID:         h.IDGenerator.Generate(),
		UserID:     u.ID,
		Name:       req.Name,
		Color:      req.Color,
		IsArchived: false,
		UpdatedAt:  now,
		CreatedAt:  now,
	}
	if err := h.Repository.CreateProject(ctx, &p); err != nil {
		return nil, errInternalServerError
	}

	return newProjectResponse(&p), nil
}

// ListProjects は GET /projects に対応するハンドラ
func (h *Handler) ListProjects(ctx context.Context, params ogen.ListProjectsParams) (*ogen.Projects, error) {
	u, ok := ctx.Value(userKey{}).(*model.User)
	if !ok {
		return nil, errUnauthorized
	}

	limit := params.Limit.Or(defaultLimit)
	ps, err := h.Repository.GetProjectsByUserID(ctx, u.ID, limit+1, params.Offset.Or(defaultOffset))
	if err != nil {
		logging.Errorf(ctx, "repository.GetProjectsByUserID failed: %s", err)
		return nil, errInternalServerError
	}

	hasNext := false
	if len(ps) == limit+1 {
		hasNext = true
		ps = ps[:limit]
	}

	return &ogen.Projects{
		Projects: newProjectsResponse(ps),
		HasNext:  hasNext,
	}, nil
}

// UpdateProject は PATCH /projects/{projectID} に対応するハンドラ
func (h *Handler) UpdateProject(ctx context.Context, req *ogen.UpdateProjectReq, params ogen.UpdateProjectParams) (*ogen.Project, error) {
	u, ok := ctx.Value(userKey{}).(*model.User)
	if !ok {
		return nil, errUnauthorized
	}

	p, err := h.Repository.GetProjectByID(ctx, params.ProjectID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return nil, errProjectNotFound
		}
		logging.Errorf(ctx, "repository.GetProjectByID failed: %s", err)
		return nil, errInternalServerError
	}

	if !u.HasProject(p) {
		logging.Errorf(ctx, "user does not have the project")
		return nil, errProjectNotFound
	}

	newProject := model.Project{
		ID:         p.ID,
		UserID:     p.UserID,
		Name:       req.Name.Or(p.Name),
		Color:      req.Color.Or(p.Color),
		IsArchived: req.IsArchived.Or(p.IsArchived),
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  ttime.Now(ctx),
	}
	if err := h.Repository.UpdateProject(ctx, &newProject); err != nil {
		logging.Errorf(ctx, "repository.SaveProject failed: %s", err)
		return nil, errInternalServerError
	}

	return newProjectResponse(&newProject), nil
}

// DeleteProject は DELETE /projects/{projectID} に対応するハンドラ
func (h *Handler) DeleteProject(ctx context.Context, params ogen.DeleteProjectParams) error {
	u, ok := ctx.Value(userKey{}).(*model.User)
	if !ok {
		return errUnauthorized
	}

	p, err := h.Repository.GetProjectByID(ctx, params.ProjectID)
	if err != nil {
		if errors.Is(err, repository.ErrModelNotFound) {
			return errProjectNotFound
		}
		logging.Errorf(ctx, "repository.GetProjectByID failed: %s", err)
		return errInternalServerError
	}

	if !u.HasProject(p) {
		logging.Errorf(ctx, "user does not have the project")
		return errProjectNotFound
	}

	if err := h.Repository.DeleteProject(ctx, p.ID); err != nil {
		logging.Errorf(ctx, "repository.DeleteProject failed: %s", err)
		return errInternalServerError
	}

	return nil
}
