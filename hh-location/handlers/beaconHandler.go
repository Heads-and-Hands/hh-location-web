package handlers

import (
	"net/http"
	"encoding/json"
	"beacon/hh-location/provider"
)

var BeaconGetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	beacons := provider.GetBeacons()
	var payload, _ = json.Marshal(beacons)
	w.Write([]byte(payload))
})
