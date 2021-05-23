package controllers

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/julienschmidt/httprouter"
	"github.com/skatiyar/hubref/database"
	"github.com/skatiyar/hubref/errors"
	"github.com/skatiyar/hubref/utilities"
)

type PathData struct {
	Path string `json:"path"`
	Type string `json:"type"`
	Data string `json:"data"`
}

func CreatePath(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	db, dbErr := database.GetDB()
	if dbErr != nil {
		utilities.SendError(w, errors.DBError(dbErr.Error()))
		return
	}

	var pathReq PathData
	if decodeErr := json.NewDecoder(r.Body).Decode(&pathReq); decodeErr != nil {
		utilities.SendError(w, errors.InvalidDataError(decodeErr.Error()))
		return
	}

	parsed, parseErr := url.ParseRequestURI(pathReq.Path)
	if parseErr != nil {
		utilities.SendError(w, errors.InvalidDataError("Invalid path value"))
		return
	}

	path := &database.Path{
		Path: parsed.EscapedPath(),
		Type: "application/json",
		Data: pathReq.Data,
	}

	result, resultErr := db.Model(path).Insert()
	if resultErr != nil {
		if database.UniqueConstraintError(resultErr) {
			utilities.SendError(w, errors.InvalidDataError("Path already exists"))
		} else {
			utilities.SendError(w, errors.DBError(resultErr.Error()))
		}
		return
	}
	if result.RowsAffected() != 1 {
		utilities.SendError(w, errors.InternalServerError("Data not added"))
		return
	}

	utilities.Send201(w, path)
}
