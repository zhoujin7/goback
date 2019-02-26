package goback

import "net/http"

type response struct {
	http.ResponseWriter
	statusCode int
}

func (resp *response) WriteHeader(code int) {
	resp.statusCode = code
	resp.ResponseWriter.WriteHeader(code)
}

func (resp *response) StatusCode() int {
	return resp.statusCode
}
