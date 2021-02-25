package model

import (
	"gorm.io/gorm"
)

type Teacher struct {
	gorm.Model
	Email     string
	Password  string
	FirstName string
	LastName  string
}
