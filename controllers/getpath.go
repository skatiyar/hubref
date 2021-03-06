package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skatiyar/hubref/database"
	"github.com/skatiyar/hubref/errors"
	"github.com/skatiyar/hubref/utilities"
)

func GetPath(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := p.ByName("path")
	db, dbErr := database.GetDB()
	if dbErr != nil {
		utilities.SendError(w, errors.DBError(dbErr.Error()))
		return
	}

	var result database.Path
	queryErr := db.Model(&result).Where("path.path = ?", path).Select()
	if queryErr != nil {
		if database.NilRowError(queryErr) {
			utilities.SendError(w, errors.NotFoundError(path+" not found"))
		} else {
			utilities.SendError(w, errors.DBError(queryErr.Error()))
		}
		return
	}

	utilities.Send200(w, result)
}
