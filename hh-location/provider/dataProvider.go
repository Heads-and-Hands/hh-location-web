package provider

import (
	"hh-location-web/hh-location/models"
)

type DataProvider interface {
	GetBeacons() []models.Beacon

	GetDevices(uid string) []models.Device
	PostDevice(d *models.Device)

	GetDevicesPositions() []models.DevicesPositions
	PostPosition(p *models.Position)
	PostOwner(Id int, OwnerId int)

	Close()
}

func GetProvider() DataProvider {
	return OrmInstance
}
