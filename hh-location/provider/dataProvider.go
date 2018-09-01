package provider

import (
	"beacon/hh-location/models"
	"beacon/hh-location/configurator"
)

type DataProvider interface {

	GetBeacons() []models.Beacon
	GetDevices(uid string) []models.Device
	GetDevicesPositions() []models.DevicesPositions

	PostPosition(p models.Position)

	Close()
}

func GetProvider() DataProvider {
	cfg := configurator.GetConfiguration()
	return GetOrmInstance(cfg)
}
