package models

type Device struct {
	ID int `db:"id"`
	UID string `db:"uid"`
	Name string `db:"name"`
}

