package middlewares

import (
	"github.com/zhoujin7/goback"
	"io"
	"log"
	"strings"
	"time"
)

// Middleware for logging client requests.
func Logger(out io.Writer) func(fn goback.HandlerFn) goback.HandlerFn {
	return func(fn goback.HandlerFn) goback.HandlerFn {
		logger := log.New(out, "*goback*", 0)
		return func(ctx *goback.Context) error {
			start := time.Now()
			err := fn(ctx)
			elapsed := time.Since(start)
			req := ctx.Request()
			resp := ctx.Response()
			logger.Printf(" -- %s - %v \"%s %s %d\" \"%s\" \"%s\" - %v",
				strings.Split(req.RemoteAddr, ":")[0], start.Format("2006-01-02 15:04:05"),
				req.Method, req.URL.Path, resp.StatusCode(), req.Referer(), req.UserAgent(), elapsed)
			return err
		}
	}
}
