package goback

import (
	"encoding/json"
	"net/http"
)

// Context provides an HTTP context.
type Context struct {
	request    *http.Request
	response   *response
	pathParams map[string][]string
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	ctx := &Context{}
	ctx.init(w, req)
	return ctx
}

func (ctx *Context) init(w http.ResponseWriter, req *http.Request) {
	ctx.request = req
	ctx.response = &response{w, 200}
	ctx.pathParams = make(map[string][]string)
}

// Request returns *http.Request.
func (ctx *Context) Request() *http.Request {
	return ctx.request
}

// Response returns *response.
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

// PathParamValue returns the first path parameter by name.
func (ctx *Context) PathParamValue(paramName string) string {
	if ctx.pathParams[paramName] != nil && len(ctx.pathParams[paramName]) > 0 {
		return ctx.pathParams[paramName][0]
	}
	return ""
}

// PathParamValueByIndex returns path parameter by name and index.
func (ctx *Context) PathParamValueByIndex(paramName string, index int) string {
	if ctx.pathParams[paramName] != nil && index >= 0 && index < len(ctx.pathParams[paramName]) {
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
