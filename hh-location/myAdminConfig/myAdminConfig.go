package myAdminConfig

import (
	"beacon/hh-location/configurator"
	"beacon/hh-location/models"
	"github.com/jinzhu/gorm"
	"github.com/qor/admin"
	"net/http"
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
