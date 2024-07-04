package service

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/zsandibe/eff_mobile_task/internal/domain"
	"github.com/zsandibe/eff_mobile_task/internal/entity"
	logger "github.com/zsandibe/eff_mobile_task/pkg"
)

func (s *service) StartTask(ctx context.Context, inp domain.CreateTaskRequest) (entity.Task, error) {
	logger.Debugf("Start task: input %+v ", inp)
	userId, err := strconv.Atoi(inp.UserId)
	if err != nil {
		return entity.Task{}, err
	}
	_, err = s.repo.GetUserById(ctx, userId)
	if err != nil {
		return entity.Task{}, domain.ErrUserNotFound
	}

	task, err := s.repo.StartTask(ctx, inp)
	if err != nil {
		return entity.Task{}, domain.ErrCreatingTask
	}
	logger.Debugf("Task: %+v ", task)
	logger.Info("Successfully started task")
	return task, nil
}

func (s *service) StopTask(ctx context.Context, taskId int, id string) error {
	logger.Debugf("Stop task: taskId %d ,  userId %s ", taskId, id)
	userId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	_, err = s.repo.GetUserById(ctx, userId)
	if err != nil {
		return domain.ErrUserNotFound
	}
	exists, _ := s.repo.IsTaskExists(ctx, taskId)
	if !exists {
		return domain.ErrTaskNotFound
	}
	logger.Info("Task successfully stopped")
	return s.repo.StopTask(ctx, taskId, id)
}

func (s *service) GetTaskProgressByUserId(ctx context.Context, userId int) ([]entity.Task, error) {
	logger.Debugf("Get task progress by userId: userId  %d ", userId)
	_, err := s.repo.GetUserById(ctx, userId)
	if err != nil {
		return []entity.Task{}, domain.ErrUserNotFound
	}

	tasks, err := s.repo.GetTaskProgressByUserId(ctx, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []entity.Task{}, domain.ErrTaskNotFound
		}
		return []entity.Task{}, err
	}

	logger.Info("Task progress successfully got")
	return tasks, nil
}
