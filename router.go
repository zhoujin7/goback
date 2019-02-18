package goback

import (
	"net/http"
	"regexp"
	"strings"
)

type Middleware func(http.Handler) http.Handler

type router struct {
	handlerFuncMap  map[string]map[*regexp.Regexp]http.HandlerFunc
	bindParamStuff  map[string]map[*regexp.Regexp]map[int]string
	middlewareChain []Middleware
}

func (router *router) add(reqMethod string, path string, handlerFunc http.HandlerFunc) {
	bindParamReg := regexp.MustCompile(`(:[a-z][[:alnum:]]*)`)
	pathReg := regexp.MustCompile(bindParamReg.ReplaceAllString(path, `[^/]+`))
	mergedHandler := http.Handler(handlerFunc)
	for i := len(router.middlewareChain) - 1; i >= 0; i-- {
		mergedHandler = router.middlewareChain[i](handlerFunc)
	}
	router.handlerFuncMap[reqMethod][pathReg] = mergedHandler.ServeHTTP

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
	reqMethods := []string{"GET", "POST", "PUT", "DELETE"}
	if !Contains(reqMethods, req.Method) {
		http.Error(w, "Method Not Allowed", 405)
		return
	}

	pathReg, handlerFunc := router.popPathRegAndHandlerFunc(req.Method, req.URL.Path)
	bindParamIndexNameMap := router.bindParamStuff[req.Method][pathReg]
	if bindParamIndexNameMap != nil {
		pathSegments := strings.Split(req.URL.Path, "/")[1:]
		err := req.ParseForm()
		CheckErr(err)
		for index, paramName := range bindParamIndexNameMap {
			if len(req.Form[paramName]) == 0 {
				paramValues := []string{pathSegments[index]}
				req.Form[paramName] = paramValues
			} else {
				paramValues := []string{pathSegments[index]}
				req.Form[paramName] = append(paramValues, req.Form[paramName]...)
			}
		}
	}

	if handlerFunc != nil {
		handlerFunc.ServeHTTP(w, req)
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

func (router *router) Use(m Middleware) {
	router.middlewareChain = append(router.middlewareChain, m)
}