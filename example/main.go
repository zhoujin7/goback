package main

import (
	"github.com/zhoujin7/goback"
	"github.com/zhoujin7/goback/middlewares"
	"log"
	"net/http"
	"os"
)

func main() {
	router := goback.Instance()

	router.Use(middlewares.Logger(os.Stdout))

	//router.Use(middlewares.BasicAuth("admin", "123"))

	router.Get("/page/:pageNum", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(goback.Context.GetBindParamValue("pageNum")))
	})
	router.Get("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("hello"))
	})
	router.Get("/a+", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("aaaaaaaaaaaa"))
	})
	router.Get("/500", func(w http.ResponseWriter, req *http.Request) {
		w.WriteHeader(500)
		w.Write([]byte("500 Internal Server Error"))
	})

	log.Fatal(goback.Run(":8080", router))
}
