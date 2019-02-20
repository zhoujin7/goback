package goback

import "net/url"

type context struct {
	bindParams url.Values
}

var Context *context

func initContext() {
	Context = &context{
		bindParams: make(map[string][]string),
	}
}

func (context *context) GetBindParamValues(paramName string) []string {
	return context.bindParams[paramName]
}

func (context *context) GetBindParamValue(paramName string) string {
	if paramValues := context.bindParams[paramName]; paramValues != nil {
		return context.bindParams[paramName][0]
	}
	return ""
}

func (context *context) setBindParamValue(paramName string, paramValue string) {
	if len(context.bindParams[paramName]) == 0 {
		context.bindParams[paramName] = []string{paramValue}
	} else {
		context.bindParams[paramName] = append(context.bindParams[paramName], paramValue)
	}
}
