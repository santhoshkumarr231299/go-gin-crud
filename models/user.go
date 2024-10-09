package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"primary_key"`
	Email string `json:"email"`
}
