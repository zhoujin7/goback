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

func (router *router) add(reqMethod string, path string, handlerFunc http.HandlerFunc) {
	bindParamReg := regexp.MustCompile(`(:[a-z][[:alnum:]]*)`)
	pathReg := regexp.MustCompile(bindParamReg.ReplaceAllString(path, `[^/]+`))
	router.handlerFuncMap[reqMethod][pathReg] = handlerFunc

	if bindParamReg.MatchString(path) {
		pathSegments := strings.Split(path, "/")[1:]
		router.bindParamStuff[reqMethod][pathReg] = make(map[int]string)
		for i := range pathSegments {
			if strings.HasPrefix(pathSegments[i], ":") {
				bindParam := strings.TrimLeft(pathSegments[i], ":")
				router.bindParamStuff[reqMethod][pathReg][i] = bindParam
			}
		}
	}
}

func (router *router) Get(path string, handlerFunc http.HandlerFunc) {
	router.add("GET", path, handlerFunc)
}

func (router *router) Post(path string, handlerFunc http.HandlerFunc) {
	router.add("POST", path, handlerFunc)
}

func (router *router) Put(path string, handlerFunc http.HandlerFunc) {
	router.add("PUT", path, handlerFunc)
}

func (router *router) Delete(path string, handlerFunc http.HandlerFunc) {
	router.add("DELETE", path, handlerFunc)
}

func (router *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	if reqMethods[req.Method] != req.Method {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	pathReg, handlerFunc := router.popPathRegAndHandlerFunc(req.Method, req.URL.Path)
	bindParamIndexNameMap := router.bindParamStuff[req.Method][pathReg]
	if bindParamIndexNameMap != nil {
		pathSegments := strings.Split(req.URL.Path, "/")[1:]
		initContext()
		for index, paramName := range bindParamIndexNameMap {
			Context.setBindParamValue(paramName, pathSegments[index])
		}
	}

	if handlerFunc != nil {
		nextHandlerFunc := handlerFunc
		for i := len(router.middlewareChain) - 1; i >= 0; i-- {
			nextHandlerFunc = router.middlewareChain[i](nextHandlerFunc)
		}
		nextHandlerFunc.ServeHTTP(w, req)
	} else {
		http.NotFound(w, req)
	}
}

func (router *router) popPathRegAndHandlerFunc(reqMethod string, path string) (*regexp.Regexp, http.HandlerFunc) {
	for pathReg, handlerFunc := range router.handlerFuncMap[reqMethod] {
		if pathReg.FindString(path) == path {
			return pathReg, handlerFunc
		}
	}
	return nil, nil
}

func (router *router) Use(m middleware) {
	router.middlewareChain = append(router.middlewareChain, m)
}
