package middleware

import (
	"log"
	"net/http"
)

func WebCommonHandler(next http.Handler) http.Handler {
	clnHandle := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		log.Println("web handler")
		next.ServeHTTP(w, r)
	})
	return http.Handler(commonHandler(clnHandle))
}
