package handlers

import (
	"beacon/hh-location/models"
	"beacon/hh-location/provider"
	"encoding/json"
	"io"
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
	provider.GetProvider().PostPosition(p)

	var success = map[string]string{"message": "ok"}
	var payload, _ = json.Marshal(success)
	w.Write([]byte(payload))
})

func PositionFromJson(from io.ReadCloser) models.Position {
	type PositionJS struct {
		UID  string
		PosX int
		PosY int
	}

	var d PositionJS
	decoder := json.NewDecoder(from)
	err := decoder.Decode(&d)
	if err != nil {
		panic(err)
	}

	devices := provider.GetProvider().GetDevices(d.UID)
	device := devices[0]

	var p = models.Position{
		DeviceID: device.ID,
		PosX:     d.PosX,
		PosY:     d.PosY,
		Time:     time.Now(),
	}
	return p
}
