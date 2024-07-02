package domain

import "fmt"

var (
	ErrCreatingUser = fmt.Errorf("user already exists")
	ErrNoUser       = fmt.Errorf("user not found")
)
