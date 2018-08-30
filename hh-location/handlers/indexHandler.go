package handlers

import (
	"net/http"
	"encoding/json"
)

var IndexHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	var success = map[string]string{"message": "Hello!"}
	var payload, _ = json.Marshal(success)
	w.Write([]byte(payload))
})
