package views

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"io/fs"
	"log"
	"net/http"

	"github.com/gorilla/csrf"
)

func Must(t Template, err error) Template {
	if err != nil {
		panic(err) //panics should be use when something went wrong that you can't recover from
	}

	return t
}

func ParseFs(fs fs.FS, patterns ...string) (Template, error) {

	// We need to create the custom function prior to parsing the templates
	tmpl := template.New(patterns[0])
	// 1. create a placeholder function of the crsField first
	tmpl = tmpl.Funcs(
		template.FuncMap{
			"csrfField": func() (template.HTML, error) {
				return "", fmt.Errorf("crsfField not implemented")
			},
		},
	)

	tmpl, err := tmpl.ParseFS(fs, patterns...)

	if err != nil {
		return Template{}, fmt.Errorf("parsing template: %w", err)
	}

	return Template{
		htmlTmpl: tmpl,
	}, nil
}

type Template struct {
	htmlTmpl *template.Template
}

func (t Template) Execute(w http.ResponseWriter, r *http.Request, data interface{}) {

	// Solve race condition by creating a copy of template for each request using Clone()
	tmpl, err := t.htmlTmpl.Clone()

	if err != nil {
		log.Printf("cloningTemplate: %v", err)
		http.Error(w, "There was an error rendering the page", http.StatusInternalServerError)
		return
	}

	// 2. when we execute we populate the placeholder that we define in 1

	tmpl = tmpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
		})

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	// Pros:
	//	1. prevents a half rendered page if there is an error
	// 	2. you can set the status so you don't get superfluous
	// Cons:
	//	1. Can cause performance issue if it's a giant page that you're rendering to in memory. Memory hit

	var buf bytes.Buffer // you are writing into memory
	if err = tmpl.Execute(&buf, data); err != nil {
		log.Printf("Executing template:%v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}

	// copies data from the buffer to the response writer
	io.Copy(w, &buf)

}
