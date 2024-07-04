package repository

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/zsandibe/eff_mobile_task/internal/domain"
	"github.com/zsandibe/eff_mobile_task/internal/entity"
	logger "github.com/zsandibe/eff_mobile_task/pkg"
)

func (r *repositoryPostgres) GetUsersList(ctx context.Context, params domain.UsersListParams) ([]entity.User, error) {
	users := make([]entity.User, 0)
	var (
		args    []interface{}
		where   []string
		orderBy []string
	)

	if params.PassportSerie != "" {
		where = append(where, "passport_serie = $"+strconv.Itoa(len(args)+1))
		args = append(args, params.PassportSerie)
	}
	if params.PassportNumber != "" {
		where = append(where, "passport_number = $"+strconv.Itoa(len(args)+1))
		args = append(args, params.PassportNumber)
	}
	if params.Name != "" {
		where = append(where, "name = $"+strconv.Itoa(len(args)+1))
		args = append(args, params.Name)
	}
	if params.Surname != "" {
		where = append(where, "surname = $"+strconv.Itoa(len(args)+1))
		args = append(args, params.Surname)
	}
	if params.Patronymic != "" {
		where = append(where, "patronymic = $"+strconv.Itoa(len(args)+1))
		args = append(args, params.Patronymic)
	}
	if params.Address != "" {
		where = append(where, "address = $"+strconv.Itoa(len(args)+1))
		args = append(args, params.Address)
	}

	query := "SELECT * FROM users"
	if len(where) > 0 {
		query += " WHERE " + strings.Join(where, " AND ")
	}
	if len(orderBy) > 0 {
		query += " ORDER BY " + strings.Join(orderBy, ", ")
	}
	if params.Limit > 0 {
		query += " LIMIT $" + strconv.Itoa(len(args)+1)
		args = append(args, params.Limit)
	}
	if params.Offset > 0 {
		query += " OFFSET $" + strconv.Itoa(len(args)+1)
		args = append(args, params.Offset)
	}

	fmt.Println(query)
	fmt.Println(args)

	rows, err := r.db.QueryContext(ctx, query, args...)
	if err != nil {
		logger.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user entity.User
		if err := rows.Scan(&user.Id, &user.PassportSerie, &user.PassportNumber, &user.Name, &user.Surname,
			&user.Patronymic, &user.Address); err != nil {
			logger.Error(err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		logger.Error(err)
		return nil, err
	}

	return users, nil
}

func (r *repositoryPostgres) AddUser(ctx context.Context, inp domain.GetUserResponse) (entity.User, error) {
	var id int
	query := `
	INSERT INTO users (passport_serie, passport_number, name, surname, patronymic, address)
	VALUES ($1, $2, $3, $4, $5, $6)
	RETURNING id
	`
	var user entity.User

	err := r.db.QueryRowContext(ctx, query,
		inp.PassportSerie,
		inp.PassportNumber,
		inp.People.Name,
		inp.People.Surname,
		inp.People.Patronymic,
		inp.People.Address).Scan(&id)
	if err != nil {
		logger.Errorf("Error in inserting user: %v", err)
		return entity.User{}, domain.ErrCreatingUser
	}

	user = entity.User{
		Id:             id,
		PassportSerie:  inp.PassportSerie,
		PassportNumber: inp.PassportNumber,
		Name:           inp.People.Name,
		Surname:        inp.People.Surname,
		Patronymic:     inp.People.Patronymic,
		Address:        inp.People.Address,
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
		logger.Error(ctx, fmt.Errorf("error with executing query: %v", err))
		return err
	}

	return nil
}

func (r *repositoryPostgres) CheckUserByPassport(ctx context.Context, passportSerie, passportNumber string) (bool, error) {
	var exists bool

	query := `
    SELECT EXISTS (
        SELECT 1
        FROM users
        WHERE passport_serie = $1
		AND passport_number = $2
    )
    `

	err := r.db.QueryRowContext(ctx, query, passportSerie, passportNumber).Scan(&exists)
	if err != nil {
		logger.Errorf("error checking if task exists: %v", err)
		return false, err
	}

	return exists, nil
}
