package goback

import (
	"net/http"
	"reflect"
	"regexp"
)

type router struct {
	GET             map[*regexp.Regexp]http.HandlerFunc
	POST            map[*regexp.Regexp]http.HandlerFunc
	middlewareChain []http.HandlerFunc
	bindParams      *bindParams
}

type bindParams struct {
	GET  map[*regexp.Regexp]map[int]string
	POST map[*regexp.Regexp]map[int]string
}

func (router *router) add(reqMethod string, path string, handlerFunc http.HandlerFunc) {
	structValue := reflect.ValueOf(router)
	fieldValue := reflect.Indirect(structValue).FieldByName(reqMethod)
	re := regexp.MustCompile(path)
	fieldValue.SetMapIndex(reflect.ValueOf(re), reflect.ValueOf(handlerFunc))
}

func (router *router) Get(path string, handlerFunc http.HandlerFunc) {
	router.add("GET", path, handlerFunc)
}

func (router *router) Post(path string, handlerFunc http.HandlerFunc) {
	router.add("POST", path, handlerFunc)
}

func (router *router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	reqMethods := []string{"GET", "POST"}
	if !Contains(reqMethods, req.Method) {
		//return 不支持的方法
		return
	}

	structValue := reflect.ValueOf(router)
	fieldValue := reflect.Indirect(structValue).FieldByName(req.Method)
	reFunMap := fieldValue.Interface().(map[*regexp.Regexp]http.HandlerFunc)
	for re := range reFunMap {
		if re.MatchString(req.URL.Path) {
			fun := reFunMap[re]
			fun.ServeHTTP(w, req)
		}
	}
}
