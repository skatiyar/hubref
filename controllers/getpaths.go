package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/skatiyar/hubref/database"
	"github.com/skatiyar/hubref/errors"
	"github.com/skatiyar/hubref/utilities"
)

func GetPaths(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, dbErr := database.GetDB()
	if dbErr != nil {
		utilities.SendError(w, errors.DBError(dbErr.Error()))
		return
	}

	var paths []database.Path
	count, queryErr := db.Model(&paths).
		Column("id", "path", "type", "created_at", "updated_at").
		Order("path").
		SelectAndCount()
	if queryErr != nil {
		utilities.SendError(w, errors.DBError(queryErr.Error()))
		return
	}

	utilities.Send200(w, struct {
		Count int             `json:"count"`
		Paths []database.Path `json:"paths"`
	}{count, paths})
}
