package main

import (
	"goback"
	"log"
	"net/http"
)

func main() {
	router := goback.Instance()
	router.Post("/page/:pageNum", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte(req.Form["pageNum"][0]))
	})
	log.Fatal(goback.Run(":8080", router))
}
