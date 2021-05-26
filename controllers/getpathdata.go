package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skatiyar/hubref/database"
	"github.com/skatiyar/hubref/errors"
	"github.com/skatiyar/hubref/utilities"
)

func GetPathData(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := p.ByName("path")
	db, dbErr := database.GetDB()
	if dbErr != nil {
		w.Header().Set("content-type", "application/json")
		utilities.SendError(w, errors.DBError(dbErr.Error()))
		return
	}

	var result database.Path
	queryErr := db.Model(&result).Where("path.path = ?", path).Select()
	if queryErr != nil {
		w.Header().Set("content-type", "application/json")
		if database.NilRowError(queryErr) {
			utilities.SendError(w, errors.NotFoundError(path+" not found"))
		} else {
			utilities.SendError(w, errors.DBError(queryErr.Error()))
		}
		return
	}

	w.Header().Set("content-type", result.Type)
	w.Header().Set("last-modified", result.UpdatedAt.UTC().Format(http.TimeFormat))
	w.WriteHeader(200)
	w.Write([]byte(result.Data))
}
