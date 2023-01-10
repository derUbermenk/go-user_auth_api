package user_repository

import (
	"database/sql"

	"github.com/derUbermenk/go-user_auth_api/repository"
	"github.com/jmoiron/sqlx"
)

type userrepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) (uR *userrepository) {
	return &userrepository{db: db}
}

func (u *userrepository) Create(userInfo map[string]interface{}) (user repository.User, err error) {
	tx := u.db.MustBegin()

	err = tx.QueryRow(`INSERT INTO users (password,email) VALUES ($1, $2) RETURNING id`, userInfo["password"], userInfo["email"]).Scan(&user.ID)

	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}

	return
}

func (u *userrepository) FindByEmail(email string) (user repository.User, err error) {
	err = u.db.Get(&user, `SELECT * FROM users WHERE email=$1`, email)

	// set user password to null
	user.Password = ""

	if err == sql.ErrNoRows {
		err = nil
		return
	}

	return
}

func (u *userrepository) FindPublic(id int) (user repository.User, err error) {
	err = u.db.Get(&user, `SELECT * FROM users WHERE id=$1`, id)

	user.Password = ""

	if err == sql.ErrNoRows {
		err = nil
		return
	}

	return
}
