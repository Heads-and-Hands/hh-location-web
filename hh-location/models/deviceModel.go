package models

type Device struct {
	Id   int    `db:"id" json:"id"`
	Uid  string `db:"uid" json:"uid"`
	Name string `db:"name" json:"name"`
}

func (Device) TableName() string {
	return "device"
}
