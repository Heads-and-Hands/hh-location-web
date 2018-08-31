package provider

import (
	"github.com/jmoiron/sqlx"
	"beacon/hh-location/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"beacon/hh-location/configurator"
	"sync"
)

type dbProvider struct {
	db *sqlx.DB
	cfg *configurator.Configuration
}

var instance *dbProvider
var once sync.Once

func GetDBInstance(config *configurator.Configuration) *dbProvider {
	once.Do(func() {
		instance = &dbProvider{cfg:config}
	})
	return instance
}

func (dbp dbProvider) Close() {
	dbp.getDB().Close()
}

func (dbp dbProvider) getDB() *sqlx.DB {
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

func (dbp dbProvider) GetBeacons() []models.Beacon {
	beacons := []models.Beacon{}
	db := dbp.getDB()
	err := db.Select(&beacons, "SELECT * from beacon")
    if err != nil {
    	log.Println(err)
	}
	return beacons
}

func (dbp dbProvider) GetDevices() []models.Device {
	devices := []models.Device{}
	db := dbp.getDB()
	err := db.Select(&devices, "SELECT * from device")
	if err != nil {
		log.Println(err)
	}
	return devices
}

func (dbp dbProvider) GetDevicesPositions() []models.DevicesPositions {
	positions := []models.DevicesPositions{}
	db := dbp.getDB()
	err := db.Select(&positions, "select position.id, position.pos_x, position.pos_y, position.time, device.id as device_id, device.name as device_name from device inner join position on (position.device_id = device.id and position.id = (select max(p.id) from position as p where p.device_id = device.id))")
	if err != nil {
		log.Println(err)
	}
	return positions
}

func (dbp dbProvider) PostPosition(p models.Position) {
	db := dbp.getDB()
    _, err := db.Exec("INSERT INTO position (device_id, pos_x, pos_y) VALUES (?, ?, ?)",
    	p.DeviceID, p.PosX, p.PosY)
	if err != nil {
		log.Println(err)
	}
}