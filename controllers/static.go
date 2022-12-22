package controllers

import (
	"html/template"
	"net/http"

	"github.com/psanti93/lenslockedv2.com/views"
)

func StaticHandler(tmpl views.Template) http.HandlerFunc {
	//closure example
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func FAQ(tmpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Nope",
		},
		{
			Question: "How many courses are there?",
			Answer:   "23",
		},
		{
			Question: "How many licks does it take to get to the center of a tootsie pop?",
			Answer:   "364",
		},
		{
			Question: "How do I contact support?",
			Answer:   `Email us - <a href="paulsantiago@gmail.com">paulysupport@gmail.com</a>`,
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, questions)
	}
}
