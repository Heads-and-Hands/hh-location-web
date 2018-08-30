package middleware

import "net/http"

func WebClientCommonHandler(next http.Handler) http.Handler {
	clnHandle := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
	return http.Handler(commonHandler(clnHandle))
}
