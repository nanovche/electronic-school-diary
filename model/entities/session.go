package model

type Session struct{
	Session_id uint `gorm:"primaryKey"`
	Email string
	Session_Token string
}
