package goback

import "net/http"

type response struct {
	http.ResponseWriter
	statusCode int
}

// WriteHeader sends an HTTP response header with status code.
func (resp *response) WriteHeader(code int) {
	resp.statusCode = code
	resp.ResponseWriter.WriteHeader(code)
}

// StatusCode returns status code.
func (resp *response) StatusCode() int {
	return resp.statusCode
}
