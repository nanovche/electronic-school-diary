package main

import (
	c "electronic-school-diary/controller"
	"electronic-school-diary/loggerutils"
	"electronic-school-diary/model"
	"electronic-school-diary/model/dal"
	"electronic-school-diary/service"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"log"
	"net/http"
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

	//instantiating repos
	adminRepo := dal.NewRepositoryAdmin(db)
	authRepo := dal.NewRepositoryAuth(db)
	teacherRepo := dal.NewRepositoryTeacher(db)
	studentRepo := dal.NewStudentRepositoryImpl(db)
	markRepo := dal.NewMarkRepository(db)
	subjectRepo := dal.NewSubjectRepository(db)
	termRepo := dal.NewTermRepository(db)

	//adding repos to repository container
	repo := &dal.RepositoryImpl{}
	repo.SetRepositoryTeacher(teacherRepo)
	repo.SetRepositoryAdmin(adminRepo)
	repo.SetAuthRepository(authRepo)
	repo.SetMarkRepository(markRepo)
	repo.SetStudentRepository(studentRepo)
	repo.SetSubjectRepository(subjectRepo)
	repo.SetTermRepository(termRepo)

	errorLogger := loggerutils.InitLogger(loggerutils.ErrorLoggerFileName, loggerutils.Error.String())
	eventLogger := loggerutils.InitLogger(loggerutils.EventLoggerFileName, loggerutils.Event.String())
	cntrl := c.NewController(repo,
					[]*log.Logger{errorLogger, eventLogger})


	//instantiating service and controller (teacher)
	teacherService := service.NewTeacherService(repo)
	studentService := service.NewStudentService(repo)
	markService := service.NewMarkService(repo)
	teacherController := c.NewTeacherController(teacherService, studentService, markService)

	//add controllers to main controller
	cntrl.SetTeacherController(teacherController)

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

/*	app.mux.HandleFunc("/", app.cntrl.IndexHandler)
	app.mux.HandleFunc("/register", app.cntrl.RegisterHandler)
	app.mux.HandleFunc("/login", app.cntrl.LoginHandler)
	app.mux.HandleFunc("/logout", app.cntrl.LogoutHandler)*/
	app.mux.HandleFunc("/assess-student", app.cntrl.GetTeacherController().AddMarkHandler)
	app.mux.HandleFunc("/update-student-mark", app.cntrl.GetTeacherController().PresentFormMarkHandler).Methods("GET")
	app.mux.HandleFunc("/update-student-mark", app.cntrl.GetTeacherController().ReadFormDataMarkHandler).Methods("POST")
	app.mux.HandleFunc("/update-student-mark", app.cntrl.GetTeacherController().UpdateMarkHandler).Methods("PUT")
	app.mux.HandleFunc("/update-student-mark", app.cntrl.GetTeacherController().DeleteMarkHandler).Methods("DELETE")
	if err := app.server.ListenAndServe(); err != nil {
		app.cntrl.ErrorLogger.Fatalln(err)
	}
}


