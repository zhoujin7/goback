package main

import (
	"goback"
	"net/http"
)

func main() {
	router := goback.Inst()
	router.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello"))
	})
	goback.Run(":8080", router)
}
