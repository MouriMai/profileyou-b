package errors

import (
	"encoding/json"
	"log"
	"net/http"
)

type ApiErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

// Meaning Example
// {
//   "message": "product 1 not found",
//   "status" : 404,
//   "error"  : "not_found"
// }

// [WILL DO REFACTORING]

func APIError(w http.ResponseWriter, errMessage string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	jsonError, err := json.Marshal(ApiErr{Message: errMessage, Status: status, Error: errMessage})
	if err != nil {
		log.Fatal(err)
	}
	w.Write(jsonError)
}

func NewBadRequestError(message string) *ApiErr {
	return &ApiErr{
		Message: message,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NotFoundError(message string) *ApiErr {
	return &ApiErr{
		Message: message,
		Status:  http.StatusNotFound,
		Error:   "Not_found",
	}
}

func InternalSeverError(message string) *ApiErr {
	return &ApiErr{
		Message: message,
		Status:  http.StatusInternalServerError,
		Error:   "Server error",
	}
}
