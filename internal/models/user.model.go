package models

type User struct {
	User_id  string `db:"id"`
	Email    string `db:"email"`
	Password string `db:"password"`
}
