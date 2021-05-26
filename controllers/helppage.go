package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Helppage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "assets/help.html")
}
