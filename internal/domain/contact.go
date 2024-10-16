package domain

type Contact struct {
	Id            int    `json:"id" db:"id"`
	Email         string `json:"email" db:"email"`
	ExtraContacts string `json:"extraContacts" db:"extra_contacts"`
}
