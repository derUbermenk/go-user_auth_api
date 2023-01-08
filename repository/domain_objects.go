package repository

type User struct {
	ID       int    `db:"id"`
	Email    string `db:"email"`
	Password string `string:"string"`
}
