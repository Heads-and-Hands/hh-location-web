package provider

import (
	"beacon/hh-location/configurator"
	"beacon/hh-location/models"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
	"sync"
	"time"
)

type ormProvider struct {
	cfg *configurator.Configuration
	db  *gorm.DB
}

var ormInstance *ormProvider
var ormOnce sync.Once

func GetOrmInstance(config *configurator.Configuration) *ormProvider {
	once.Do(func() {
		ormInstance = &ormProvider{
			cfg: config,
			db:  getDB(config),
		}
	})
	return ormInstance
}

func (dbp ormProvider) Close() {
	dbp.db.Close()
}

func getDB(cfg *configurator.Configuration) *gorm.DB {
	dbString := cfg.DbString
	log.Println(dbString)

	newDb, err := gorm.Open("mysql", dbString)
	if err != nil {
		log.Println(err)
	}
	return newDb
}

func (dbp ormProvider) GetBeacons() []models.Beacon {
	beacons := []models.Beacon{}
	dbp.db.Table("beacon").Find(&beacons)
	return beacons
}

func (dbp ormProvider) GetDevices(uid string) []models.Device {
	devices := []models.Device{}
	if uid != "" {
		dbp.db.Table("device").Where("uid like ?", uid).First(&devices)
	} else {
		dbp.db.Table("device").Find(&devices)
	}

	return devices
}

func (dbp ormProvider) GetDevicesPositions() []models.DevicesPositions {
	devices := []models.Device{}
	dbp.db.Table("device").Find(&devices)

	positions := []models.DevicesPositions{}
	for _, elem := range devices {
		p := models.Position{}
		dbp.db.Table("position").Where("device_id = ?", elem.ID).Last(&p)
		dp := models.DevicesPositions{
			ID:         p.ID,
			PosY:       p.PosY,
			PosX:       p.PosX,
			Time:       p.Time,
			DeviceID:   elem.ID,
			DeviceName: elem.Name,
		}
		positions = append(positions, dp)
	}

	return positions
}

func (dbp ormProvider) PostPosition(p models.Position) {
	p.Time = time.Now()
	dbp.db.Table("position").Create(&p)
}
