package middlewares

import (
	"github.com/zhoujin7/goback"
	"net/http"
	"regexp"
)

// BasicAuth returns middleware for basic authentication.
func BasicAuth(account string, password string, restrictedPaths []string) func(fn goback.HandlerFn) goback.HandlerFn {
	pathRegs := make([]*regexp.Regexp, len(restrictedPaths))
	for i, path := range restrictedPaths {
		pathRegs[i] = regexp.MustCompile(path)
	}

	return func(fn goback.HandlerFn) goback.HandlerFn {
		return func(ctx *goback.Context) error {
			req := ctx.Request()
			for _, pathReg := range pathRegs {
				restricted := pathReg.MatchString(req.URL.Path)
				if restricted {
					if userId, pwd, ok := req.BasicAuth(); ok && userId == account && pwd == password {
						return fn(ctx)
					} else {
						ctx.Response().Header().Set("WWW-Authenticate", "Basic realm=Restricted")
						return ctx.HTML(http.StatusUnauthorized, "Unauthorized")
					}
				}
			}
			return fn(ctx)
		}
	}
}
