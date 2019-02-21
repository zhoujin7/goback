package goback

import "net/http"

type Response struct {
	http.ResponseWriter
	Status int
}

func (w *Response) WriteHeader(code int) {
	w.Status = code
	w.ResponseWriter.WriteHeader(code)
}