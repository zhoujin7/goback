package main

import (
	"errors"
	"github.com/zhoujin7/goback"
	"github.com/zhoujin7/goback/middlewares"
	"log"
	"os"
)

func main() {
	router := goback.Instance()

	router.Use(middlewares.Logger(os.Stdout))

	router.Use(middlewares.BasicAuth("admin", "123", []string{
		`/admin.*`,
		//`/[a`,
	}))

	router.Get("/admin", func(ctx *goback.Context) error {
		return ctx.HTML(200, "<h1 style='color:red;'>Welcome Admin</h1>")
	})

	router.Get("/page/:pageNum", func(ctx *goback.Context) error {
		return ctx.String(200, ctx.PathParamValue("pageNum"))
	})

	router.Get("/", func(ctx *goback.Context) error {
		return ctx.HTML(200, `<h1>Welcome back!</h1>`)
	})

	router.Get("/a+", func(ctx *goback.Context) error {
		return ctx.String(200, `aaaaaaaa`)
	})

	router.Get("/500", func(ctx *goback.Context) error {
		return errors.New("500 Internal Server Error")
	})

	log.Fatal(goback.Run(":8080", router))
}
