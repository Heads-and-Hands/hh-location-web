package handlers

import (
	"beacon/hh-location/provider"
	"encoding/json"
	"log"
	"net/http"
)

var BeaconGetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Println("beacon handler")

	beacons := provider.GetProvider().GetBeacons()

	var payload, _ = json.Marshal(beacons)
	w.Write([]byte(payload))
})
