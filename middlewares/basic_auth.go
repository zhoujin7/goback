package middlewares

import "net/http"

var BasicAuth = func(next http.HandlerFunc) http.HandlerFunc {
	account := "admin"
	password := "123"
	return func(w http.ResponseWriter, req *http.Request) {
		if userId, pwd, ok := req.BasicAuth(); ok && userId == account && pwd == password {
			next.ServeHTTP(w, req)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
	}
}
