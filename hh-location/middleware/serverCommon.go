package middleware

import "net/http"

var MobileClientCommonHandler = func(next http.Handler) http.Handler {
	srvHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
	return http.Handler(commonHandler(srvHandler))
}
