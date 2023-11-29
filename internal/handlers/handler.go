package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	loghelper "github.com/Kawar1mi/crud-noframework/internal/log_helper"
)

func WrapError(w http.ResponseWriter, err error) {
	WrapErrorWithStatus(w, err, http.StatusBadRequest)
}

func WrapErrorWithStatus(w http.ResponseWriter, err error, httpStatus int) {

	loghelper.AddError(err)

	var m = map[string]string{
		"result": "error",
		"data":   err.Error(),
	}

	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosnniff")
	w.WriteHeader(httpStatus)
	fmt.Fprintln(w, string(res))
}

func WrapOk(w http.ResponseWriter, m map[string]any) {
	res, _ := json.Marshal(m)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, string(res))
}
