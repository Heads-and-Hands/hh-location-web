package handlers

import (
	"encoding/json"
	"hh-location-web/hh-location/models"
	"hh-location-web/hh-location/provider"
	"io"
	"log"
	"net/http"
	"time"
)

var PositionGetHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	positions := provider.GetProvider().GetDevicesPositions()
	var payload, _ = json.Marshal(positions)
	w.Write([]byte(payload))
})

var PositionPostHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	p := PositionFromJson(r.Body)
	var success map[string]string
	if p != nil {
		provider.GetProvider().PostPosition(p)
		success = map[string]string{"message": "ok"}
	} else {
		success = map[string]string{"message": "error. incorrect uid"}
	}

	var payload, _ = json.Marshal(success)
	w.Write([]byte(payload))
})

func PositionFromJson(from io.ReadCloser) *models.Position {

	var d models.PositionJS
	decoder := json.NewDecoder(from)
	err := decoder.Decode(&d)
	if err != nil {
		log.Println(err)
	}

	devices := provider.GetProvider().GetDevices(d.Uid)
	if len(devices) == 0 {
		return nil
	}
	device := devices[0]

	var p = models.Position{
		DeviceId: device.Id,
		PosX:     d.PosX,
		PosY:     d.PosY,
		Time:     time.Now(),
	}
	return &p
}
