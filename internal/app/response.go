package app

import "encoding/json"

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func NewResponse(status int, message string) *Response {
	return &Response{Status: status, Message: message}
}

func (r *Response) ToJson() []byte {
	// the error here is impossible due to well known object.
	resp, _ := json.Marshal(r)
	return resp
}
