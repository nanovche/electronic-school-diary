package view

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func LayoutFiles() []string {
	files, err := filepath.Glob("templates/layout/*.html")
	if err != nil {
		log.Panic(err)
	}
	return files
}

type View struct {
	Login Page
	Register Page
}

type Page struct {
	Template *template.Template
	Layout string
}

func(page *Page) Render(w http.ResponseWriter, data interface{}) error {
	return page.Template.ExecuteTemplate(w, page.Layout, data)
}
