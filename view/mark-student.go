package view

import (
	"html/template"
	"log"
	"path/filepath"
)

//update/delete student mark view
type MarkStudentView struct{
	View
	RateStudent Page
	PresentUpdateStudentOptions Page
	UpdateDeleteStudentMark Page
}

var MarkStudent MarkStudentView

func StudentMarkFiles() []string {
	files, err := filepath.Glob("templates/studentmark/*.html")
	if err != nil {
		log.Panic(err)
	}
	files = append(files, LayoutFiles()...)
	return files
}

func init() {

	files := StudentMarkFiles()

	MarkStudent.RateStudent = Page{
		Template: template.Must(template.New("base").ParseFiles(files[0], files[3])),
		Layout:   "base",
	}

	MarkStudent.PresentUpdateStudentOptions = Page{
		Template: template.Must(template.New("base").ParseFiles(files[2],files[3])),
		Layout:   "base",
	}

	MarkStudent.UpdateDeleteStudentMark = Page{
		Template: template.Must(template.New("base").ParseFiles(files[1], files[3])),
		Layout:   "base",
	}

}