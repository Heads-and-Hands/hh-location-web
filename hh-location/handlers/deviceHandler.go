package handlers

import (
	"net/http"
	"encoding/json"
	"beacon/hh-location/provider"
)

var DeviceHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	devices := provider.GetProvider().GetDevices()
	var payload, _ = json.Marshal(devices)
	w.Write([]byte(payload))
})
