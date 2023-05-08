package models

import "gorm.io/gorm"

type UserModel struct {
	gorm.Model
	Name     string
	Email    string
	Password string
}
