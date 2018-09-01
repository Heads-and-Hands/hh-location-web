package models

import (
	"time"
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