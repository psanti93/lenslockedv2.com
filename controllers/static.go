package controllers

import (
	"html/template"
	"net/http"
)

type Static struct {
	Template Template
}

func (static Static) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}

func StaticHandler(tmpl Template) http.HandlerFunc {
	//closure example
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, nil)
	}
}

func FAQ(tmpl Template) http.HandlerFunc {
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
			Answer:   `Email us - <a class="underline" href="paulsantiago@gmail.com">paulysupport@gmail.com</a>`,
		},
	}

	return func(w http.ResponseWriter, r *http.Request) {
		tmpl.Execute(w, questions)
	}
}
