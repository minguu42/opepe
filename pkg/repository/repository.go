// Package repository はデータベースへの操作を抽象化する
package repository

//go:generate mockgen -source=$GOFILE -destination=../../gen/mock/$GOFILE -package=mock

import (
	"context"

	"github.com/minguu42/opepe/pkg/entity"
)

type Repository interface {
	GetUserByAPIKey(ctx context.Context, apiKey string) (*entity.User, error)

	SaveProject(ctx context.Context, p *entity.Project) error
	GetProjectsByUserID(ctx context.Context, userID string, limit, offset int) ([]*entity.Project, error)
	GetProjectByID(ctx context.Context, id string) (*entity.Project, error)
	DeleteProject(ctx context.Context, id string) error

	SaveTask(ctx context.Context, t *entity.Task) error
	GetTasksByProjectID(ctx context.Context, projectID string, limit, offset int) ([]*entity.Task, error)
	GetTaskByID(ctx context.Context, id string) (*entity.Task, error)
	DeleteTask(ctx context.Context, id string) error
}
