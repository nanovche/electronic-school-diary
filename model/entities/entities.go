package model

import "time"

// "gorm.io/gorm"

type Administrator struct {
	Administrator_ID uint `gorm:"primaryKey"`
	FirstName  string
	LastName   string
}

type Teacher struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type Parent struct {
	Parent_ID uint
	Email     string
	Password  string
	FirstName string
	LastName  string
}

type Student struct {
	Student_ID uint `gorm:"primaryKey"`
	FullName  string
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
	Inserted_At time.Time
}

type Teacher_Subject struct{
	Teacher_Subject_ID uint `gorm:"primaryKey"`
	Subject_ID uint `gorm:"foreignKey:subject_id"`
	Teacher_ID uint `gorm:"foreignKey:teacher_ID"`
}

type WeekDay struct {
	Weekday_classes_ID uint `gorm:"primaryKey"`
	Consecutive_class int
	Subject_ID uint `gorm:"foreignKey:subject_id"`
	Date_recorded time.Time
}

type Term struct{
	Term_id uint `gorm:"primaryKey"`
	Start_date time.Time
	End_date time.Time
}