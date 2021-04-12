package models

type Device struct {
	Id   int    `db:"id" json:"id"`
	Uid  string `db:"uid" json:"uid"`
	Name string `db:"name" json:"name"`
	Nickname string `db:"nickname" json:"nickname"`
	OSName string `db:"os_type" json:"os_type"`
	OSVersion string `db:"os_version" json:"os_version"`
	Shell string `db:"shell" json:"shell"`
	Resolution string `db:"resolution" json:"resolution"`
	Type string `db:"type" json:"type"`
	Comment string `db:"comment" json:"comment"`
	TokenUid string `db:"token_uid" json:"token_uid"`
	OwnerId   int    `db:"owner_id" json:"owner_id"`
	Private bool    `db:"private" json:"private"`
}

func (Device) TableName() string {
	return "device"
}
