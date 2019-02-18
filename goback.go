package goback

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"
)

var singleton *router
var once sync.Once

func Instance() *router {
	reqMethods := []string{"GET", "POST", "PUT", "DELETE"}
	once.Do(func() {
		singleton = &router{
			handlerFuncMap: make(map[string]map[*regexp.Regexp]http.HandlerFunc),
			bindParamStuff: make(map[string]map[*regexp.Regexp]map[int]string),
		}
	})
	for _, reqMethod := range reqMethods {
		singleton.handlerFuncMap[reqMethod] = make(map[*regexp.Regexp]http.HandlerFunc)
		singleton.bindParamStuff[reqMethod] = make(map[*regexp.Regexp]map[int]string)
	}
	return singleton
}

func Run(addr string, router *router) error {
	fmt.Printf("Listen on %s\n", addr)
	return http.ListenAndServe(addr, router)
}
