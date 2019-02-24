package goback

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Context struct {
	request    *http.Request
	response   *response
	pathParams url.Values
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	ctx := &Context{}
	ctx.init(w, req)
	return ctx
}

func (ctx *Context) init(w http.ResponseWriter, req *http.Request) {
	ctx.request = req
	ctx.response = &response{w, 200}
	ctx.pathParams = make(url.Values)
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

func (ctx *Context) PathParamValue(paramName string) string {
	if ctx.pathParams[paramName] != nil {
		return ctx.pathParams[paramName][0]
	}
	return ""
}

func (ctx *Context) PathParamValueByIndex(paramName string, index int) string {
	if ctx.pathParams[paramName] != nil {
		return ctx.pathParams[paramName][index]
	}
	return ""
}

func (ctx *Context) setPathParamValue(paramName string, paramValue string) {
	if len(ctx.pathParams[paramName]) == 0 {
		ctx.pathParams[paramName] = []string{paramValue}
	} else {
		ctx.pathParams[paramName] = append(ctx.pathParams[paramName], paramValue)
	}
}
