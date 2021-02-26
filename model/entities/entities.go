package model

import "gorm.io/gorm"

// "gorm.io/gorm"

type Teacher struct {
	gorm.Model
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type Student struct {
	Student_ID uint `gorm:"primaryKey"`
	FirstName  string
	LastName   string
	Number     uint8
}

type Subject struct {
	Subject_ID uint `gorm:"primaryKey"`
	Title      string
}

type Mark struct {
	Mark_ID    uint `gorm:"primaryKey"`
	Student_ID uint `gorm:"foreignKey:student_ID"`
	Teacher_ID uint `gorm:"foreignKey:teacher_ID"`
	Subject_ID uint `gorm:"foreignKey:subject_id"`
	Value      int
}
