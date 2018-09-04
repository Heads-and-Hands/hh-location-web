package models

import (
	"time"
)

type Position struct {
	Id       int       `db:"id",json:"id"`
	PosX     int       `db:"pos_x",json:"posX"`
	PosY     int       `db:"pos_y",json:"posY"`
	Time     time.Time `db:"time",json:"time"`
	DeviceId int       `db:"device_id",json:"deviceId"`
}

type DevicesPositions struct {
	Id         int       `db:"id",json:"id"`
	PosX       int       `db:"pos_x",json:"posX"`
	PosY       int       `db:"pos_y",json:"posY"`
	Time       time.Time `db:"time",json:"time"`
	DeviceId   int       `db:"device_id",json:"deviceId"`
	DeviceName string    `db:"device_name",json:"deviceName"`
}

type PositionJS struct {
	Uid  string `json:"uid"`
	PosX int `json:"posX"`
	PosY int `json:"posY"`
}
