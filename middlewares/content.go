package middlewares

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skatiyar/hubref/errors"
	"github.com/skatiyar/hubref/utilities"
)

func AcceptJSON(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("content-type", "application/json")
		if r.Header.Get("content-type") != "application/json" {
			utilities.SendError(w, errors.InvalidDataError("Invalid content type"))
			return
		}
		h(w, r, ps)
	}
}
