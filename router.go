package goback

import (
	"log"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

// HandlerFn defines a function to serve HTTP requests.
type HandlerFn func(ctx *Context) error

// Middleware defines a function to process middleware.
type Middleware func(fn HandlerFn) HandlerFn

// Router contains the route rules map and chained middleware
type Router struct {
	handlerFnMap    map[string]map[*regexp.Regexp]HandlerFn
	pathParamStore  map[string]map[*regexp.Regexp]map[int]string
	middlewareChain []Middleware
	pool            *sync.Pool
}

func (r *Router) add(reqMethod string, path string, fn HandlerFn) {
	pathParamReg := regexp.MustCompile(`(:[a-z][[:alnum:]]*)`)
	pathReg := regexp.MustCompile(pathParamReg.ReplaceAllString(path, `[^/]+`))
	r.handlerFnMap[reqMethod][pathReg] = fn

	if pathParamReg.MatchString(path) {
		pathSegments := strings.Split(path, "/")[1:]
		r.pathParamStore[reqMethod][pathReg] = make(map[int]string)
		for i := range pathSegments {
			if strings.HasPrefix(pathSegments[i], ":") {
				pathParam := strings.TrimLeft(pathSegments[i], ":")
				r.pathParamStore[reqMethod][pathReg][i] = pathParam
			}
		}
	}
}

// Get registers a GET request handler for a path.
func (r *Router) Get(path string, fn HandlerFn) {
	r.add("GET", path, fn)
}

// Post registers a POST request handler for a path.
func (r *Router) Post(path string, fn HandlerFn) {
	r.add("POST", path, fn)
}

// Put registers a PUT request handler for a path.
func (r *Router) Put(path string, fn HandlerFn) {
	r.add("PUT", path, fn)
}

// Delete registers a DELETE request handler for a path.
func (r *Router) Delete(path string, fn HandlerFn) {
	r.add("DELETE", path, fn)
}

// ServeHTTP implements`http.Handler`interface, which serves HTTP requests.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if !reqMethods[req.Method] {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	ctx := r.pool.Get().(*Context)
	defer r.pool.Put(ctx)
	ctx.init(w, req)
	pathReg, fn := r.popPathRegAndHandlerFn(req.Method, req.URL.Path)
	pathParamIndexNameMap := r.pathParamStore[req.Method][pathReg]
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
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}
	} else {
		http.NotFound(w, req)
	}
}

func (r *Router) popPathRegAndHandlerFn(reqMethod string, path string) (*regexp.Regexp, HandlerFn) {
	for pathReg, fn := range r.handlerFnMap[reqMethod] {
		if pathReg.FindString(path) == path {
			return pathReg, fn
		}
	}
	return nil, nil
}

// Use method is used to Load middleware.
func (r *Router) Use(m Middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}
