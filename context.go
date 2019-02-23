package goback

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Context struct {
	request    *http.Request
	response   *response
	bindParams url.Values
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	ctx := &Context{}
	ctx.init(w, req)
	ctx.bindParams = make(url.Values)
	return ctx
}

func (ctx *Context) init(w http.ResponseWriter, req *http.Request) {
	ctx.response = &response{w, 200}
	ctx.request = req
}

func (ctx *Context) Request() *http.Request {
	return ctx.request
}

func (ctx *Context) Response() *response {
	return ctx.response
}

// HTML sends an HTTP response with status code.
func (ctx *Context) HTML(code int, html string) (err error) {
	ctx.response.Header().Set("Content-Type", "text/html; charset=utf-8")
	ctx.response.WriteHeader(code)
	_, err = ctx.response.Write([]byte(html))
	return
}

// String sends a string response with status code.
func (ctx *Context) String(code int, s string) (err error) {
	ctx.response.Header().Set("Content-Type", "text/plain; charset=utf-8")
	ctx.response.WriteHeader(code)
	_, err = ctx.response.Write([]byte(s))
	return
}

// JSON sends a JSON response with status code.
func (ctx *Context) JSON(code int, i interface{}) (err error) {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	ctx.response.Header().Set("Content-Type", "application/json;charset=UTF-8")
	ctx.response.WriteHeader(code)
	_, err = ctx.response.Write(b)
	return
}

func (ctx *Context) BindParamFirstValue(paramName string) string {
	if ctx.bindParams[paramName] != nil {
		return ctx.bindParams[paramName][0]
	}
	return ""
}

func (ctx *Context) BindParamValue(paramName string, index int) string {
	if ctx.bindParams[paramName] != nil {
		return ctx.bindParams[paramName][index]
	}
	return ""
}

func (ctx *Context) setBindParamValue(paramName string, paramValue string) {
	if len(ctx.bindParams[paramName]) == 0 {
		ctx.bindParams[paramName] = []string{paramValue}
	} else {
		ctx.bindParams[paramName] = append(ctx.bindParams[paramName], paramValue)
	}
}
