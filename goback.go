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

// Instance returns a Router instance.
func Instance() *Router {
	instance := &Router{
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

// Run outputs the listening address then calls http.ListenAndServe.
func Run(addr string, router *Router) error {
	fmt.Printf("Listen on %s\n", addr)
	return http.ListenAndServe(addr, router)
}
