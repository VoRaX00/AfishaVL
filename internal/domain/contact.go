package domain

type Contact struct {
	Id            int    `json:"id" db:"id"`
	Email         string `json:"email" db:"email"`
	ExtraContacts string `json:"extra_contacts" db:"extra_contacts"`
}
