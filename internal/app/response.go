package app

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewResponse(status int, message string) *Response {
	return &Response{Status: status, Message: strings.ToLower(message)}
}

func (r *Response) ToJson() []byte {
	// the error here is impossible due to well known object.
	resp, _ := json.Marshal(r)
	return resp
}

func (r *Response) WriteResponse(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(r.Status)
	_, _ = w.Write(r.ToJson())
}
