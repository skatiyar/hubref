package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skatiyar/hubref/database"
	"github.com/skatiyar/hubref/errors"
	"github.com/skatiyar/hubref/utilities"
)

func DeletePath(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := p.ByName("path")
	db, dbErr := database.GetDB()
	if dbErr != nil {
		utilities.SendError(w, errors.DBError(dbErr.Error()))
		return
	}

	var pathRes database.Path
	if sresErr := db.Model(&pathRes).
		Where("path.path = ?", path).
		Select(); sresErr != nil {
		if database.NilRowError(sresErr) {
			utilities.SendError(w, errors.NotFoundError(path+" not found"))
		} else {
			utilities.SendError(w, errors.DBError(sresErr.Error()))
		}
		return
	}
	result, resultErr := db.Model(&pathRes).
		Where("path.path = ?", path).
		Delete()
	if resultErr != nil {
		utilities.SendError(w, errors.DBError(resultErr.Error()))
		return
	}
	if result.RowsAffected() != 1 {
		utilities.SendError(w, errors.DBError("Unable to delete: "+path))
		return
	}

	utilities.Send200(w, pathRes)
}
