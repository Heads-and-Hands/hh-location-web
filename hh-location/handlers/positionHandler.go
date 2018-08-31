package handlers

import (
	"net/http"
	"encoding/json"
	"beacon/hh-location/models"
	"beacon/hh-location/provider"
)

var PositionGetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
	positions := provider.GetProvider().GetDevicesPositions()
	var payload, _ = json.Marshal(positions)
	w.Write([]byte(payload))
})

var PositionPostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){

	p := models.PositionFromJson(r.Body)
	provider.GetProvider().PostPosition(p)

	var success = map[string]string{"message": "ok"}
	var payload, _ = json.Marshal(success)
	w.Write([]byte(payload))
})

