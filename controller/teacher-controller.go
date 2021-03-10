package controller

import (
	"electronic-school-diary/service"
	"electronic-school-diary/service/utils"
	"electronic-school-diary/view"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TeacherController struct {
	Controller Controller
	TeacherService service.ITeacherService
	StudentService service.IStudentService
	MarkService service.IMarkService
}

func NewTeacherController(teacherService service.ITeacherService,
	studentService service.IStudentService, markService service.IMarkService, Controller Controller) TeacherController {
	return TeacherController{TeacherService: teacherService,
		StudentService: studentService, MarkService: markService, Controller: Controller}
}

func (tc TeacherController) AddMarkHandlerGet(w http.ResponseWriter, r *http.Request) error{

			//get name from jwt -> get id -> get his classes from subjectservice
			subjects, err := tc.TeacherService.
				GetIRepository().
				GetTeacherRepository().
				GetClassesForToday()// should be known since //only teachers will have access to this method
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return err
			}

			studentNames, err := tc.StudentService.GetAllStudentNames()
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return err
			}

			data := make([][]string, 3)
			data[0] = studentNames
			data[1] = utils.GetMarksAsWords()
			data[2] = subjects

			w.WriteHeader(http.StatusOK)
			return view.MarkStudent.RateStudent.Render(w, data)
}

func (tc TeacherController) AddMarkHandlerPost(w http.ResponseWriter, r *http.Request) error{

			err := r.ParseForm()
			if err != nil {
				//tc.ErrorLogger.Printf("failed parsing form: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
			}

			markValString := r.Form.Get("markValue")
			studentName := r.Form.Get("studentName")
			subjectTitle := r.Form.Get("subjectTitle")
			teacherName := r.Form.Get("teacherName")
			date := r.Form.Get("date")

			err = tc.TeacherService.AssessStudent(studentName, teacherName, subjectTitle, markValString, date)

			w.WriteHeader(http.StatusOK)
			return view.MarkStudent.RateStudent.Render(w, nil)
}

func (tc TeacherController) PresentFormMarkHandler(w http.ResponseWriter, r *http.Request) error {

		studentNames, err := tc.StudentService.GetAllStudentNames()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return err
		}

		w.WriteHeader(http.StatusOK)
		return view.MarkStudent.PresentUpdateStudentOptions.Render(w, studentNames)

}

func (tc TeacherController) ReadFormDataMarkHandler(w http.ResponseWriter, r *http.Request) error {

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

	return view.MarkStudent.UpdateDeleteStudentMark.Render(w, marks)

}

func (tc TeacherController) UpdateMarkHandler(w http.ResponseWriter, r *http.Request) error {

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

		return nil
}

func (tc TeacherController) DeleteMarkHandler(w http.ResponseWriter, r *http.Request) error {

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

		return nil
}

