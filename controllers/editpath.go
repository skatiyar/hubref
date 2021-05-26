package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/skatiyar/hubref/database"
	"github.com/skatiyar/hubref/errors"
	"github.com/skatiyar/hubref/utilities"
)

func EditPath(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	path := p.ByName("path")
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

	pathRes := database.Path{
		Data:      pathReq.Data,
		UpdatedAt: time.Now(),
	}
	uResult, uResultErr := db.Model(&pathRes).
		Where("path.path = ?", path).
		Returning("*").
		UpdateNotZero()
	if uResultErr != nil {
		if database.NilRowError(uResultErr) {
			utilities.SendError(w, errors.InvalidDataError("Invalid path: "+path))
		} else {
			utilities.SendError(w, errors.DBError(uResultErr.Error()))
		}
		return
	}
	if uResult.RowsAffected() != 1 {
		utilities.SendError(w, errors.InvalidDataError("Unable to update path: "+path))
		return
	}

	utilities.Send200(w, pathRes)
}
