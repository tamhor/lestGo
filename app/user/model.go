package app

import "github.com/google/uuid"

type User struct {
	Uuid     uuid.UUID `json:"uuid"`
	Username string    `gorm:"unique" json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
}

func (User) TableName() string {
	return "users"
}
