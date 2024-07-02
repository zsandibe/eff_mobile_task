package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/zsandibe/eff_mobile_task/internal/domain"
	"github.com/zsandibe/eff_mobile_task/internal/entity"
)

type Repository interface {
	AddUser(ctx context.Context, inp domain.GetUserResponse) (entity.User, error)
	GetUserById(ctx context.Context, id int) (entity.User, error)
	UpdateUserData(ctx context.Context, userId int, params domain.UserDataUpdatingRequest) error
	DeleteUserById(ctx context.Context, userId int) error
	StartTask(ctx context.Context, inp domain.CreateTaskRequest) (entity.Task, error)
	StopTask(ctx context.Context, taskId int, id string) error
}

type repositoryPostgres struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repositoryPostgres {
	return &repositoryPostgres{
		db: db,
	}
}
