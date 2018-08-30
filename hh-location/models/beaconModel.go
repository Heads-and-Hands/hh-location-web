package models

type Beacon struct {
	ID int `db:"id"`
	Name string `db:"name"`
	PosX int `db:"pos_x"`
	PosY int `db:"pos_y"`
}
