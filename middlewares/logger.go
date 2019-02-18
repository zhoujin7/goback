package middlewares

import (
	"net/http"
)

func Logger(handlerFunc http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

	}
}
