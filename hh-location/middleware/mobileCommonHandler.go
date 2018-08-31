package middleware

import (
	"net/http"
	"log"
)

var MobileCommonHandler = func(next http.Handler) http.Handler {
	srvHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Println("mobile handler")
		next.ServeHTTP(w, r)
	})
	return http.Handler(commonHandler(srvHandler))
}
