package models

type Beacon struct {
	ID int `db:"id"`
	Name string `db:"name"`
	Correction int `db:"correction"`
	PosX int `db:"pos_x"`
	PosY int `db:"pos_y"`
}
