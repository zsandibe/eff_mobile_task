package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/zsandibe/eff_mobile_task/internal/domain"
	"github.com/zsandibe/eff_mobile_task/internal/entity"
	logger "github.com/zsandibe/eff_mobile_task/pkg"
)

func (r *repositoryPostgres) AddUser(ctx context.Context, inp domain.GetUserResponse) (entity.User, error) {
	var id int
	query := `
	INSERT INTO users (passport_serie, passport_number, name, surname, patronymic, address)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
	`
	var user entity.User

	err := r.db.QueryRowContext(ctx, query,
		inp.PassportCredentials.PassportSerie,
		inp.PassportCredentials.PassportNumber,
		inp.Name,
		inp.Surname,
		inp.Patronymic,
		inp.Address).Scan(&id)
	if err != nil {
		logger.Errorf("Error in inserting user: %v", err)
		return entity.User{}, domain.ErrCreatingUser
	}

	user = entity.User{
		Id:             id,
		PassportSerie:  inp.PassportCredentials.PassportSerie,
		PassportNumber: inp.PassportCredentials.PassportNumber,
		Name:           inp.Name,
		Surname:        inp.Surname,
		Patronymic:     inp.Patronymic,
		Address:        inp.Address,
	}

	return user, nil
}

func (r *repositoryPostgres) GetUserById(ctx context.Context, id int) (entity.User, error) {
	var user entity.User

	query := `
		SELECT users.id,users.passport_serie,users.passport_number,users.name,users.surname,users.patronymic,users.address
		FROM users WHERE id = $1
	`

	if err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.Id,
		&user.PassportSerie,
		&user.PassportNumber,
		&user.Name,
		&user.Surname,
		&user.Patronymic,
		&user.Address,
	); err != nil {
		return user, err
	}

	return user, nil
}

func (r *repositoryPostgres) UpdateUserData(ctx context.Context, userId int, params domain.UserDataUpdatingRequest) error {
	query := `
	UPDATE users SET
	passport_serie = COALESCE(NULLIF($1, ''), passport_serie),
	passport_number = COALESCE(NULLIF($2, ''), passport_number),
	name = COALESCE(NULLIF($3, ''), name),
	surname = COALESCE(NULLIF($4, ''), surname),
	patronymic = COALESCE(NULLIF($5, ''), patronymic),
	address = COALESCE(NULLIF($6, ''), address)
	WHERE id = $7
	`

	_, err := r.db.ExecContext(ctx, query,
		params.PassportSerie, params.PassportNumber, params.Name, params.Surname, params.Patronymic, params.Address, userId)
	if err != nil {
		logger.Error(ctx, fmt.Errorf("error with executing query: %v", err))
		return err
	}

	return nil
}

func (r *repositoryPostgres) DeleteUserById(ctx context.Context, userId int) error {
	query := `
		DELETE FROM users WHERE id = $1
	`

	_, err := r.db.ExecContext(ctx, query, userId)
	if err != nil {
		if err == sql.ErrNoRows {
			return domain.ErrNoUser
		}
		logger.Error(ctx, fmt.Errorf("error with executing query: %v", err))
		return err
	}

	return nil
}
