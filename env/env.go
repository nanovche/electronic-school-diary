package env

import (
	"database/sql"
	"html/template"
	"log"
)

type Env struct {
	db        *sql.DB
	logger    *log.Logger
	templates *template.Template
}

func (env *Env) InitDB() error {
	var err error
	env.db, err = sql.Open("mysql", "root:23sl_@/school?parseTime=true")
	return err
}
