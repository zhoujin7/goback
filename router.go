package goback

import (
	"net/http"
	"regexp"
	"strings"
)

type HandlerFn func(ctx *Context) error

type Middleware func(fn HandlerFn) HandlerFn

type router struct {
	handlerFuncMap  map[string]map[*regexp.Regexp]HandlerFn
	bindParamStuff  map[string]map[*regexp.Regexp]map[int]string
	middlewareChain []Middleware
}

func (r *router) add(reqMethod string, path string, fn HandlerFn) {
	bindParamReg := regexp.MustCompile(`(:[a-z][[:alnum:]]*)`)
	pathReg := regexp.MustCompile(bindParamReg.ReplaceAllString(path, `[^/]+`))
	r.handlerFuncMap[reqMethod][pathReg] = fn

	if bindParamReg.MatchString(path) {
		pathSegments := strings.Split(path, "/")[1:]
		r.bindParamStuff[reqMethod][pathReg] = make(map[int]string)
		for i := range pathSegments {
			if strings.HasPrefix(pathSegments[i], ":") {
				bindParam := strings.TrimLeft(pathSegments[i], ":")
				r.bindParamStuff[reqMethod][pathReg][i] = bindParam
			}
		}
	}
}

func (r *router) Get(path string, fn HandlerFn) {
	r.add("GET", path, fn)
}

func (r *router) Post(path string, fn HandlerFn) {
	r.add("POST", path, fn)
}

func (r *router) Put(path string, fn HandlerFn) {
	r.add("PUT", path, fn)
}

func (r *router) Delete(path string, fn HandlerFn) {
	r.add("DELETE", path, fn)
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if !reqMethods[req.Method] {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := NewContext(w, req)

	pathReg, fn := r.popPathRegAndHandlerFn(req.Method, req.URL.Path)
	bindParamIndexNameMap := r.bindParamStuff[req.Method][pathReg]
	if bindParamIndexNameMap != nil {
		pathSegments := strings.Split(req.URL.Path, "/")[1:]
		for index, paramName := range bindParamIndexNameMap {
			ctx.setBindParamValue(paramName, pathSegments[index])
		}
	}

	if fn != nil {
		nextFn := fn
		for i := len(r.middlewareChain) - 1; i >= 0; i-- {
			nextFn = r.middlewareChain[i](nextFn)
		}
		nextFn(ctx)
	} else {
		http.NotFound(w, req)
	}
}

func (r *router) popPathRegAndHandlerFn(reqMethod string, path string) (*regexp.Regexp, HandlerFn) {
	for pathReg, fn := range r.handlerFuncMap[reqMethod] {
		if pathReg.FindString(path) == path {
			return pathReg, fn
		}
	}
	return nil, nil
}

func (r *router) Use(m Middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}
