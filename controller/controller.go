package controller

import (
	"electronic-school-diary/loggerutils"
	"html/template"
	"log"
	"net/http"

	"gorm.io/gorm"
)

type Controller struct {
	Db       *gorm.DB
	ErrorLogger   *log.Logger
	EventLogger  *log.Logger
	Template *template.Template
}

func NewController(db *gorm.DB, loggers []*log.Logger, template *template.Template) *Controller {

	var errorLogger, eventLogger *log.Logger
	for _, logger := range loggers {
		if logger.Prefix() == loggerutils.Error.String() {
			errorLogger = logger
		} else {
			eventLogger = logger
		}
	}

	return &Controller{
		Db:       db,
		ErrorLogger: errorLogger,
		EventLogger: eventLogger,
		Template: template,
	}
}


func (cntrl *Controller) HandleRegister(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("C:\\go-work\\src\\electronic-school-diary\\view\\register.html"))
	t.Execute(w, struct{}{})

}
