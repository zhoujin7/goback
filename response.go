package goback

import "net/http"

type Response struct {
	Writer http.ResponseWriter
	Status int
}
