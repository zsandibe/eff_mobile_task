package service

import (
	"context"

	"github.com/zsandibe/eff_mobile_task/config"
	"github.com/zsandibe/eff_mobile_task/internal/domain"
	"github.com/zsandibe/eff_mobile_task/internal/entity"
	"github.com/zsandibe/eff_mobile_task/internal/repository"
)

type Service interface {
	AddUser(ctx context.Context, inp *domain.GetUserRequest) (entity.User, error)
	GetUserById(ctx context.Context, id int) (entity.User, error)
	UpdateUserData(ctx context.Context, userId int, params domain.UserDataUpdatingRequest) error
	DeleteUserById(ctx context.Context, userId int) error
	StartTask(ctx context.Context, inp domain.CreateTaskRequest) (entity.Task, error)
	StopTask(ctx context.Context, taskId int, id string) error
	GetTaskProgressByUserId(ctx context.Context, userId int) ([]entity.Task, error)
	GetUsersList(ctx context.Context, params domain.UsersListParams) ([]entity.User, error)
}

type service struct {
	repo repository.Repository
	conf config.Config
}

func NewService(repo repository.Repository) *service {
	return &service{
		repo: repo,
	}
}
