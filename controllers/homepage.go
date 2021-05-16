package controllers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var indexTemplate = template.Must(template.ParseGlob("templates/index.go.tpl"))

func Homepage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	if err := indexTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}
