package views

import (
	"fmt"
	"html/template"
	"io/fs"
	"log"
	"net/http"
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
	tmpl = tmpl.Funcs(
		template.FuncMap{
			"csrfField": func() template.HTML {
				return `<input type="hidden" />`
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

func (t Template) Execute(w http.ResponseWriter, data interface{}) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	if err := t.htmlTmpl.Execute(w, data); err != nil {
		log.Printf("Executing template:%v", err)
		http.Error(w, "There was an error executing the template", http.StatusInternalServerError)
		return
	}

}
