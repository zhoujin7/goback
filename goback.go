package goback

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"
)

var reqMethods = map[string]bool{
	"GET":    true,
	"POST":   true,
	"PUT":    true,
	"DELETE": true,
}

// Instance returns a router instance.
func Instance() *router {
	instance := &router{
		handlerFnMap:   make(map[string]map[*regexp.Regexp]HandlerFn),
		pathParamStore: make(map[string]map[*regexp.Regexp]map[int]string),
		pool: &sync.Pool{
			New: func() interface{} {
				return newContext(nil, nil)
			},
		},
	}
	for reqMethod := range reqMethods {
		instance.handlerFnMap[reqMethod] = make(map[*regexp.Regexp]HandlerFn)
		instance.pathParamStore[reqMethod] = make(map[*regexp.Regexp]map[int]string)
	}
	return instance
}

// Run outputs the listening address and returns http.ListenAndServe.
func Run(addr string, router *router) error {
	fmt.Printf("Listen on %s\n", addr)
	return http.ListenAndServe(addr, router)
}
