package middlewares

import "net/http"

func BasicAuth(account string, password string) func(next http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, req *http.Request) {
			if userId, pwd, ok := req.BasicAuth(); ok && userId == account && pwd == password {
				next.ServeHTTP(w, req)
				return
			}
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
		}
	}
}
