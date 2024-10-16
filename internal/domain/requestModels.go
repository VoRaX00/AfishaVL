package domain

type UserToLogin struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type UserRegister struct {
	Phone    string
	Password string
}
