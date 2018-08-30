package provider

import (
	"github.com/jmoiron/sqlx"
	"beacon/hh-location/models"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"beacon/hh-location/config"
)

func GetDB() *sqlx.DB {
	cfg := config.GetConfiguration()
	dbString := cfg.DbString
	log.Println(dbString)
	db, err := sqlx.Connect("mysql", dbString)
	if err != nil {
		log.Println(err)
	}
	return db
}

func GetBeacons() []models.Beacon {
	beacons := []models.Beacon{}
	db := GetDB()
	err := db.Select(&beacons, "SELECT * from beacon")
    if err != nil {
    	log.Println(err)
	}
	db.Close()
	return beacons
}

func GetDevices() []models.Device {
	devices := []models.Device{}
	db := GetDB()
	err := db.Select(&devices, "SELECT * from device")
	if err != nil {
		log.Println(err)
	}
	db.Close()
	return devices
}

func GetDevicesPositions() []models.DevicesPositions {
	positions := []models.DevicesPositions{}
	db := GetDB()
	err := db.Select(&positions, "select position.id, position.pos_x, position.pos_y, position.time, device.id as device_id, device.name as device_name from device inner join position on (position.device_id = device.id and position.id = (select max(p.id) from position as p where p.device_id = device.id))")
	if err != nil {
		log.Println(err)
	}
	db.Close()
	return positions
}

func PostPosition(p models.Position) {
	db := GetDB()

    _, err := db.Exec("INSERT INTO position (device_id, pos_x, pos_y) VALUES (?, ?, ?)",
    	p.DeviceID, p.PosX, p.PosY)
	if err != nil {
		log.Println(err)
	}

	db.Close()
}