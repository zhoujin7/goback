package goback

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"
)

var singleton *router
var once sync.Once

func Inst() *router {
	once.Do(func() {
		singleton = &router{
			GET:  make(map[*regexp.Regexp]http.HandlerFunc),
			POST: make(map[*regexp.Regexp]http.HandlerFunc),
		}
	})
	return singleton
}

func Run(addr string, router *router) error {
	fmt.Printf("Listen on %s\n", addr)
	return http.ListenAndServe(addr, router)
}
