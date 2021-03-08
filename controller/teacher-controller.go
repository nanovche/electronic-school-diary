package controller

import (
	"electronic-school-diary/service"
	"electronic-school-diary/service/utils"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TeacherController struct {
	TeacherService service.ITeacherService
	StudentService service.IStudentService
	MarkService service.IMarkService

}

func NewTeacherController(teacherService service.ITeacherService,
	studentService service.IStudentService, markService service.IMarkService) TeacherController {
	return TeacherController{TeacherService: teacherService,
		StudentService: studentService, MarkService: markService}
}

func (tc TeacherController) AddMarkHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

		case "GET": {

			//get name from jwt -> get id -> get his classes from subjectservice
			subjects, err := tc.TeacherService.
				GetIRepository().
				GetTeacherRepository().
				GetClassesForToday()// should be known since //only teachers will have access to this method
			if err != nil {
				//tc.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			studentNames, err := tc.StudentService.GetAllStudentNames()
			if err != nil {
				//tc.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			data := make([][]string, 3)
			data[0] = studentNames
			data[1] = utils.GetMarksAsSliceOfStrings()
			data[2] = subjects

			t := template.Must(template.ParseFiles("templates/assess-student.tmpl", "templates/base.tmpl"))
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

			err = tc.TeacherService.AssessStudent(studentName, teacherName, subjectTitle, markValString)

			if err != nil {
				//tc.ErrorLogger.Printf("failed assessing student: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			t, err := template.ParseFiles("templates/assess-student.tmpl")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
			}
			if err = t.Execute(w, struct{}{}); err != nil {
				//tc.ErrorLogger.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			return
		}
	}
}

func (tc TeacherController) PresentFormMarkHandler(w http.ResponseWriter, r *http.Request) {

		studentNames, err := tc.StudentService.GetAllStudentNames()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		t := template.Must(template.ParseFiles("templates/update-student-mark.tmpl", "templates/base.tmpl"))
		if err := t.Execute(w, studentNames); err != nil {
			//tc.ErrorLogger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		w.WriteHeader(http.StatusOK)
		return

}

func (tc TeacherController) ReadFormDataMarkHandler(w http.ResponseWriter, r *http.Request) {

		err := r.ParseForm()
		if err != nil {
			//tc.ErrorLogger.Printf("failed parsing form: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		studentName := r.Form.Get("studentName")
		studentID, err := tc.StudentService.GetStudentIDByName(studentName)
		if err != nil {
			//tc.ErrorLogger.Printf("failed assessing student: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		//should be resolved by jwt (teachername)
		teacherID, err := tc.TeacherService.GetTeacherIDByName("Biser")
		if err != nil {
			//tc.ErrorLogger.Printf("failed assessing student: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		marks, err := tc.TeacherService.GetAllMarksOfStudentByOneTeacher(studentID, teacherID)
		if err != nil {
			//tc.ErrorLogger.Printf("failed assessing student: %s", err)
			w.WriteHeader(http.StatusInternalServerError)
		}

		t := template.Must(template.ParseFiles("templates/student-marks.tmpl", "templates/base.tmpl"))
		if err := t.Execute(w, marks); err != nil {
			//tc.ErrorLogger.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		return
}

func (tc TeacherController) UpdateMarkHandler(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close() //handle
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			//log error
			w.WriteHeader(http.StatusInternalServerError)
		}

		values, err := url.ParseQuery(string(data))
		if err != nil {
			//log error
			w.WriteHeader(http.StatusInternalServerError)
		}

		var xmlReqData []string
		for i, _ := range values {
			xmlReqData = append(xmlReqData, values[i][0])
		}

		//handle
		markID := xmlReqData[0]
		date := xmlReqData[1]
		mark := xmlReqData[2]


		err = tc.TeacherService.UpdateMark(markID, date, mark)
		if err != nil {
			fmt.Println(err)
		}
}

func (tc TeacherController) DeleteMarkHandler(w http.ResponseWriter, r *http.Request) {

		defer r.Body.Close() //handle
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			//log error
			w.WriteHeader(http.StatusInternalServerError)
		}

		values, err := url.ParseQuery(string(data))
		if err != nil {
			//log error
			w.WriteHeader(http.StatusInternalServerError)
		}

		var xmlReqData []string
		for i, _ := range values {
			xmlReqData = append(xmlReqData, values[i][0])
		}
		markID := xmlReqData[0]

		err = tc.TeacherService.DeleteMark(markID)
		if err != nil {
			fmt.Println(err)
		}
}

