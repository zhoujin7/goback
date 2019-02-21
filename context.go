package goback

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Context struct {
	request  *http.Request
	response *Response
	HandlerFn func(Context) error
	bindParams url.Values
}

// HTML sends an HTTP response with status code.
func (ctx *Context)HTML(w http.ResponseWriter, code int, html string) {

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)
	w.Write([]byte(html))
}

// String sends a string response with status code.
func (ctx *Context)String(w http.ResponseWriter, code int, s string) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(code)
	w.Write([]byte(s))
}

// JSON sends a JSON response with status code.
func (ctx *Context)JSON(w http.ResponseWriter, code int, i interface{}) (err error) {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(code)
	w.Write(b)
	return
}

/*
var Context *context

func initContext() {
	Context = &context{
		bindParams: make(map[string][]string),
	}
}

func (ctx *context) GetBindParamFirstValue(paramName string) string {
	if ctx.bindParams[paramName] != nil {
		return ctx.bindParams[paramName][0]
	}
	return ""
}

func (ctx *context) GetBindParamValue(paramName string, index int) string {
	if ctx.bindParams[paramName] != nil {
		return ctx.bindParams[paramName][index]
	}
	return ""
}

func (ctx *context) setBindParamValue(paramName string, paramValue string) {
	if len(ctx.bindParams[paramName]) == 0 {
		ctx.bindParams[paramName] = []string{paramValue}
	} else {
		ctx.bindParams[paramName] = append(ctx.bindParams[paramName], paramValue)
	}
}
*/