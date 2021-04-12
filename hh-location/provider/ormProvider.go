package provider

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"hh-location-web/hh-location/configurator"
	"hh-location-web/hh-location/models"
	"log"
	"time"
)

type ormProvider struct {
	db  *gorm.DB
}

var OrmInstance *ormProvider

func init() {
	newDB := getDB(configurator.GetConfiguration())
	once.Do(func() {
		OrmInstance = &ormProvider{
			db: newDB,
		}
	})
}

func (dbp ormProvider) Close() {
	dbp.db.Close()
	dbp.db = nil
}

func getDB(cfg *configurator.Configuration) *gorm.DB {
	dbString := cfg.DbString
	log.Println(dbString)

	newDb, err := gorm.Open("mysql", dbString)
	if err != nil {
		log.Println(err)
		return nil
	}
	return newDb
}

func (dbp ormProvider) GetBeacons() []models.Beacon {
	var beacons []models.Beacon
	dbp.db.Table("beacon").Find(&beacons)
	return beacons
}

func (dbp ormProvider) GetDevices(uid string) []models.Device {
	var devices []models.Device
	if uid != "" {
		dbp.db.Table("device").Where("uid like ?", uid).First(&devices)
	} else {
		dbp.db.Table("device").Find(&devices)
	}
	return devices
}

func (dbp ormProvider) GetDevicesPositions() []models.DevicesPositions {
	var devices []models.Device
	dbp.db.Table("device").Find(&devices)

	var positions []models.DevicesPositions
	for _, elem := range devices {
		p := models.Position{}
		dbp.db.Table("position").Where("device_id = ?", elem.Id).Last(&p)
		dp := models.DevicesPositions{
			Id:         p.Id,
			PosY:       p.PosY,
			PosX:       p.PosX,
			Time:       p.Time,
			DeviceId:   elem.Id,
			DeviceName: elem.Name,
		}
		positions = append(positions, dp)
	}

	return positions
}

func (dbp ormProvider) PostDevice(d *models.Device) {
	device := models.Device{}
	dbp.db.Table("device").Where("name like ?", d.Name).First(&device)
	if device.Uid != "" {
		device.Uid = d.Uid
		dbp.db.Save(&device)
	} else {
		dbp.db.Table("device").Create(d)
	}
}

func (dbp ormProvider) PostPosition(p *models.Position) {
	p.Time = time.Now()
	dbp.db.Table("position").Create(p)
}

func (dbp ormProvider) PostOwner(deviceId int, ownerId int) {
	dbp.db.Model(&models.Device{}).Where("id = ?", deviceId).Update("owner_id", ownerId)
}