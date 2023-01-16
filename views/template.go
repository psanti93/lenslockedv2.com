package views

import (
	"fmt"
	"html/template"
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
			"csrfField": func() template.HTML {
				return `<!-- TODO: Implement the csrfField -->`
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

	// Solve race condition by creating a copy of template for each request using clone
	tmpl, err := t.htmlTmpl.Clone()

	if err != nil {
		log.Printf("cloningTemplate: %v", err)
		http.Error(w, "There was an error rendering the page", http.StatusInternalServerError)
		return
	}

	// 2. when we execute the populate the placeholder that we define in 1

	tmpl = tmpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return csrf.TemplateField(r)
			},
		})

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err = tmpl.Execute(w, data); err != nil {
		log.Printf("Executing template:%v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}

}
