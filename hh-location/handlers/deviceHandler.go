package handlers

import (
	"beacon/hh-location/provider"
	"encoding/json"
	"net/http"
)

var DeviceGetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	devices := provider.GetProvider().GetDevices("")
	var payload, _ = json.Marshal(devices)
	w.Write([]byte(payload))
})
