package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Address  string `json:"address"`
	Role     bool   `json:"role"`
}

func (u *User) TableName() string {
	return "Users"
}
