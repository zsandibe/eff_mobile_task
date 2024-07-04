package service

import (
	"context"
	"database/sql"
	"errors"

	"github.com/zsandibe/eff_mobile_task/internal/domain"
	"github.com/zsandibe/eff_mobile_task/internal/entity"
	logger "github.com/zsandibe/eff_mobile_task/pkg"
)

func (s *service) AddUser(ctx context.Context, inp *domain.GetUserRequest) (entity.User, error) {
	logger.Debugf("Add user: %+v ", inp)
	checkedUser, err := getInfoFromApi(inp.PassportSerie, inp.PassportNumber, s.conf)
	if err != nil {
		return entity.User{}, domain.ErrUserNotFound
	}

	exists, _ := s.repo.CheckUserByPassport(ctx, inp.PassportSerie, inp.PassportNumber)
	if exists {
		return entity.User{}, domain.ErrUserExists
	}

	resp := domain.GetUserResponse{
		PassportSerie:  inp.PassportSerie,
		PassportNumber: inp.PassportNumber,
		People: domain.People{
			Name:       checkedUser.Name,
			Surname:    checkedUser.Surname,
			Patronymic: checkedUser.Patronymic,
			Address:    checkedUser.Address,
		},
	}

	user, err := s.repo.AddUser(ctx, resp)
	if err != nil {
		return entity.User{}, domain.ErrCreatingUser
	}
	logger.Debug(user)
	logger.Info("User successfully created")
	return user, nil
}

func (s *service) GetUserById(ctx context.Context, id int) (entity.User, error) {
	logger.Debugf("Get user by id: %d", id)
	user, err := s.repo.GetUserById(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return entity.User{}, domain.ErrUserNotFound
		}
		return entity.User{}, err
	}
	logger.Debug(user)
	logger.Info("User successfully got")
	return user, nil
}

func (s *service) UpdateUserData(ctx context.Context, userId int, params domain.UserDataUpdatingRequest) error {
	logger.Debugf("Update user data:  %d , %+v ", userId, params)
	_, err := s.repo.GetUserById(ctx, userId)
	if err != nil {
		return domain.ErrUserNotFound
	}

	if err := s.repo.UpdateUserData(ctx, userId, params); err != nil {
		return err
	}
	logger.Info("User`s data was successfully updated")
	return nil
}

func (s *service) DeleteUserById(ctx context.Context, userId int) error {
	logger.Debugf("Delete user by id: %d ", userId)
	if err := s.repo.DeleteUserById(ctx, userId); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.ErrUserNotFound
		}
		return err
	}
	logger.Info("User was successfully updated")
	return nil
}

func (s *service) GetUsersList(ctx context.Context, params domain.UsersListParams) ([]entity.User, error) {
	logger.Debugf("Get users list: %v ", params)
	return s.repo.GetUsersList(ctx, params)
}
