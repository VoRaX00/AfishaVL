package domain

type User struct {
	Id       int    `json:"id" db:"id"`
	Login    string `json:"login" db:"login"`
	Password string `json:"password" db:"password_hash"`
}
