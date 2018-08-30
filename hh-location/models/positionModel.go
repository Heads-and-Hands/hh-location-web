package models

import (
	"time"
	"io"
	"encoding/json"
)

type Position struct {
	ID int `db:"id"`
	PosX int `db:"pos_x"`
	PosY int `db:"pos_y"`
	Time time.Time `db:"time"`
	DeviceID int `db:"device_id"`
}

type DevicesPositions struct {
	ID int `db:"id"`
	PosX int `db:"pos_x"`
	PosY int `db:"pos_y"`
	Time time.Time `db:"time"`
	DeviceID int `db:"device_id"`
	DeviceName string `db:"device_name"`
}

func PositionFromJson(from io.ReadCloser) Position {
	type PositionJS struct {
		DeviceID int
		PosX int
		PosY int
		Time time.Time
	}

	var d PositionJS
	decoder := json.NewDecoder(from)
	err := decoder.Decode(&d)
	if err != nil {
		panic(err)
	}

	var p = Position{
		DeviceID:d.DeviceID,
		PosX:d.PosX,
		PosY:d.PosX,
		Time:d.Time,
	}
	return p
}