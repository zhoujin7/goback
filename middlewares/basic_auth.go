package middlewares

import (
	"github.com/zhoujin7/goback"
	"net/http"
)

func BasicAuth(account string, password string) func(fn goback.HandlerFn) goback.HandlerFn {
	return func(fn goback.HandlerFn) goback.HandlerFn {
		return func(ctx *goback.Context) error {
			if userId, pwd, ok := ctx.Request().BasicAuth(); ok && userId == account && pwd == password {
				return fn(ctx)
			}
			return ctx.HTML(http.StatusUnauthorized, "Unauthorized")
		}
	}
}
