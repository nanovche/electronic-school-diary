package controller

import (
	"electronic-school-diary/loggerutils"
	"electronic-school-diary/model/dal"
	"github.com/google/uuid"
	"log"
	"net/http"
)

type Loggers struct {
	ErrorLogger   *log.Logger
	EventLogger  *log.Logger
}

type Controller struct {
	Loggers
	repo dal.IRepository
	teacherController TeacherController
}

func NewController(repo dal.IRepository, loggers []*log.Logger) *Controller {

	var errorLogger, eventLogger *log.Logger
	for _, logger := range loggers {
		if logger.Prefix() == loggerutils.Error.String() {
			errorLogger = logger
		} else {
			eventLogger = logger
		}
	}

	lgs := Loggers{errorLogger, eventLogger}

	return &Controller{Loggers : lgs,repo : repo}
}

func(cntrl *Controller) GetTeacherController() TeacherController{
	return cntrl.teacherController
}

func(cntrl *Controller) SetTeacherController(teacherController TeacherController) {
	cntrl.teacherController = teacherController
}


/*func (cntrl *Controller) RegisterHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

		case "GET": {
			t := template.Must(template.ParseFiles("C:\\go-work\\src\\electronic-school-diary\\view\\register.html"))
			if err := t.Execute(w, struct{}{}); err != nil {
				cntrl.ErrorLogger.Println("failed to render template: %s ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
		}
		case "POST": {
			defer http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			if err := r.ParseForm(); err != nil {
				cntrl.ErrorLogger.Println("failed to parse form: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			teacher := model.Teacher{
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Email:     r.Form.Get("email"),
				Password:  r.Form.Get("psw"),
				FirstName: r.Form.Get("firstname"),
				LastName:  r.Form.Get("lastname"),
			}

			rawPass := teacher.Password
			var hashedPassword string
			var err error
			if hashedPassword, err = pwd.EncryptPassword(rawPass); err != nil {
				cntrl.ErrorLogger.Println(err)
				if hashedPassword, err = pwd.EncryptPassword(rawPass) ; err != nil {
					cntrl.ErrorLogger.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
					return
				}
			}
			teacher.Password = hashedPassword

			repo := cntrl.Repository.GetAdminRepository()
			fmt.Println(repo)
			if err := cntrl.Repository.GetAdminRepository().CreateTeacher(&teacher); err != nil {
				cntrl.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			cntrl.EventLogger.Println("Teacher %s", teacher.FirstName, " created successfully.")
			w.WriteHeader(http.StatusCreated)
			return
		}
	}
}*/

/*func (cntrl *Controller) IndexHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "POST": {
		cookie, _ := r.Cookie("SessionID")
		_, err := cntrl.Repository.GetAuthRepository().RetrieveEmail(cookie.Value)
		if err != nil {
			cntrl.ErrorLogger.Println(err)
		}
		t := template.Must(template.ParseFiles("C:\\go-work\\src\\electronic-school-diary\\view\\index.html"))
		if err := t.Execute(w, struct{}{}); err != nil {
			cntrl.ErrorLogger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		return
	}

	}
}*/

/*func (cntrl *Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

		case "GET": {
			t := template.Must(template.ParseFiles("C:\\go-work\\src\\electronic-school-diary\\view\\login.html"))
			if err := t.Execute(w, struct{}{}); err != nil {
				cntrl.ErrorLogger.Println("failed to render template: %s ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.WriteHeader(http.StatusOK)
		}

		case "POST": {
			if err := r.ParseForm(); err != nil {
				cntrl.ErrorLogger.Println("failed to parse form: %s ", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fEmail := r.Form.Get("name")
			fPassword := r.Form.Get("pwd")

			authRepo := cntrl.Repository.GetAuthRepository()
			dbPass, err := authRepo.GetTeacherPass(fEmail)
			if err != nil {
				cntrl.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			if pwd.PasswordsAreEqual(fPassword, dbPass) {
				defer http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

				cookie := createCookie(300)
				session := model.Session{
					Email: fEmail,
					Session_Token: cookie.Value,
				}
				if err = authRepo.StoreSession(session); err != nil {
					cntrl.ErrorLogger.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
				}
				http.SetCookie(w, &cookie)
			} else {
				defer http.Redirect(w, r, "/login", http.StatusSeeOther)
				w.WriteHeader(http.StatusUnauthorized)
			}
		}
	}
	return
}*/

/*func (cntrl *Controller) LogoutHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "GET":
		{
			cookie, _ := r.Cookie("SessionID")
			if err := cntrl.Repository.GetAuthRepository().DeleteSession(cookie.Value); err != nil {
				cntrl.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			cookieToBeDeleted := deleteCookie(cookie.Name)
			http.SetCookie(w, &cookieToBeDeleted)
			t := template.Must(template.ParseFiles("C:\\go-work\\src\\electronic-school-diary\\view\\index.html"))
			if err := t.Execute(w, struct{}{}); err != nil {
				cntrl.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}
	}
}*/

func createCookie (age int) http.Cookie {
	UUID := uuid.New()
	value := UUID.String()
	cookie := http.Cookie{Name : "SessionID", Value: value , MaxAge: age}
	return cookie
}

func deleteCookie (cookieName string) http.Cookie {
	cookie := http.Cookie{Name : cookieName, MaxAge: -1}
	return cookie
}


