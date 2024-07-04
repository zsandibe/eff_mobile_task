package repository

import (
	"context"
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
		return entity.Task{}, err
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

	finishedAt := time.Now()

	query := `
        UPDATE task_progress
        SET finished_at = $1, 
            time_difference = $1 - started_at
        WHERE id = $2 AND 
              user_id = $3
			  AND finished_at IS NULL
    `
	_, err = r.db.ExecContext(ctx, query, finishedAt, taskId, userId)
	if err != nil {
		logger.Errorf("error in stopping task: %v", err)
		return err
	}

	return nil
}

func (r *repositoryPostgres) GetTaskProgressByUserId(ctx context.Context, userId int) ([]entity.Task, error) {

	query := `
	SELECT t.id,t.user_id,
	t.name,t.description,
	t.started_at,t.finished_at,
	t.time_difference
	FROM task_progress t
	WHERE t.user_id = $1
	ORDER BY t.time_difference DESC
	`

	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		logger.Errorf("error querying with context: %v", err)
		return nil, err
	}
	defer rows.Close()

	var tasks []entity.Task
	for rows.Next() {
		var task entity.Task

		if err := rows.Scan(
			&task.Id,
			&task.UserId,
			&task.Name,
			&task.Description,
			&task.StartedAt,
			&task.FinishedAt,
			&task.TimeDifference,
		); err != nil {
			logger.Errorf("problems with scanning rows: %v", err)
			return tasks, err
		}

		tasks = append(tasks, task)
	}

	if err := rows.Close(); err != nil {
		logger.Error(err)
		return tasks, err
	}

	if err := rows.Err(); err != nil {
		logger.Error(err)
		return tasks, err
	}

	return tasks, nil
}

func (r *repositoryPostgres) IsTaskExists(ctx context.Context, taskId int) (bool, error) {
	var exists bool

	query := `
    SELECT EXISTS (
        SELECT 1
        FROM task_progress
        WHERE id = $1
    )
    `

	err := r.db.QueryRowContext(ctx, query, taskId).Scan(&exists)
	if err != nil {
		logger.Errorf("error checking if task exists: %v", err)
		return false, err
	}

	return exists, nil
}
