package goback

import (
	"fmt"
	"net/http"
	"regexp"
)

var reqMethods = map[string]string{
	"GET":    "GET",
	"POST":   "POST",
	"PUT":    "PUT",
	"DELETE": "DELETE",
}

func Instance() *router {
	instance := &router{
		handlerFuncMap: make(map[string]map[*regexp.Regexp]http.HandlerFunc),
		bindParamStuff: make(map[string]map[*regexp.Regexp]map[int]string),
	}
	for _, reqMethod := range reqMethods {
		instance.handlerFuncMap[reqMethod] = make(map[*regexp.Regexp]http.HandlerFunc)
		instance.bindParamStuff[reqMethod] = make(map[*regexp.Regexp]map[int]string)
	}
	return instance
}

func Run(addr string, router *router) error {
	fmt.Printf("Listen on %s\n", addr)
	return http.ListenAndServe(addr, router)
}
