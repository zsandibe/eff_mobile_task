package repository

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/zsandibe/eff_mobile_task/internal/domain"
	"github.com/zsandibe/eff_mobile_task/internal/entity"
	logger "github.com/zsandibe/eff_mobile_task/pkg"
)

func (r *repositoryPostgres) StartTask(ctx context.Context, inp domain.CreateTaskRequest) (entity.Task, error) {
	var id int
	userId, err := strconv.Atoi(inp.UserId)
	if err != nil {
		return entity.Task{}, err
	}
	query := `
		INSERT INTO task_progress (user_id, name, description,started_at)
		VALUES ($1, $2, $3,$4)
		RETURNING id
		`
	var task entity.Task

	startTime := time.Now()

	err = r.db.QueryRowContext(ctx, query,
		userId,
		inp.Name,
		inp.Description,
		startTime).Scan(&id)
	if err != nil {
		logger.Errorf("Error in inserting user: %v", err)
		return entity.Task{}, domain.ErrCreatingUser
	}

	task = entity.Task{
		Id:          id,
		UserId:      userId,
		Name:        inp.Name,
		Description: inp.Description,
		StartedAt:   time.Now(),
	}

	return task, nil

}

func (r *repositoryPostgres) StopTask(ctx context.Context, taskId int, id string) error {
	userId, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	// Получаем время начала задачи
	var startedAt time.Time
	err = r.db.QueryRowContext(ctx, "SELECT started_at FROM task_progress WHERE id = $1 AND user_id = $2", taskId, userId).Scan(&startedAt)
	if err != nil {
		logger.Errorf("error fetching start time: %v", err)
		return err
	}

	if startedAt.IsZero() {
		logger.Errorf("start time is zero for task %d and user %d", taskId, userId)
		return fmt.Errorf("start time is not set for task %d and user %d", taskId, userId)
	}

	// Вычисляем разницу времени
	finishedAt := time.Now()
	duration := finishedAt.Sub(startedAt)
	fmt.Println(duration)
	if duration < 0 {
		logger.Errorf("negative duration calculated for task %d and user %d", taskId, userId)
		return fmt.Errorf("negative duration for task %d and user %d", taskId, userId)
	}

	// Обновляем запись с новым временем завершения и разницей времени
	query := `
        UPDATE task_progress
        SET finished_at = $1, 
            time_difference = $2
        WHERE id = $3 AND 
              user_id = $4
    `
	_, err = r.db.ExecContext(ctx, query, finishedAt, duration, taskId, userId)
	if err != nil {
		logger.Errorf("error in stopping task: %v", err)
		return err
	}

	return nil
}
