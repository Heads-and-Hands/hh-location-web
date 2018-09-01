package provider

import (
	"beacon/hh-location/models"
	"beacon/hh-location/configurator"
	"github.com/jinzhu/gorm"
)

type DataProvider interface {

	GetDB() *gorm.DB
	GetBeacons() []models.Beacon
	GetDevices() []models.Device
	GetDevicesPositions() []models.DevicesPositions

	PostPosition(p models.Position)

	Close()
}

func GetProvider() DataProvider {
	cfg := configurator.GetConfiguration()
	return GetOrmInstance(cfg)
}
