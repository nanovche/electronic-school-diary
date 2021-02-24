package main

import (
	"electronic-school-diary/env"
	"electronic-school-diary/model"
	"log"

	"gorm.io/gorm"
)

func main() {

	var db *gorm.DB
	var err error
	if db, err = model.NewDB(); err != nil {
		log.Fatal(err)
	}

	env := env.Env{
		Db: db,
	}

	model.Migrate(&env)

}
