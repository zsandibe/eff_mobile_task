package domain

import "fmt"

var (
	ErrCreatingUser      = fmt.Errorf("problems with creating user")
	ErrUserNotFound      = fmt.Errorf("user not found")
	ErrTaskNotFound      = fmt.Errorf("task not found")
	ErrUserExists        = fmt.Errorf("user already exists")
	ErrTaskHasNotStarted = fmt.Errorf("task hasn`t started")
	ErrCreatingTask      = fmt.Errorf("problems with creating task")
)
