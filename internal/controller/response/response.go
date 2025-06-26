package response

import (
	"encoding/json"
	"net/http"
)

type HttpResponse struct {
	Status     bool        `json:"status"`
	StatusCode string      `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func HttpSuccessResponse(w http.ResponseWriter, status bool, code int, statusCode, message string, data interface{}) {
	res := &HttpResponse{
		Status: status,
		StatusCode: statusCode,
		Message: message,
		Data: data,
	}

	jsonResp, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonResp)
}

func HttpErrorResponse(w http.ResponseWriter, status bool, code int, statusCode, message string) {
	res := &HttpResponse{
		Status: status,
		StatusCode: statusCode,
		Message: message,
	}

	jsonResp, _ := json.Marshal(res)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(jsonResp)
}