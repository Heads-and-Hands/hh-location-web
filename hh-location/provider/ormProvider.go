package provider

import (
	"sync"
	"beacon/hh-location/configurator"
	"beacon/hh-location/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/jinzhu/gorm"
	"time"
)

type ormProvider struct {
	cfg *configurator.Configuration
	db *gorm.DB
}

var ormInstance *ormProvider
var ormOnce sync.Once

func GetOrmInstance(config *configurator.Configuration) *ormProvider {
	once.Do(func() {
		ormInstance = &ormProvider{cfg:config}
	})
	return ormInstance
}

func (dbp ormProvider) Close() {
	dbp.GetDB().Close()
}

func (dbp ormProvider) GetDB() *gorm.DB {
	dbString := dbp.cfg.DbString
	log.Println(dbString)

	if dbp.db != nil {
		return dbp.db
	}

	newDb ,err := gorm.Open("mysql", dbString)
	if err != nil {
		log.Println(err)
	}
	dbp.db = newDb
	return newDb
}

func (dbp ormProvider) GetBeacons() []models.Beacon {
	beacons := []models.Beacon{}
	dbp.GetDB().Table("beacon").Find(&beacons)
	return beacons
}

func (dbp ormProvider) GetDevices() []models.Device {
	devices := []models.Device{}
	dbp.GetDB().Table("device").Find(&devices)
	return devices
}

func (dbp ormProvider) GetDevicesPositions() []models.DevicesPositions {
	db := dbp.GetDB()
	devices := []models.Device{}
	db.Table("device").Find(&devices)

	positions := []models.DevicesPositions{}
	for _, elem := range devices {
		p := models.Position{}
		db.Table("position").Where("device_id = ?", elem.ID).Last(&p)
		dp := models.DevicesPositions {
			ID: p.ID,
			PosY: p.PosY,
			PosX: p.PosX,
			Time: p.Time,
			DeviceID: elem.ID,
			DeviceName: elem.Name,
		}
		positions = append(positions, dp)
	}

	return positions
}

func (dbp ormProvider) PostPosition(p models.Position) {
	db := dbp.GetDB()
	p.Time = time.Now()
	db.Table("position").Create(&p)
}