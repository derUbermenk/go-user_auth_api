package user_repository

import "github.com/jmoiron/sqlx"

type userrepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) (uR *userrepository) {
	return &userrepository{db: db}
}
