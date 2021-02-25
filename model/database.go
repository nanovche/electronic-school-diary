package model

import (
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

