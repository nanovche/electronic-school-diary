package model

import (
	"electronic-school-diary/env"
	entities "electronic-school-diary/model/entities"
	"fmt"

	"github.com/go-gorm/mysql"
	"gorm.io/gorm"
)

func NewDB() (*gorm.DB, error) {

	db, err := gorm.Open(mysql.Open("root:23sl_@/electronic_diary?parseTime=true"), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize database session: %s", err)
	}
	return db, nil

}

func Migrate(env *env.Env) {
	env.Db.AutoMigrate(&entities.Mark{})
}
