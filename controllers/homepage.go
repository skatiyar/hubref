package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Homepage(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	http.ServeFile(w, r, "assets/index.html")
}
