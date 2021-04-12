package provider

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"hh-location-web/hh-location/configurator"
	"hh-location-web/hh-location/models"
	"log"
	"sync"
)

type nativeDbProvider struct {
	db  *sqlx.DB
	cfg *configurator.Configuration
}

var instance *nativeDbProvider
var once sync.Once

func GetDBInstance(config *configurator.Configuration) *nativeDbProvider {
	once.Do(func() {
		instance = &nativeDbProvider{cfg: config}
	})
	return instance
}

func (dbp nativeDbProvider) Close() {
	dbp.getDB().Close()
}

func (dbp nativeDbProvider) getDB() *sqlx.DB {
	dbString := dbp.cfg.DbString
	log.Println(dbString)

	if dbp.db != nil {
		return dbp.db
	}
	newDb, err := sqlx.Connect("mysql", dbString)
	if err != nil {
		log.Println(err)
	}
	dbp.db = newDb
	return newDb
}

func (dbp nativeDbProvider) GetBeacons() []models.Beacon {
	beacons := []models.Beacon{}
	db := dbp.getDB()
	err := db.Select(&beacons, "SELECT * from beacon")
	if err != nil {
		log.Println(err)
	}
	return beacons
}

func (dbp nativeDbProvider) GetDevices(uid string) []models.Device {
	devices := []models.Device{}
	db := dbp.getDB()
	err := db.Select(&devices, "SELECT * from device")
	if err != nil {
		log.Println(err)
	}
	return devices
}

func (dbp nativeDbProvider) GetDevicesPositions() []models.DevicesPositions {
	positions := []models.DevicesPositions{}
	db := dbp.getDB()
	err := db.Select(&positions, "select position.id, position.pos_x, position.pos_y, position.time, device.id as device_id, device.name as device_name from device inner join position on (position.device_id = device.id and position.id = (select max(p.id) from position as p where p.device_id = device.id))")
	if err != nil {
		log.Println(err)
	}
	return positions
}

func (dbp nativeDbProvider) PostPosition(p models.Position) {
	db := dbp.getDB()
	_, err := db.Exec("INSERT INTO position (device_id, pos_x, pos_y) VALUES (?, ?, ?)",
		p.DeviceId, p.PosX, p.PosY)
	if err != nil {
		log.Println(err)
	}
}
