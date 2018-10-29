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
	db  *gorm.DB
}

var ormInstance *ormProvider
var ormDB *gorm.DB
var ormOnce sync.Once

func GetOrmInstance(config *configurator.Configuration) *ormProvider {
	ormDB = getDB(config)
	if ormDB == nil {
		return nil
	} else {
		once.Do(func() {
			ormInstance = &ormProvider{
				db: ormDB,
			}
		})
		return ormInstance
	}
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
		return nil
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
