package middleware

import (
	"net/http"
	"log"
	"encoding/json"
)

var commonHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI)
		next.ServeHTTP(w, r)
	})
}

var NotImplementedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var success = map[string]string{"message": "Not implemented"}
	var payload, _ = json.Marshal(success)
	w.Write([]byte(payload))
})