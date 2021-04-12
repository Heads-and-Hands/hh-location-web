package handlers

import (
	"encoding/json"
	"hh-location-web/hh-location/provider"
	"log"
	"net/http"
)

var BeaconGetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Println("beacon handler")

	beacons := provider.GetProvider().GetBeacons()

	var payload, _ = json.Marshal(beacons)
	w.Write([]byte(payload))
})
