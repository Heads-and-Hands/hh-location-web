package middleware

import (
	"net/http"
	"log"
)

func WebCommonHandler(next http.Handler) http.Handler {
	clnHandle := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		log.Println("web handler")
		next.ServeHTTP(w, r)
	})
	return http.Handler(commonHandler(clnHandle))
}
