package models

type User struct {
	ID         int    `db:"id"`
	FirstName  string `db:"first_name"`
	LastName   string `db:"last_name"`
	Username   string `db:"username"`
	IsVerified bool   `db:"is_verified"`
}
