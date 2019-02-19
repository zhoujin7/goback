package middlewares

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"time"
)

var Logger = func(next http.HandlerFunc) http.HandlerFunc {
	logger := log.New(os.Stdout, "*goback*", 0)
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		recorder := httptest.NewRecorder()
		next.ServeHTTP(w, req)
		end := time.Since(start)
		logger.Printf(" -- %s - %v \"%s %s %d\" \"%s\" \"%s\" - %v",
			strings.Split(req.RemoteAddr, ":")[0], start.Format("2006-01-02 15:04:05"),
			req.Method, req.URL.Path, recorder.Code, req.Referer(), req.UserAgent(), end)
	}
}
