package models

type User struct {
	UserID   int64  `db:"user_id" json:"user_id"`
	Username string `db:"username" json:"user_name"`
	Password string `db:"password" json:"-"`
	Token    string `db:"-" json:"token"`
}
