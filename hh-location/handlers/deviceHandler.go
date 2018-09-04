package handlers

import (
	"beacon/hh-location/models"
	"beacon/hh-location/provider"
	"encoding/json"
	"net/http"
)

var DeviceGetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	devices := provider.GetProvider().GetDevices("")
	var payload, _ = json.Marshal(devices)
	w.Write([]byte(payload))
})

var DevicePostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var d models.Device
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&d)
	if err != nil {
		panic(err)
	}
	provider.GetProvider().PostDevice(&d)
	var success = map[string]string{"message": "ok"}
	var payload, _ = json.Marshal(success)
	w.Write([]byte(payload))
})
