package goback

import (
	"net/http"
	"regexp"
	"strings"
)

type middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

type router struct {
	handlerFuncMap  map[string]map[*regexp.Regexp]http.HandlerFunc
	bindParamStuff  map[string]map[*regexp.Regexp]map[int]string
	middlewareChain []middleware
}

func (r *router) add(reqMethod string, path string, handlerFunc http.HandlerFunc) {
	bindParamReg := regexp.MustCompile(`(:[a-z][[:alnum:]]*)`)
	pathReg := regexp.MustCompile(bindParamReg.ReplaceAllString(path, `[^/]+`))
	r.handlerFuncMap[reqMethod][pathReg] = handlerFunc

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

func (r *router) Get(path string, handlerFunc http.HandlerFunc) {
	r.add("GET", path, handlerFunc)
}

func (r *router) Post(path string, handlerFunc http.HandlerFunc) {
	r.add("POST", path, handlerFunc)
}

func (r *router) Put(path string, handlerFunc http.HandlerFunc) {
	r.add("PUT", path, handlerFunc)
}

func (r *router) Delete(path string, handlerFunc http.HandlerFunc) {
	r.add("DELETE", path, handlerFunc)
}

func (r *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if !reqMethods[req.Method] {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	pathReg, handlerFunc := r.popPathRegAndHandlerFunc(req.Method, req.URL.Path)
	bindParamIndexNameMap := r.bindParamStuff[req.Method][pathReg]
	if bindParamIndexNameMap != nil {
		pathSegments := strings.Split(req.URL.Path, "/")[1:]
		initContext()
		for index, paramName := range bindParamIndexNameMap {
			Context.setBindParamValue(paramName, pathSegments[index])
		}
	}

	if handlerFunc != nil {
		nextHandlerFunc := handlerFunc
		for i := len(r.middlewareChain) - 1; i >= 0; i-- {
			nextHandlerFunc = r.middlewareChain[i](nextHandlerFunc)
		}
		nextHandlerFunc.ServeHTTP(w, req)
	} else {
		http.NotFound(w, req)
	}
}

func (r *router) popPathRegAndHandlerFunc(reqMethod string, path string) (*regexp.Regexp, http.HandlerFunc) {
	for pathReg, handlerFunc := range r.handlerFuncMap[reqMethod] {
		if pathReg.FindString(path) == path {
			return pathReg, handlerFunc
		}
	}
	return nil, nil
}

func (r *router) Use(m middleware) {
	r.middlewareChain = append(r.middlewareChain, m)
}
