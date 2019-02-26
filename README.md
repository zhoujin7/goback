# goback
Goback is a simple Web framework written in Go.
[![GoDoc](https://godoc.org/github.com/zhoujin7/goback?status.svg)](https://godoc.org/github.com/zhoujin7/goback)
[![Build Status](https://www.travis-ci.org/zhoujin7/goback.svg?branch=master)](https://www.travis-ci.org/zhoujin7/goback)
[![codecov](https://codecov.io/gh/zhoujin7/goback/branch/master/graph/badge.svg)](https://codecov.io/gh/zhoujin7/goback)

## Features
Router with regex pattern matching and URI path parameter binding support

Middleware support

Friendly to REST API

## Installation
```
> go get github.com/zhoujin7/goback/...
```

## Example
```go
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

	router.Get("/", func(ctx *goback.Context) error {
		return ctx.HTML(200, `<h1>Welcome back!</h1>`)
	})

	log.Fatal(goback.Run(":8080", router))
}
```
Learn more [example](https://github.com/zhoujin7/goback/tree/master/example).
