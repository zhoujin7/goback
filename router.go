package goback

import (
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

type HandlerFn func(ctx *Context) error

type Middleware func(fn HandlerFn) HandlerFn

type router struct {
	handlerFnMap    map[string]map[*regexp.Regexp]HandlerFn
	pathParamStuff  map[string]map[*regexp.Regexp]map[int]string
	middlewareChain []Middleware
	pool            *sync.Pool
}

func (r *router) add(reqMethod string, path string, fn HandlerFn) {
	pathParamReg := regexp.MustCompile(`(:[a-z][[:alnum:]]*)`)
	pathReg := regexp.MustCompile(pathParamReg.ReplaceAllString(path, `[^/]+`))
	r.handlerFnMap[reqMethod][pathReg] = fn

	if pathParamReg.MatchString(path) {
		pathSegments := strings.Split(path, "/")[1:]
		r.pathParamStuff[reqMethod][pathReg] = make(map[int]string)
		for i := range pathSegments {
			if strings.HasPrefix(pathSegments[i], ":") {
				pathParam := strings.TrimLeft(pathSegments[i], ":")
				r.pathParamStuff[reqMethod][pathReg][i] = pathParam
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

	ctx := r.pool.Get().(*Context)
	ctx.init(w, req)
	pathReg, fn := r.popPathRegAndHandlerFn(req.Method, req.URL.Path)
	pathParamIndexNameMap := r.pathParamStuff[req.Method][pathReg]
	if pathParamIndexNameMap != nil {
		pathSegments := strings.Split(req.URL.Path, "/")[1:]
		for index, paramName := range pathParamIndexNameMap {
			ctx.setPathParamValue(paramName, pathSegments[index])
		}
	}

	if fn != nil {
		nextFn := fn
		for i := len(r.middlewareChain) - 1; i >= 0; i-- {
			nextFn = r.middlewareChain[i](nextFn)
		}
		err := nextFn(ctx)
		if err != nil {
			log.Println(err)
		}
	} else {
		http.NotFound(w, req)
	}
}

func (r *router) popPathRegAndHandlerFn(reqMethod string, path string) (*regexp.Regexp, HandlerFn) {
	for pathReg, fn := range r.handlerFnMap[reqMethod] {
		if pathReg.FindString(path) == path {
			return pathReg, fn
		}
	}
	return nil, nil
}

func (r *router) Use(m Middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}
