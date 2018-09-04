package models

type Beacon struct {
	Id         int    `db:"id",json:"id"`
	Uid        int    `db:"uid",json:"uid"`
	Name       string `db:"name",json:"name"`
	Correction int    `db:"correction",json:"correction"`
	PosX       int    `db:"pos_x",json:"posX"`
	PosY       int    `db:"pos_y",json:"posY"`
}

func (Beacon) TableName() string {
	return "beacon"
}
