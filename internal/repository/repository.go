package repository

import "github.com/jmoiron/sqlx"

type Repository interface {
}

type repositoryPostgres struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *repositoryPostgres {
	return &repositoryPostgres{
		db: db,
	}
}
