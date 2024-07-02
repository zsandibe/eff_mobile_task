package service

import (
	"context"

	"github.com/zsandibe/eff_mobile_task/internal/domain"
	"github.com/zsandibe/eff_mobile_task/internal/entity"
)

func (s *service) StartTask(ctx context.Context, inp domain.CreateTaskRequest) (entity.Task, error) {
	return s.repo.StartTask(ctx, inp)
}

func (s *service) StopTask(ctx context.Context, taskId int, id string) error {
	return s.repo.StopTask(ctx, taskId, id)
}
