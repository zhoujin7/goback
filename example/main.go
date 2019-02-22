package main

import (
	"github.com/zhoujin7/goback"
	"github.com/zhoujin7/goback/middlewares"
	"log"
	"os"
)

func main() {
	router := goback.Instance()

	router.Use(middlewares.Logger(os.Stdout))

	//router.Use(middlewares.BasicAuth("admin", "123"))

	router.Get("/page/:pageNum", func(ctx *goback.Context) error {

		return ctx.String(200, ctx.GetBindParamFirstValue("pageNum"))
	})
	router.Get("/", func(ctx *goback.Context) error {
		return ctx.HTML(200,`<h1>hello</h1>`)
	})

	router.Get("/a+", func(ctx *goback.Context) error {
		return ctx.String(200,`aaaaaaaa`)
	})

	router.Get("/500", func(ctx *goback.Context) error {
		return ctx.String(500,`500 Internal Server Error`)
	})

	log.Fatal(goback.Run(":8080", router))
}
