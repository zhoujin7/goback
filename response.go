package goback

import "net/http"

type response struct {
	http.ResponseWriter
	statusCode int
}

func (w *response) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *response) StatusCode() int {
	return w.statusCode
}