package main

import (
	c "electronic-school-diary/controller"
	"electronic-school-diary/model"
	"html/template"
	"log"
	"net/http"
	"electronic-school-diary/loggerutils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type application struct{
	server *http.Server
	mux *mux.Router
	cntrl *c.Controller
}

func main() {

	var db *gorm.DB
	var err error
	if db, err = model.NewDB(); err != nil {
		log.Fatal(err)
	}

	errorLogger := loggerutils.InitLogger(loggerutils.ErrorLoggerFileName, loggerutils.Error.String())
	eventLogger := loggerutils.InitLogger(loggerutils.EventLoggerFileName, loggerutils.Event.String())
	cntrl := c.NewController(db,
					[]*log.Logger{errorLogger, eventLogger},
					&template.Template{})


	handler := mux.NewRouter()
	server := &http.Server{
		Addr: ":8080",
		Handler:  handler,
	}

	application := application{
		server: server,
		mux: handler,
		cntrl: cntrl,
	}

	application.launch()

}

func(app *application) launch() {

	app.mux.HandleFunc("/register", app.cntrl.HandleRegister )
	if err := app.server.ListenAndServe(); err != nil {
		app.cntrl.ErrorLogger.Fatalln(err)
	}
}


