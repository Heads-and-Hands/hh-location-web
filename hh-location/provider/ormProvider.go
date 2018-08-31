package provider

import (
	"sync"
	"beacon/hh-location/configurator"
	"beacon/hh-location/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"github.com/jinzhu/gorm"
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
	dbp.getDB().Close()
}

func (dbp ormProvider) getDB() *gorm.DB {
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
	dbp.getDB().Find(&beacons)
	return beacons
}

func (dbp ormProvider) GetDevices() []models.Device {
	devices := []models.Device{}
	dbp.getDB().Find(&devices)
	return devices
}

func (dbp ormProvider) GetDevicesPositions() []models.DevicesPositions {
	positions := []models.DevicesPositions{}
	db := dbp.getDB()
	err := db.Select(&positions, "select position.id, position.pos_x, position.pos_y, position.time, device.id as device_id, device.name as device_name from device inner join position on (position.device_id = device.id and position.id = (select max(p.id) from position as p where p.device_id = device.id))")
	if err != nil {
		log.Println(err)
	}
	db.Close()
	return positions
}

func (dbp ormProvider) PostPosition(p models.Position) {
	db := dbp.getDB()
	//_, err := db.Exec("INSERT INTO position (device_id, pos_x, pos_y) VALUES (?, ?, ?)",
	//	p.DeviceID, p.PosX, p.PosY)
	//if err != nil {
	//	log.Println(err)
	//}
	db.Close()
}