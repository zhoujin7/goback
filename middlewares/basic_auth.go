package middlewares

import (
	"github.com/zhoujin7/goback"
	"net/http"
)

func BasicAuth(account string, password string) func(next goback.HandlerFn) goback.HandlerFn {
	return func(next goback.HandlerFn) goback.HandlerFn {
		return func(ctx *goback.Context) error {
			if userId, pwd, ok := req.BasicAuth(); ok && userId == account && pwd == password {
				next.ServeHTTP(w, req)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		}
	}
}
