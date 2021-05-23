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

func ErrorResponse(statusCode int, msg, err string) []byte {
	return []byte(`{"statusCode":` + strconv.Itoa(statusCode) +
		`,"data":{"error":"` + strings.Replace(err, `"`, `\"`, -1) + `"` +
		`,"message":"` + strings.Replace(msg, `"`, `\"`, -1) + `"}}`)
}

func SendError(rw http.ResponseWriter, err error) {
	if val, ok := err.(errors.Error); ok {
		rw.WriteHeader(val.Code)
		rw.Write(ErrorResponse(val.Code, val.Message, val.RawError.Error()))
	} else {
		rw.WriteHeader(500)
		rw.Write(ErrorResponse(500, "UNWRAPPED_ERROR", err.Error()))
	}
}

func Send201(rw http.ResponseWriter, data interface{}) {
	rw.WriteHeader(201)

	if jsonErr := json.NewEncoder(rw).Encode(Response{201, data}); jsonErr != nil {
		SendError(rw, errors.InternalServerError(jsonErr.Error()))
		return
	}
}

func Send200(rw http.ResponseWriter, data interface{}) {
	rw.WriteHeader(200)

	if jsonErr := json.NewEncoder(rw).Encode(Response{200, data}); jsonErr != nil {
		SendError(rw, errors.InternalServerError(jsonErr.Error()))
		return
	}
}
