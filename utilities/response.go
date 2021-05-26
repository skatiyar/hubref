package utilities

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/skatiyar/hubref/errors"
)

type Response struct {
	StatusCode int         `json:"statusCode"`
	Data       interface{} `json:"data"`
}

func ErrorResponse(statusCode int, cls, err string) []byte {
	return []byte(`{"statusCode":` + strconv.Itoa(statusCode) +
		`,"data":{"message":"` + strings.Replace(err, `"`, `\"`, -1) + `"` +
		`,"error":"` + strings.Replace(cls, `"`, `\"`, -1) + `"}}`)
}

func SendError(rw http.ResponseWriter, err error) {
	if val, ok := err.(errors.Error); ok {
		rw.WriteHeader(val.Code)
		rw.Write(ErrorResponse(val.Code, val.Class, val.RawError.Error()))
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write(ErrorResponse(http.StatusInternalServerError, "UNWRAPPED_ERROR", err.Error()))
	}
}

func Send201(rw http.ResponseWriter, data interface{}) {
	rw.WriteHeader(http.StatusCreated)

	if jsonErr := json.NewEncoder(rw).Encode(Response{http.StatusCreated, data}); jsonErr != nil {
		SendError(rw, errors.InternalServerError(jsonErr.Error()))
		return
	}
}

func Send200(rw http.ResponseWriter, data interface{}) {
	rw.WriteHeader(http.StatusOK)

	if jsonErr := json.NewEncoder(rw).Encode(Response{http.StatusOK, data}); jsonErr != nil {
		SendError(rw, errors.InternalServerError(jsonErr.Error()))
		return
	}
}
