package provider

import (
	"beacon/hh-location/configurator"
	"beacon/hh-location/models"
)

type DataProvider interface {
	GetBeacons() []models.Beacon

	GetDevices(uid string) []models.Device
	PostDevice(d *models.Device)

	GetDevicesPositions() []models.DevicesPositions
	PostPosition(p *models.Position)

	Close()
}

func GetProvider() DataProvider {
	cfg := configurator.GetConfiguration()
	return GetOrmInstance(cfg)
}
