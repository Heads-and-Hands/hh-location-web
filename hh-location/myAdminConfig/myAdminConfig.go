package myAdminConfig

import (
	"github.com/qor/admin"
	"beacon/hh-location/models"
	"net/http"
	"github.com/jinzhu/gorm"
	"beacon/hh-location/configurator"
)

func Init() *http.ServeMux {

	dbString := configurator.GetConfiguration().DbString
	DB, _ := gorm.Open("mysql", dbString)
	Admin := admin.New(&admin.AdminConfig{DB: DB})
	Admin.AddResource(&models.Beacon{})
	Admin.AddResource(&models.Device{})

	m := http.NewServeMux()
	Admin.MountTo("/admin", m)

	return m
}
