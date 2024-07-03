package entity

import "time"

type Task struct {
	Id             int       `json:"id" db:"id"`
	UserId         int       `json:"user_id" db:"user_id"`
	Name           string    `json:"name" db:"name"`
	Description    string    `json:"description,omitempty" db:"description"`
	StartedAt      time.Time `json:"started_at,omitempty" db:"started_at"`
	FinishedAt     time.Time `json:"finished_at,omitempty" db:"finished_at"`
	TimeDifference string    `json:"time_difference,omitempty" db:"time_difference"`
}
