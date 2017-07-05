package views

import (
	"net/http"
	"html/template"
	"path/filepath"
)

var LayoutDir string = "views/layouts"

//func NewView(files ...string) *View {
func NewView(layout string, files ...string) *View {
	files = append(files, layoutFiles()...)
	/*
	files = append(files, 
		"views/layouts/footer.gohtml",
		"views/layouts/navbar.gohtml",
		"views/layouts/bootstrap.gohtml")
	*/
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}

	return &View {
		Template: t,
		Layout: layout,
	}
}

type View struct {
	Template *template.Template
	Layout string
}

func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "/*.gohtml")
	if err != nil {
		panic(err)
	}
	return files
}

func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	return v.Template.ExecuteTemplate(w, v.Layout, data)
}
