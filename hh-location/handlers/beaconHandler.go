package handlers

import (
	"net/http"
	"encoding/json"
	"beacon/hh-location/provider"
	"log"
)

var BeaconGetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Println("beacon handler")

	beacons := provider.GetProvider().GetBeacons()

	var payload, _ = json.Marshal(beacons)
	w.Write([]byte(payload))
})
