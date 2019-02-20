package main

import (
	"github.com/zhoujin7/goback"
	"github.com/zhoujin7/goback/middlewares"
	"log"
	"net/http"
)

func main() {
	router := goback.Instance()

	router.Use(middlewares.Logger)

	//router.Use(middlewares.BasicAuth)

	router.Get("/page/:pageNum", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(req.Form["pageNum"][0]))
	})
	router.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
	})
	router.Get("/a+", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("aaaaaaaaaaaa"))
	})

	log.Fatal(goback.Run(":8080", router))
}
