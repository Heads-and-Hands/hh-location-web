package handlers

import (
	"encoding/json"
	"hh-location-web/hh-location/models"
	"hh-location-web/hh-location/provider"
	"log"
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
		log.Println(err)
	}
	provider.GetProvider().PostDevice(&d)
	var success = map[string]string{"message": "ok"}
	var payload, _ = json.Marshal(success)
	w.Write([]byte(payload))
})

type OwnerData struct {
	Id  int    `json:"id"`
	OwnerId int `json:"owner_id"`
}

var OwnerPostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var d OwnerData
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&d)
	if err != nil {
		log.Println(err)
	}
	provider.GetProvider().PostOwner(d.Id, d.OwnerId)
	var success = map[string]string{"message": "ok"}
	var payload, _ = json.Marshal(success)
	w.Write([]byte(payload))
})