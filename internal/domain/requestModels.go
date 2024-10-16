package domain

type UserToLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
