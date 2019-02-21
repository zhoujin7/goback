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

func (ctx *context) GetBindParamFirstValue(paramName string) string {
	if ctx.bindParams[paramName] != nil {
		return ctx.bindParams[paramName][0]
	}
	return ""
}

func (ctx *context) GetBindParamValue(paramName string, index int) string {
	if ctx.bindParams[paramName] != nil {
		return ctx.bindParams[paramName][index]
	}
	return ""
}

func (ctx *context) setBindParamValue(paramName string, paramValue string) {
	if len(ctx.bindParams[paramName]) == 0 {
		ctx.bindParams[paramName] = []string{paramValue}
	} else {
		ctx.bindParams[paramName] = append(ctx.bindParams[paramName], paramValue)
	}
}
