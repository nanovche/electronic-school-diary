package controller

import (
	"electronic-school-diary/service"
	"electronic-school-diary/utils"
	"html/template"
	"net/http"
	)

type TeacherController struct {
	TeacherService service.ITeacherService
}

func NewTeacherController(teacherService service.ITeacherService) TeacherController {
	return TeacherController{TeacherService: teacherService}
}

func (tc TeacherController) AddMarkHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

		case "GET": {

			subjects, err := tc.TeacherService.
				GetIRepository().
				GetTeacherRepository().
				GetTeacherSubjects(4) // should be known since //only teachers will have access to this method
			if err != nil {
				//tc.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			var subjectTitles []string
			for _, subject := range subjects {
				subjectTitles = append(subjectTitles, subject.Title)
			}


			students, err := tc.TeacherService.
				GetIRepository().
				GetStudentRepository().
				GetAllStudents()
			if err != nil {
				//tc.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			var studentNames []string
				for _, student := range students {
					studentNames = append(studentNames, student.FirstName + " " + student.LastName)
				}
				var marksAsWords []string
				for w,_ := range utils.GetMarks() {
					marksAsWords = append(marksAsWords, w)
				}

				data := make([][]string, 3)
				data[0] = studentNames
				data[1] = marksAsWords
				data[2] = subjectTitles

				t := template.Must(template.ParseFiles("view/assess-student.html"))
				if err := t.Execute(w, data); err != nil {
					//tc.ErrorLogger.Println(err)
					w.WriteHeader(http.StatusInternalServerError)
				}

				w.WriteHeader(http.StatusOK)
				return
			}

		case "POST": {
			err := r.ParseForm()
			if err != nil {
				//tc.ErrorLogger.Printf("failed parsing form: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
			}

			markValString := r.Form.Get("markValue")
			studentName := r.Form.Get("studentName")
			subjectTitle := r.Form.Get("subjectTitle")
			teacherName := r.Form.Get("teacherName")

			var markAsNumber int
			for w, n := range utils.GetMarks() {
				if w == markValString {
					markAsNumber = n
					break
				}
			}
			err = tc.TeacherService.AssessStudent(studentName, teacherName, subjectTitle, markAsNumber)

			if err != nil {
				//tc.ErrorLogger.Printf("failed assessing student: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			t := template.Must(template.ParseFiles("C:\\go-work\\src\\electronic-school-diary\\view\\assess-student.html"))
			if err := t.Execute(w, struct{}{}); err != nil {
				//tc.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}

