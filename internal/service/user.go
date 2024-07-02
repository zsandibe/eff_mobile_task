package service

import (
	"context"
	"fmt"

	"github.com/zsandibe/eff_mobile_task/internal/domain"
	"github.com/zsandibe/eff_mobile_task/internal/entity"
)

func (s *service) AddUser(ctx context.Context, inp *domain.GetUserRequest) (entity.User, error) {
	checkedUser, err := getInfoByPassport(inp.PassportSerie, inp.PassportNumber)
	if err != nil {
		return entity.User{}, err
	}
	fmt.Println("REPO", checkedUser.PassportCredentials.PassportSerie, checkedUser.PassportCredentials.PassportNumber)
	return s.repo.AddUser(ctx, checkedUser)
}

func (s *service) GetUserById(ctx context.Context, id int) (entity.User, error) {
	return s.repo.GetUserById(ctx, id)
}

func (s *service) UpdateUserData(ctx context.Context, userId int, params domain.UserDataUpdatingRequest) error {
	return s.repo.UpdateUserData(ctx, userId, params)
}

func (s *service) DeleteUserById(ctx context.Context, userId int) error {
	return s.repo.DeleteUserById(ctx, userId)
}
