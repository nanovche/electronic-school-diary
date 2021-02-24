package env

import (
	"html/template"
	"log"

	"gorm.io/gorm"
)

type Env struct {
	Db        *gorm.DB
	Logger    *log.Logger
	Templates *template.Template
}

// func (env *Env) InitDB() error {
// 	var err error
// 	env.db, err = sql.Open("mysql", "root:23sl_@/school?parseTime=true")
// 	return err
// }
