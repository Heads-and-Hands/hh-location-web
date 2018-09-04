package middleware

import (
	"encoding/json"
	"log"
	"net/http"
)

var commonHandler = func(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s \n", r.Method, r.RequestURI)

		w.Header().Set("Access-Control-Allow-Origin", "*")

		var token string = ""

		t := r.Header.Get("token")
		if t != "" {
			token = t
		}

		keys, ok := r.URL.Query()["token"]
		if ok && len(keys[0]) > 0 {
			token = keys[0]
		}

		if token == "" {
			http.Error(w, "Not authorized", 401)
			return
		}

		next.ServeHTTP(w, r)
	})
}

var NotImplementedHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var success = map[string]string{"message": "Not implemented"}
	var payload, _ = json.Marshal(success)
	w.Write([]byte(payload))
})
